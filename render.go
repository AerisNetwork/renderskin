package renderskin

import (
	"image"
	"math"
)

// RenderFace renders a 2-dimensional image of the face of a Minecraft player's skin.
func RenderFace(img *image.NRGBA, opts Options) (*image.NRGBA, error) {
	if err := validateSkin(img); err != nil {
		return nil, err
	}

	var (
		skin   = convertToNRGBA(img)
		output = removeTransparency(extract(skin, HeadFront))
	)

	if opts.Overlay {
		overlaySkin := fixTransparency(skin)
		composite(output, extract(overlaySkin, HeadOverlayFront), 0, 0)
	}
	return scale(output, opts.Scale), nil
}

// RenderHead renders a 3-dimensional image of the head of a Minecraft player's skin.
func RenderHead(img *image.NRGBA, opts Options) (*image.NRGBA, error) {
	if err := validateSkin(img); err != nil {
		return nil, err
	}

	var (
		skin        = convertToNRGBA(img)
		scaleDouble = float64(opts.Scale)
		output      = image.NewNRGBA(image.Rect(0, 0, 13*opts.Scale+int(math.Floor(scaleDouble*0.855)), 16*opts.Scale))
		frontHead   = removeTransparency(extract(skin, HeadFront))
		topHead     = rotate90(flipHorizontal(removeTransparency(extract(skin, HeadTop))))
		rightHead   = removeTransparency(extract(skin, HeadRight))
	)

	if opts.Overlay {
		overlaySkin := fixTransparency(skin)
		composite(frontHead, extract(overlaySkin, HeadOverlayFront), 0, 0)
		composite(topHead, rotate90(flipHorizontal(extract(overlaySkin, HeadOverlayTop))), 0, 0)
		composite(rightHead, extract(overlaySkin, HeadOverlayRight), 0, 0)
	}

	// Front Head
	compositeTransform(output, scale(frontHead, opts.Scale), frontMatrix, 8*scaleDouble, 12*scaleDouble)
	// Top Head
	compositeTransform(output, scale(topHead, opts.Scale), plantMatrix, 4*scaleDouble, -4*scaleDouble)
	// Right Head
	compositeTransform(output, scale(rightHead, opts.Scale), sideMatrix, 0, 4*scaleDouble)

	if opts.Square {
		return squareAndCenter(output), nil
	}
	return output, nil
}

// RenderBody renders a 3-dimensional image of the full body of a Minecraft player's skin.
func RenderBody(img *image.NRGBA, opts Options) (*image.NRGBA, error) {
	if err := validateSkin(img); err != nil {
		return nil, err
	}

	var (
		skin                       = convertToNRGBA(img)
		scaleDouble                = float64(opts.Scale)
		offset                     = armOffset(opts.Slim)
		isOldSkin                  = oldSkin(skin)
		output                     = image.NewNRGBA(image.Rect(0, 0, 17*opts.Scale+int(math.Ceil(scaleDouble*0.32)), 39*opts.Scale))
		frontHead                  = removeTransparency(extract(skin, HeadFront))
		topHead                    = rotate90(flipHorizontal(removeTransparency(extract(skin, HeadTop))))
		rightHead                  = removeTransparency(extract(skin, HeadRight))
		frontTorso                 = removeTransparency(extract(skin, TorsoFront))
		frontLeftArm  *image.NRGBA = nil
		topLeftArm    *image.NRGBA = nil
		frontRightArm              = removeTransparency(extract(skin, rightArmFront(opts.Slim)))
		topRightArm                = removeTransparency(extract(skin, rightArmTop(opts.Slim)))
		rightRightArm              = removeTransparency(extract(skin, RightArmRight))
		frontLeftLeg  *image.NRGBA = nil
		frontRightLeg              = removeTransparency(extract(skin, RightLegFront))
		rightRightLeg              = removeTransparency(extract(skin, RightLegRight))
	)

	if isOldSkin {
		frontLeftArm = flipHorizontal(frontRightArm)
		topLeftArm = flipHorizontal(topRightArm)
		frontLeftLeg = flipHorizontal(frontRightLeg)
	} else {
		frontLeftArm = removeTransparency(extract(skin, leftArmFront(opts.Slim)))
		topLeftArm = removeTransparency(extract(skin, leftArmTop(opts.Slim)))
		frontLeftLeg = removeTransparency(extract(skin, LeftLegFront))
	}

	if opts.Overlay {
		overlaySkin := fixTransparency(skin)
		composite(topHead, rotate90(flipHorizontal(extract(overlaySkin, HeadOverlayTop))), 0, 0)
		composite(frontHead, extract(overlaySkin, HeadOverlayFront), 0, 0)
		composite(rightHead, extract(overlaySkin, HeadOverlayRight), 0, 0)
		if !isOldSkin {
			composite(frontTorso, extract(overlaySkin, TorsoOverlayFront), 0, 0)
			composite(frontLeftArm, extract(overlaySkin, leftArmOverlayFront(opts.Slim)), 0, 0)
			composite(topLeftArm, extract(overlaySkin, leftArmOverlayTop(opts.Slim)), 0, 0)
			composite(frontRightArm, extract(overlaySkin, rightArmOverlayFront(opts.Slim)), 0, 0)
			composite(topRightArm, extract(overlaySkin, rightArmOverlayTop(opts.Slim)), 0, 0)
			composite(rightRightArm, extract(overlaySkin, RightArmOverlayRight), 0, 0)
			composite(frontLeftLeg, extract(overlaySkin, LeftLegOverlayFront), 0, 0)
			composite(frontRightLeg, extract(overlaySkin, RightLegOverlayFront), 0, 0)
			composite(rightRightLeg, extract(overlaySkin, RightLegOverlayRight), 0, 0)
		}
	}

	// Right Side of Right Leg
	compositeTransform(output, scale(rightRightLeg, opts.Scale), sideMatrix, 4*scaleDouble, 23*scaleDouble)
	// Front of Right Leg
	compositeTransform(output, scale(frontRightLeg, opts.Scale), frontMatrix, 8*scaleDouble, 31*scaleDouble)
	// Front of Left Leg
	compositeTransform(output, scale(frontLeftLeg, opts.Scale), frontMatrix, 12*scaleDouble, 31*scaleDouble)
	// Front of Torso
	compositeTransform(output, scale(frontTorso, opts.Scale), frontMatrix, 8*scaleDouble, 19*scaleDouble)
	// Front of Right Arm
	compositeTransform(output, scale(frontRightArm, opts.Scale), frontMatrix, float64(4+offset)*scaleDouble, 19*scaleDouble)
	// Front of Left Arm
	compositeTransform(output, scale(frontLeftArm, opts.Scale), frontMatrix, 16*scaleDouble, 19*scaleDouble)
	// Top of Left Arm
	compositeTransform(output, scale(rotate270(topLeftArm), opts.Scale), plantMatrix, 15*scaleDouble, float64(offset-1)*scaleDouble)
	// Right Side of Right Arm
	compositeTransform(output, scale(rightRightArm, opts.Scale), sideMatrix, float64(offset)*scaleDouble, float64(15-offset)*scaleDouble)
	// Top of Right Arm
	compositeTransform(output, scale(rotate90(topRightArm), opts.Scale), plantMatrix, 15*scaleDouble, 11*scaleDouble)
	// Front of Head
	compositeTransform(output, scale(frontHead, opts.Scale), frontMatrix, 10*scaleDouble, 13*scaleDouble)
	// Top of Head
	compositeTransform(output, scale(topHead, opts.Scale), plantMatrix, 5*scaleDouble, -5*scaleDouble)
	// Right Side of Head
	compositeTransform(output, scale(rightHead, opts.Scale), sideMatrix, 2*scaleDouble, 3*scaleDouble)

	if opts.Square {
		return squareAndCenter(output), nil
	}
	return output, nil
}

// RenderFrontBody renders a 2-dimensional image of the front of a Minecraft player's skin.
func RenderFrontBody(img *image.NRGBA, opts Options) (*image.NRGBA, error) {
	if err := validateSkin(img); err != nil {
		return nil, err
	}

	var (
		skin                    = convertToNRGBA(img)
		offset                  = armOffset(opts.Slim)
		isOldSkin               = oldSkin(skin)
		output                  = image.NewNRGBA(image.Rect(0, 0, 16, 32))
		frontHead               = removeTransparency(extract(skin, HeadFront))
		frontTorso              = removeTransparency(extract(skin, TorsoFront))
		leftArm    *image.NRGBA = nil
		rightArm                = removeTransparency(extract(skin, rightArmFront(opts.Slim)))
		leftLeg    *image.NRGBA = nil
		rightLeg                = removeTransparency(extract(skin, RightLegFront))
	)

	if isOldSkin {
		leftArm = flipHorizontal(rightArm)
		leftLeg = flipHorizontal(rightLeg)
	} else {
		leftArm = removeTransparency(extract(skin, leftArmFront(opts.Slim)))
		leftLeg = removeTransparency(extract(skin, LeftLegFront))
	}

	if opts.Overlay {
		overlaySkin := fixTransparency(skin)
		composite(frontHead, extract(overlaySkin, HeadOverlayFront), 0, 0)
		if !isOldSkin {
			composite(frontTorso, extract(overlaySkin, TorsoOverlayFront), 0, 0)
			composite(leftArm, extract(overlaySkin, leftArmOverlayFront(opts.Slim)), 0, 0)
			composite(rightArm, extract(overlaySkin, rightArmOverlayFront(opts.Slim)), 0, 0)
			composite(leftLeg, extract(overlaySkin, LeftLegOverlayFront), 0, 0)
			composite(rightLeg, extract(overlaySkin, RightLegOverlayFront), 0, 0)
		}
	}

	// Face
	composite(output, frontHead, 4, 0)
	// Torso
	composite(output, frontTorso, 4, 8)
	// Left Arm
	composite(output, leftArm, 12, 8)
	// Right Arm
	composite(output, rightArm, offset, 8)
	// Left Leg
	composite(output, leftLeg, 8, 20)
	// Right Leg
	composite(output, rightLeg, 4, 20)

	if opts.Square {
		output = squareAndCenter(output)
	}
	return scale(output, opts.Scale), nil
}

// RenderBackBody renders a 2-dimensional image of the back of a Minecraft player's skin.
func RenderBackBody(img image.Image, opts Options) (*image.NRGBA, error) {
	if err := validateSkin(img); err != nil {
		return nil, err
	}

	var (
		skin                      = convertToNRGBA(img)
		offset                    = armOffset(opts.Slim)
		isOldSkin                 = oldSkin(skin)
		output                    = image.NewNRGBA(image.Rect(0, 0, 16, 32))
		backHead                  = removeTransparency(extract(skin, HeadBack))
		backTorso                 = removeTransparency(extract(skin, TorsoBack))
		backLeftArm  *image.NRGBA = nil
		backRightArm              = removeTransparency(extract(skin, rightArmBack(opts.Slim)))
		backLeftLeg  *image.NRGBA = nil
		backRightLeg              = removeTransparency(extract(skin, RightLegBack))
	)

	if isOldSkin {
		backLeftArm = flipHorizontal(backRightArm)
		backLeftLeg = flipHorizontal(backRightLeg)
	} else {
		backLeftArm = removeTransparency(extract(skin, leftArmBack(opts.Slim)))
		backLeftLeg = removeTransparency(extract(skin, LeftLegBack))
	}

	if opts.Overlay {
		overlaySkin := fixTransparency(skin)
		composite(backHead, extract(overlaySkin, HeadOverlayBack), 0, 0)
		if !isOldSkin {
			composite(backTorso, extract(overlaySkin, TorsoOverlayBack), 0, 0)
			composite(backLeftArm, extract(overlaySkin, leftArmOverlayBack(opts.Slim)), 0, 0)
			composite(backRightArm, extract(overlaySkin, rightArmOverlayBack(opts.Slim)), 0, 0)
			composite(backLeftLeg, extract(overlaySkin, LeftLegOverlayBack), 0, 0)
			composite(backRightLeg, extract(overlaySkin, RightLegOverlayBack), 0, 0)
		}
	}

	// Face
	composite(output, backHead, 4, 0)
	// Torso
	composite(output, backTorso, 4, 8)
	// Left Arm
	composite(output, backLeftArm, offset, 8)
	// Right Arm
	composite(output, backRightArm, 12, 8)
	// Left Leg
	composite(output, backLeftLeg, 4, 20)
	// Right Leg
	composite(output, backRightLeg, 8, 20)

	if opts.Square {
		output = squareAndCenter(output)
	}
	return scale(output, opts.Scale), nil
}

// RenderLeftBody renders a 2-dimensional image of the left side of a Minecraft player's skin.
func RenderLeftBody(img *image.NRGBA, opts Options) (*image.NRGBA, error) {
	if err := validateSkin(img); err != nil {
		return nil, err
	}

	var (
		skin                     = convertToNRGBA(img)
		isOldSkin                = oldSkin(skin)
		output                   = image.NewNRGBA(image.Rect(0, 0, 16, 32))
		leftHead                 = removeTransparency(extract(skin, HeadLeft))
		leftLeftArm *image.NRGBA = nil
		leftLeftLeg *image.NRGBA = nil
	)

	if isOldSkin {
		leftLeftArm = flipHorizontal(removeTransparency(extract(skin, rightArmLeft(false))))
		leftLeftLeg = flipHorizontal(removeTransparency(extract(skin, RightLegLeft)))
	} else {
		leftLeftArm = removeTransparency(extract(skin, leftArmLeft(opts.Slim)))
		leftLeftLeg = removeTransparency(extract(skin, LeftLegLeft))
	}

	if opts.Overlay {
		overlaySkin := fixTransparency(skin)
		composite(leftHead, extract(overlaySkin, HeadOverlayLeft), 0, 0)
		if !isOldSkin {
			composite(leftLeftArm, extract(overlaySkin, leftArmOverlayLeft(opts.Slim)), 0, 0)
			composite(leftLeftLeg, extract(overlaySkin, LeftLegOverlayLeft), 0, 0)
		}
	}

	// Left Head
	composite(output, leftHead, 4, 0)
	// Left Arm
	composite(output, leftLeftArm, 6, 8)
	// Left Leg
	composite(output, leftLeftLeg, 6, 20)

	if opts.Square {
		output = squareAndCenter(output)
	}
	return scale(output, opts.Scale), nil
}

// RenderRightBody renders a 2-dimensional image of the right side of a Minecraft player's skin.
func RenderRightBody(img *image.NRGBA, opts Options) (*image.NRGBA, error) {
	if err := validateSkin(img); err != nil {
		return nil, err
	}

	var (
		skin          = convertToNRGBA(img)
		output        = image.NewNRGBA(image.Rect(0, 0, 16, 32))
		rightHead     = removeTransparency(extract(skin, HeadRight))
		rightRightArm = removeTransparency(extract(skin, RightArmRight))
		rightRightLeg = removeTransparency(extract(skin, RightLegRight))
	)

	if opts.Overlay {
		overlaySkin := fixTransparency(skin)
		composite(rightHead, extract(overlaySkin, HeadOverlayRight), 0, 0)
		if !oldSkin(skin) {
			composite(rightRightArm, extract(overlaySkin, RightArmOverlayRight), 0, 0)
			composite(rightRightLeg, extract(overlaySkin, RightLegOverlayRight), 0, 0)
		}
	}

	// Right Head
	composite(output, rightHead, 4, 0)
	// Right Arm
	composite(output, rightRightArm, 6, 8)
	// Right Leg
	composite(output, rightRightLeg, 6, 20)

	if opts.Square {
		output = squareAndCenter(output)
	}
	return scale(output, opts.Scale), nil
}
