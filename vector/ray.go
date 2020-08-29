package vector

import (
	"fmt"
	"math"
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

// WithNormal stores information useful for hit testing.
type WithNormal struct {
	Normal    Vec3
	FrontFace bool
}

// SetFaceNormal determines if the ray hit an outward facing face of an object, and sets the normal vector accordingly
func (w WithNormal) SetFaceNormal(r Ray, outwardNormal Vec3) {
	if r.Direction.Dot(outwardNormal) < 0 {
		w.FrontFace = true
		w.Normal = outwardNormal
	} else {
		w.FrontFace = false
		w.Normal = outwardNormal.Mul(-1.0)
	}
}

// HitRecord stores the hit point coordinates, ray coordinate where the hit occured and corresponding WithNormal struct of the hit object
type HitRecord struct {
	WithNormal
	P Point3
	T float64
}

// Hittable represents a hit test on any object that can intersect with a ray. Returns true & hit information when ray intersects with the tested object.
type Hittable interface {
	Hit(r Ray, tMin float64, tMax float64) (hit bool, hr *HitRecord)
}

// Sphere is represented as an origin and radius
type Sphere struct {
	Center Point3
	Radius float64
}

// Hit test for Sphere object. Solves the intersection equation of ray and sphere
func (s Sphere) Hit(r Ray, tMin float64, tMax float64) (hit bool, hr *HitRecord) {
	var oc = r.Origin.SubV(s.Center)
	var a = r.Direction.LengthSquared()
	var halfB = oc.Dot(r.Direction)
	var c = oc.LengthSquared() - s.Radius*s.Radius
	var discriminant = halfB*halfB - a*c

	if discriminant > 0.0 {
		var root = math.Sqrt(discriminant)

		var temp = (-halfB - root) / a
		if temp < tMax && temp > tMin {
			var p = r.At(temp)
			var oN = p.SubV(s.Center).Div(s.Radius)
			var hr = HitRecord{P: p, T: temp}
			(hr.WithNormal).SetFaceNormal(r, oN)
			return true, &hr
		}

		temp = (-halfB + root) / a
		if temp < tMax && temp > tMin {
			var p = r.At(temp)
			var oN = p.SubV(s.Center).Div(s.Radius)
			var hr = HitRecord{P: p, T: temp}
			(hr.WithNormal).SetFaceNormal(r, oN)
			return true, &hr
		}
	}

	return false, nil
}
