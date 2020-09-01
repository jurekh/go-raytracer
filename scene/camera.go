package scene

import vec "jurekh/raytracing/vector"

// Camera holds camera properties
type Camera struct {
	origin          vec.Point3
	lowerLeftCorner vec.Point3
	horizontal      vec.Vec3
	vertical        vec.Vec3
}

// NewCamera creates a default camera
func NewCamera() (r Camera) {
	var c = Camera{}
	var aspectRatio = 16.0 / 9.0
	var viewportHeight = 2.0
	var viewportWidth = aspectRatio * viewportHeight
	var focalLength = 1.0

	c.origin = vec.Vec3{X: 0, Y: 0, Z: 0}
	c.horizontal = vec.Vec3{X: viewportWidth, Y: 0, Z: 0}
	c.vertical = vec.Vec3{X: 0, Y: viewportHeight, Z: 0}
	c.lowerLeftCorner = c.origin.SubV(c.horizontal.Div(2.0)).SubV(c.vertical.Div(2.0)).SubV(vec.Vec3{X: 0, Y: 0, Z: focalLength})
	return c
}

// GetRay creates a ray for the camera at specified coordinates
func (c Camera) GetRay(u float64, v float64) (rr vec.Ray) {
	dir := c.lowerLeftCorner.AddV(c.horizontal.Mul(u)).AddV(c.vertical.Mul(v)).SubV(c.origin)
	return vec.Ray{Origin: c.origin, Direction: dir}
}
