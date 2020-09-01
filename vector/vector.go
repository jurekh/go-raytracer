package vector

import (
	"fmt"
	"math"
	"math/rand"
)

// Vec3 represents a 3D vector with X, Y and Z as its component dimensions
type Vec3 struct {
	X, Y, Z float64
}

// ToSlice converts a Vec3 to a slice of float64
func (v Vec3) ToSlice() []float64 {
	return []float64{v.X, v.Y, v.Z}
}

// NewVec3 converts a slice of float64 to a Vec3
func NewVec3(v [3]float64) Vec3 {
	return Vec3{v[0], v[1], v[2]}
}

// AddV is the addition operator for two Vec3 structs
func (v Vec3) AddV(w Vec3) Vec3 {
	return Vec3{v.X + w.X, v.Y + w.Y, v.Z + w.Z}
}

// Add adds a scalar value to all Vec3 components
func (v Vec3) Add(f float64) Vec3 {
	return Vec3{v.X + f, v.Y + f, v.Z + f}
}

// SubV is the subtraction operator for two Vec3 structs
func (v Vec3) SubV(w Vec3) Vec3 {
	return Vec3{v.X - w.X, v.Y - w.Y, v.Z - w.Z}
}

// Sub subtracts a scalar value from all Vec3 components
func (v Vec3) Sub(f float64) Vec3 {
	return v.Add(-f)
}

// Mul multiples a scalar by the Vec3
func (v Vec3) Mul(f float64) Vec3 {
	return Vec3{v.X * f, v.Y * f, v.Z * f}
}

// Div divides the Vec3 by a scalar
func (v Vec3) Div(f float64) Vec3 {
	return v.Mul(1.0 / f)
}

// Dot is a dot product of two Vec3 structs
func (v Vec3) Dot(w Vec3) float64 {
	return v.X*w.X + v.Y*w.Y + v.Z*w.Z
}

// Cross is a cross product of two Vec3 structs
func (v Vec3) Cross(w Vec3) Vec3 {
	return Vec3{v.Y*w.Z - v.Z*w.Y,
		v.Z*w.X - v.X*w.Z,
		v.X*w.Y - v.Y*w.X}
}

// Normalize makes a Vec3 have a unit length
func (v Vec3) Normalize() Vec3 {
	return v.Div(v.Length())
}

// Length returns the lenght of a Vec3
func (v Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

// LengthSquared returns the square of the lenght of a Vec3 struct
func (v Vec3) LengthSquared() float64 {
	return v.Dot(v)
}

// String makes it nice to print a Vec3
func (v Vec3) String() string {
	return fmt.Sprintf("[%f, %f, %f]", v.X, v.Y, v.Z)
}

// DegToRad converts degrees to radians
func DegToRad(deg float64) float64 {
	return deg * math.Pi / 180.0
}

// Random returs a uniform random number from range [min, max)
func Random(min float64, max float64) float64 {
	return min + (max-min)*rand.Float64()
}
