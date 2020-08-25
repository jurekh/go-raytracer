package vector

import "fmt"

type Point3 = Vec3

type Ray struct {
	Origin    Vec3
	Direction Vec3
}

func (r Ray) At(t float64) Vec3 {
	return r.Origin.AddV(r.Direction.Mul(t))
}

func (r Ray) String() string {
	return fmt.Sprintf("[Ray Origin:%v, Direction:%v]", r.Origin, r.Direction)
}
