package renderskin

import (
	_ "embed"
	"fmt"
	"image"
	"image/draw"
	"math"
)

var (
	zeroPoint = image.Point{}
)

// oldSkin returns true if the skin is a legacy format missing overlay information except for the head.
func oldSkin(img image.Image) bool {
	return img.Bounds().Dy() < 64
}

// validateSkin ...
func validateSkin(img image.Image) error {
	width, height := img.Bounds().Dx(), img.Bounds().Dy()
	if width == 64 && (height == 32 || height == 64) {
		return nil
	}
	return fmt.Errorf("invalid skin dimensions (received=%dx%d, expected=64x32 or 64x64)", width, height)
}

// convertToNRGBA ...
func convertToNRGBA(img image.Image) *image.NRGBA {
	if nrgba, ok := img.(*image.NRGBA); ok {
		return nrgba
	}
	result := image.NewNRGBA(img.Bounds())
	draw.Draw(result, img.Bounds(), img, zeroPoint, draw.Src)
	return result
}

// extract ...
func extract(img *image.NRGBA, rect image.Rectangle) *image.NRGBA {
	output := image.NewNRGBA(image.Rect(0, 0, rect.Dx(), rect.Dy()))
	for x := rect.Min.X; x < rect.Max.X; x++ {
		for y := rect.Min.Y; y < rect.Max.Y; y++ {
			srcIndex := y*img.Stride + x*4
			dstIndex := (y-rect.Min.Y)*output.Stride + (x-rect.Min.X)*4
			copy(output.Pix[dstIndex:dstIndex+4], img.Pix[srcIndex:srcIndex+4])
		}
	}
	return output
}

// scale ...
func scale(img *image.NRGBA, factor int) *image.NRGBA {
	if factor < 2 {
		return img
	}
	bounds := img.Bounds().Size()
	output := image.NewNRGBA(image.Rect(0, 0, bounds.X*factor, bounds.Y*factor))
	for x := 0; x < bounds.X; x++ {
		for y := 0; y < bounds.Y; y++ {
			srcIndex := y*img.Stride + x*4
			color := img.Pix[srcIndex : srcIndex+4]
			for sx := 0; sx < factor; sx++ {
				for sy := 0; sy < factor; sy++ {
					dstIndex := (y*factor+sy)*output.Stride + (x*factor+sx)*4
					copy(output.Pix[dstIndex:dstIndex+4], color)
				}
			}
		}
	}
	return output
}

// removeTransparency ...
func removeTransparency(img *image.NRGBA) *image.NRGBA {
	output := clone(img)
	for i := 3; i < len(output.Pix); i += 4 {
		output.Pix[i] = math.MaxUint8
	}
	return output
}

// flipHorizontal ...
func flipHorizontal(img *image.NRGBA) *image.NRGBA {
	bounds := img.Bounds()
	output := image.NewNRGBA(bounds)
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			srcIndex := y*img.Stride + x*4
			dstIndex := y*output.Stride + (bounds.Max.X-x-1)*4
			copy(output.Pix[dstIndex:dstIndex+4], img.Pix[srcIndex:srcIndex+4])
		}
	}
	return output
}

// fixTransparency ...
func fixTransparency(img *image.NRGBA) *image.NRGBA {
	checkColor := img.Pix[0:4]
	if checkColor[3] == 0 {
		return img
	}
	output := clone(img)
	for i := 0; i < len(output.Pix); i += 4 {
		if equalSlice(checkColor, output.Pix[i:i+4]) {
			output.Pix[i+3] = 0
		}
	}
	return output
}

// clone ...
func clone(img *image.NRGBA) *image.NRGBA {
	output := image.NewNRGBA(img.Bounds())
	draw.Draw(output, img.Bounds(), img, zeroPoint, draw.Src)
	return output
}

// armOffset ...
func armOffset(slim bool) int {
	if slim {
		return 1
	}
	return 0
}

// composite ...
func composite(dst, src *image.NRGBA, offsetX, offsetY int) {
	dstBounds := dst.Bounds()
	srcBounds := src.Bounds()
	for x := srcBounds.Min.X; x < srcBounds.Max.X; x++ {
		for y := srcBounds.Min.Y; y < srcBounds.Max.Y; y++ {
			dstX, dstY := offsetX+x, offsetY+y
			if dstX < dstBounds.Min.X || dstY < dstBounds.Min.Y || dstX >= dstBounds.Max.X || dstY >= dstBounds.Max.Y {
				continue
			}

			srcIndex := y*src.Stride + x*4
			dstIndex := dstY*dst.Stride + dstX*4
			compositeColors(dst.Pix[dstIndex:dstIndex+4], src.Pix[srcIndex:srcIndex+4])
		}
	}
}

// compositeTransform ...
func compositeTransform(dst, src *image.NRGBA, matrix Matrix2x2, offsetX, offsetY float64) {
	srcBounds := src.Bounds()
	dstBounds := dst.Bounds()

	inverse := matrix.inverse()
	transformed := TransformRect(matrix, src.Bounds())
	transformedOffsetX, transformedOffsetY := matrix.transform(offsetX, offsetY)

	for x := transformed.Min.X; x < transformed.Max.X; x++ {
		for y := transformed.Min.Y; y < transformed.Max.Y; y++ {
			dstX := x + int(transformedOffsetX)
			dstY := y + int(transformedOffsetY)

			if dstX < dstBounds.Min.X || dstY < dstBounds.Min.Y || dstX >= dstBounds.Max.X || dstY >= dstBounds.Max.Y {
				continue
			}

			srcX, srcY := inverse.transform(float64(x), float64(y))
			if int(srcX) < srcBounds.Min.X || int(srcY) < srcBounds.Min.Y || int(srcX) >= srcBounds.Max.X || int(srcY) >= srcBounds.Max.Y {
				continue
			}

			srcIndex := int(srcY)*src.Stride + int(srcX)*4
			dstIndex := dstY*dst.Stride + dstX*4
			compositeColors(dst.Pix[dstIndex:dstIndex+4], src.Pix[srcIndex:srcIndex+4])
		}
	}
}

// compositeColors ...
func compositeColors(dst, src []uint8) {
	srcAlpha := uint32(src[3]) * 0x101
	invAlpha := ((1<<16 - 1) - srcAlpha) * 0x101

	dst[0] = uint8((uint32(dst[0])*invAlpha/(1<<16-1) + uint32(src[0])*srcAlpha/0xff) >> 8)
	dst[1] = uint8((uint32(dst[1])*invAlpha/(1<<16-1) + uint32(src[1])*srcAlpha/0xff) >> 8)
	dst[2] = uint8((uint32(dst[2])*invAlpha/(1<<16-1) + uint32(src[2])*srcAlpha/0xff) >> 8)
	dst[3] = uint8((uint32(dst[3])*invAlpha/(1<<16-1) + srcAlpha) >> 8)
}

// rotate90 ...
func rotate90(img *image.NRGBA) *image.NRGBA {
	bounds := img.Bounds().Size()
	output := image.NewNRGBA(image.Rect(0, 0, bounds.Y, bounds.X))
	for x := 0; x < bounds.X; x++ {
		for y := 0; y < bounds.Y; y++ {
			srcIndex := y*img.Stride + x*4
			dstIndex := x*output.Stride + y*4
			copy(output.Pix[dstIndex:dstIndex+4], img.Pix[srcIndex:srcIndex+4])
		}
	}
	return output
}

// rotate270 ...
func rotate270(img *image.NRGBA) *image.NRGBA {
	bounds := img.Bounds().Size()
	output := image.NewNRGBA(image.Rect(0, 0, bounds.Y, bounds.X))
	for x := 0; x < bounds.X; x++ {
		for y := 0; y < bounds.Y; y++ {
			srcIndex := y*img.Stride + x*4
			dstIndex := (bounds.X-x-1)*output.Stride + y*4
			copy(output.Pix[dstIndex:dstIndex+4], img.Pix[srcIndex:srcIndex+4])
		}
	}
	return output
}

// squareAndCenter ...
func squareAndCenter(img *image.NRGBA) *image.NRGBA {
	size := max(img.Rect.Size().X, img.Rect.Size().Y)
	offsetX := (size - img.Rect.Size().X) / 2
	offsetY := (size - img.Rect.Size().Y) / 2

	output := image.NewNRGBA(image.Rect(0, 0, size, size))
	composite(output, img, offsetX, offsetY)
	return output
}

// equalSlice ...
func equalSlice[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
