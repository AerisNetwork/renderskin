package renderskin

import (
	"image"
	"math"
)

const (
	angle30Deg     = math.Pi / 6
	isometricScale = 0.86603
)

var (
	sideMatrix  = newRotation(angle30Deg).multiply(newSkewX(angle30Deg)).multiply(newScaleY(isometricScale))
	frontMatrix = newRotation(-angle30Deg).multiply(newSkewX(-angle30Deg)).multiply(newScaleY(isometricScale))
	plantMatrix = newRotation(angle30Deg).multiply(newSkewX(-angle30Deg)).multiply(newScaleY(isometricScale))
)

// Matrix2x2 represents a 2x2 transformation matrix.
type Matrix2x2 [4]float64

// multiply performs matrix multiplication with another 2x2 matrix.
func (m Matrix2x2) multiply(other Matrix2x2) Matrix2x2 {
	return Matrix2x2{
		m[0]*other[0] + m[1]*other[2],
		m[0]*other[1] + m[1]*other[3],
		m[2]*other[0] + m[3]*other[2],
		m[2]*other[1] + m[3]*other[3],
	}
}

// determinant calculates the determinant of the matrix.
func (m Matrix2x2) determinant() float64 {
	return m[0]*m[3] - m[1]*m[2]
}

// inverse returns the inverse of the matrix.
func (m Matrix2x2) inverse() Matrix2x2 {
	invDet := 1.0 / m.determinant()
	return Matrix2x2{
		m[3] * invDet,
		-m[1] * invDet,
		-m[2] * invDet,
		m[0] * invDet,
	}
}

// transform applies the matrix transformation to a point.
func (m Matrix2x2) transform(x, y float64) (float64, float64) {
	return m[0]*x + m[1]*y, m[2]*x + m[3]*y
}

// newScaleY creates a matrix that scales the y-axis.
func newScaleY(scale float64) Matrix2x2 {
	return Matrix2x2{1, 0, 0, scale}
}

// newSkewX creates a matrix that skews along the x-axis.
func newSkewX(angle float64) Matrix2x2 {
	return Matrix2x2{1, math.Tan(angle), 0, 1}
}

// newRotation creates a rotation matrix for the given angle in radians.
func newRotation(angle float64) Matrix2x2 {
	cos, sin := math.Cos(angle), math.Sin(angle)
	return Matrix2x2{cos, -sin, sin, cos}
}

// TransformRect applies a 2x2 transformation matrix to an image rectangle and returns the transformed rectangle.
func TransformRect(m Matrix2x2, rect image.Rectangle) image.Rectangle {
	corners := [4]image.Point{
		rect.Min,
		{rect.Max.X, rect.Min.Y},
		{rect.Min.X, rect.Max.Y},
		rect.Max,
	}
	minX, minY := math.MaxInt, math.MaxInt
	maxX, maxY := math.MinInt, math.MinInt
	for _, corner := range corners {
		tx, ty := m.transform(float64(corner.X), float64(corner.Y))
		x, y := int(math.Floor(tx)), int(math.Floor(ty))

		if x < minX {
			minX = x
		}
		if x+1 > maxX {
			maxX = x + 1
		}
		if y < minY {
			minY = y
		}
		if y+1 > maxY {
			maxY = y + 1
		}
	}
	return image.Rect(minX, minY, maxX, maxY)
}
