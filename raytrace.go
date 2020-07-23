package main

import (
	"image"
	"image/color"
	"image/png"
	vec "jurekh/raytracing/vector"
	"log"
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
	unitDirection := r.Direction
	unitDirection.Normalize()
	t := 0.5 * (unitDirection.Y + 1.0)
	c1 := &vec.Vec3{1.0, 1.0, 1.0}
	c2 := &vec.Vec3{0.5, 0.7, 1.0}
	c1.Mul(1.0 - t).AddV(*c2.Mul(t))
	return colorf64(c1.X, c1.Y, c1.Z, 1.0)
}

func main() {
	var aspectRatio = 16.0 / 9.0
	var imageWidth = 400
	var imageHeight = int(float64(imageWidth) / aspectRatio)

	var viewportHeight = 2.0
	var viewportWidth = aspectRatio * viewportHeight
	var focalLength = 1.0

	origin := vec.Vec3{0, 0, 0}
	horizontal := vec.Vec3{viewportWidth, 0, 0}
	vertical := vec.Vec3{0, viewportHeight, 0}
	lowerLeftCorner := origin
	h2 := horizontal
	h2.Div(2.0)
	v2 := vertical
	v2.Div(2.0)
	llc := &lowerLeftCorner
	llc.SubV(h2).SubV(v2).SubV(vec.Vec3{0, 0, focalLength})

	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))

	for j := 0; j < imageHeight; j++ {
		for i := 0; i < imageWidth; i++ {
			u := float64(i) / float64(imageWidth-1)
			v := 1.0 - float64(j)/float64(imageHeight-1)
			dh := horizontal
			dv := vertical
			dh.Mul(u)
			dv.Mul(v)
			dir := *llc
			dir.AddV(dh)
			dir.AddV(dv)
			dir.SubV(origin)
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
