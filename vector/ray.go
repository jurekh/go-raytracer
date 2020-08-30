package vector

import (
	"fmt"
)

// Point3 represents a point in 3d space
type Point3 = Vec3

// Ray is defined by Origin and Direction towards which the Ray points
type Ray struct {
	Origin    Vec3
	Direction Vec3
}

// At returns a point along the Ray starting at Origin and pointing towards t units of Direction
func (r Ray) At(t float64) Vec3 {
	return r.Origin.AddV(r.Direction.Mul(t))
}

func (r Ray) String() string {
	return fmt.Sprintf("[Ray Origin:%v, Direction:%v]", r.Origin, r.Direction)
}
