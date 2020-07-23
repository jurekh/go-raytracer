package vector

type Point3 = Vec3

type Ray struct {
	Origin    Vec3
	Direction Vec3
}

func (r Ray) At(t float64) Vec3 {
	o := r.Origin
	d := r.Direction
	return o.AddV(d.Mul(t))
}
