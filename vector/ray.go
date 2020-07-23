package vector

type Point3 = Vec3

type Ray struct {
	Origin    Vec3
	Direction Vec3
}

func (r Ray) At(t float64) Vec3 {
	return r.Origin.AddV(r.Direction.Mul(t))
}
