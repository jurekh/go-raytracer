package scene

import vec "jurekh/raytracing/vector"

// WithNormal stores information useful for hit testing.
type WithNormal struct {
	Normal    vec.Vec3
	FrontFace bool
}

// SetFaceNormal determines if the ray hit an outward facing face of an object, and sets the normal vector accordingly
func (w *WithNormal) SetFaceNormal(r vec.Ray, outwardNormal vec.Vec3) {
	if r.Direction.Dot(outwardNormal) < 0 {
		(*w).FrontFace = true
		(*w).Normal = outwardNormal
	} else {
		(*w).FrontFace = false
		(*w).Normal = outwardNormal.Mul(-1.0)
	}
}

// HitRecord stores the hit point coordinates, ray coordinate where the hit occured and corresponding WithNormal struct of the hit object
type HitRecord struct {
	WithNormal
	P vec.Point3
	T float64
}

// Hittable represents a hit test on any object that can intersect with a ray. Returns true & hit information when ray intersects with the tested object.
type Hittable interface {
	Hit(r vec.Ray, tMin float64, tMax float64) (hit bool, hr *HitRecord)
}

// HittableList represents a list of hittable objects
type HittableList struct {
	Hittables []Hittable
}

// Hit on HittableList returns the first hit result in terms of distance along the ray
func (hl HittableList) Hit(r vec.Ray, tMin float64, tMax float64) (hit bool, hr *HitRecord) {
	var closestSoFar = tMax
	var closestHr *HitRecord = nil

	for _, h := range hl.Hittables {
		currentHit, currentHr := h.Hit(r, tMin, closestSoFar)
		if currentHit {
			closestSoFar = (*currentHr).T
			closestHr = currentHr
		}
	}

	return closestHr != nil, closestHr
}

// Add on HittableList appends a hittable object to the internal list tested with Hit
func (hl *HittableList) Add(h Hittable) {
	hl.Hittables = append(hl.Hittables, h)
}
