package main

import (
	"image"
	"image/color"
	"image/png"
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

func rayColor(r vec.Ray) color.NRGBA {
	var center = vec.Vec3{0.0, 0.0, -1.0}
	var t = hitSphere(center, 0.5, r)
	if t > 0.0 {
		var N = r.At(t).SubV(center).Normalize()
		return colorf64(N.X+1.0, N.Y+1.0, N.Z+1.0, 1.0)
	}
	unitDirection := r.Direction.Normalize()
	t = 0.5 * (unitDirection.Y + 1.0)
	c1 := vec.Vec3{1.0, 1.0, 1.0}
	c2 := vec.Vec3{0.5, 0.7, 1.0}
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
	var aspectRatio = 16.0 / 9.0
	var imageWidth = 400
	var imageHeight = int(float64(imageWidth) / aspectRatio)

	var viewportHeight = 2.0
	var viewportWidth = aspectRatio * viewportHeight
	var focalLength = 1.0

	origin := vec.Vec3{0, 0, 0}
	hor := vec.Vec3{viewportWidth, 0, 0}
	vrt := vec.Vec3{0, viewportHeight, 0}
	llc := origin.SubV(hor.Div(2.0)).SubV(vrt.Div(2.0)).SubV(vec.Vec3{0, 0, focalLength})

	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))

	for j := 0; j < imageHeight; j++ {
		for i := 0; i < imageWidth; i++ {
			u := float64(i) / float64(imageWidth-1)
			v := 1.0 - float64(j)/float64(imageHeight-1)
			dir := llc.AddV(hor.Mul(u)).AddV(vrt.Mul(v)).SubV(origin)
			ray := vec.Ray{Origin: origin, Direction: dir}
			c := rayColor(ray)
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
