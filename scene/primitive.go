package scene

import (
	vec "jurekh/raytracing/vector"
	"math"
)

// Sphere is represented as an origin and radius
type Sphere struct {
	Center vec.Point3
	Radius float64
}

// Hit test for Sphere object. Solves the intersection equation of ray and sphere
func (s Sphere) Hit(r vec.Ray, tMin float64, tMax float64) (hit bool, hr *HitRecord) {
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
			(&hr).WithNormal.SetFaceNormal(r, oN)
			return true, &hr
		}

		temp = (-halfB + root) / a
		if temp < tMax && temp > tMin {
			var p = r.At(temp)
			var oN = p.SubV(s.Center).Div(s.Radius)
			var hr = HitRecord{P: p, T: temp}
			(&hr).WithNormal.SetFaceNormal(r, oN)
			return true, &hr
		}
	}

	return false, nil
}
