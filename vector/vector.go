package vector

import (
	"fmt"
	"log"
	"math"
)

type Vec3 struct {
	X, Y, Z float64
}

func (v *Vec3) ToSlice() []float64 {
	return []float64{v.X, v.Y, v.Z}
}

func NewVec3(v [3]float64) Vec3 {
	return Vec3{v[0], v[1], v[2]}
}

func (v *Vec3) AddV(w Vec3) *Vec3 {
	v.X += w.X
	v.Y += w.Y
	v.Z += w.Z
	return v
}

func (v *Vec3) Add(f float64) *Vec3 {
	v.X += f
	v.Y += f
	v.Z += f
	return v
}

func (v *Vec3) SubV(w Vec3) *Vec3 {
	v.X -= w.X
	v.Y -= w.Y
	v.Z -= w.Z
	return v
}

func (v *Vec3) Sub(f float64) *Vec3 {
	return v.Add(-f)
}

func (v *Vec3) Mul(f float64) *Vec3 {
	v.X *= f
	v.Y *= f
	v.Z *= f
	return v
}

func (v *Vec3) Div(f float64) *Vec3 {
	if f == 0.0 {
		log.Fatal("div by zero")
	}
	return v.Mul(1.0 / f)
}

func (v *Vec3) Dot(w Vec3) float64 {
	return v.X*w.X + v.Y*w.Y + v.Z*w.Z
}

func (v *Vec3) Cross(w Vec3) *Vec3 {
	return &Vec3{v.Y*w.Z - v.Z*w.Y,
		v.Z*w.X - v.X*w.Z,
		v.X*w.Y - v.Y*w.X}
}

func (v *Vec3) Normalize() *Vec3 {
	return v.Div(v.Length())
}

func (v *Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v *Vec3) LengthSquared() float64 {
	return v.Dot(*v)
}

func (v *Vec3) String() string {
	return fmt.Sprintf("[%d, %d, %d]", v.X, v.Y, v.Z)
}
