package main

import (
	"image"
	"image/color"
	"image/png"
	scn "jurekh/raytracing/scene"
	vec "jurekh/raytracing/vector"
	"log"
	"math"
	"os"
)

func colorf64(r, g, b, a float64) color.NRGBA {
	return color.NRGBA{
		R: uint8(255.999 * r),
		G: uint8(255.999 * g),
		B: uint8(255.999 * b),
		A: uint8(255.999 * a),
	}
}

func rayColor(r vec.Ray, world *scn.HittableList) color.NRGBA {
	hit, hr := (*world).Hit(r, 0, math.Inf(0))
	if hit {
		// fmt.Printf("Hit an object: %v\n", *hr)
		var N = (*hr).WithNormal.Normal.AddV(vec.Vec3{X: 1.0, Y: 1.0, Z: 1.0}).Div(2.0)
		return colorf64(N.X, N.Y, N.Z, 1.0)
	}

	unitDirection := r.Direction.Normalize()
	t := 0.5 * (unitDirection.Y + 1.0)
	c1 := vec.Vec3{X: 1.0, Y: 1.0, Z: 1.0}
	c2 := vec.Vec3{X: 0.5, Y: 0.7, Z: 1.0}
	c3 := c1.Mul(1.0 - t).AddV(c2.Mul(t))
	return colorf64(c3.X, c3.Y, c3.Z, 1.0)
}

func hitSphere(center vec.Point3, radius float64, r vec.Ray) float64 {
	var oc = r.Origin.SubV(center)
	var a = r.Direction.LengthSquared()
	var halfB = oc.Dot(r.Direction)
	var c = oc.LengthSquared() - radius*radius
	var discriminant = halfB*halfB - a*c
	if discriminant < 0.0 {
		return -1.0
	}
	return (-halfB - math.Sqrt(discriminant)) / a
}

func main() {
	// Image
	var aspectRatio = 16.0 / 9.0
	var imageWidth = 400
	var imageHeight = int(float64(imageWidth) / aspectRatio)

	// World
	var world = scn.HittableList{}
	(&world).Add(scn.Sphere{Center: vec.Point3{X: 0, Y: 0, Z: -1}, Radius: 0.5})
	(&world).Add(scn.Sphere{Center: vec.Point3{X: 0, Y: -100.5, Z: -1}, Radius: 100})

	// Camera
	var viewportHeight = 2.0
	var viewportWidth = aspectRatio * viewportHeight
	var focalLength = 1.0

	origin := vec.Vec3{X: 0, Y: 0, Z: 0}
	hor := vec.Vec3{X: viewportWidth, Y: 0, Z: 0}
	vrt := vec.Vec3{X: 0, Y: viewportHeight, Z: 0}
	llc := origin.SubV(hor.Div(2.0)).SubV(vrt.Div(2.0)).SubV(vec.Vec3{X: 0, Y: 0, Z: focalLength})

	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))

	for j := 0; j < imageHeight; j++ {
		for i := 0; i < imageWidth; i++ {
			u := float64(i) / float64(imageWidth-1)
			v := 1.0 - float64(j)/float64(imageHeight-1)
			dir := llc.AddV(hor.Mul(u)).AddV(vrt.Mul(v)).SubV(origin)
			ray := vec.Ray{Origin: origin, Direction: dir}
			c := rayColor(ray, &world)
			img.Set(i, j, c)
		}
	}

	f, err := os.Create("image.png")

	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
