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

func clamp(v, min, max float64) float64 {
	if v > max {
		return max
	}
	if v < min {
		return min
	}
	return v
}

func multiSampleColorf64(r, g, b, a float64, samples int) color.NRGBA {
	var scale = 1.0 / float64(samples)
	return color.NRGBA{
		R: uint8(256 * clamp(r*scale, 0.0, 0.999)),
		G: uint8(256 * clamp(g*scale, 0.0, 0.999)),
		B: uint8(256 * clamp(b*scale, 0.0, 0.999)),
		A: uint8(256 * clamp(a, 0.0, 0.999)),
	}
}

func rayColor(r vec.Ray, world *scn.HittableList) vec.Vec3 {
	hit, hr := (*world).Hit(r, 0, math.Inf(0))
	if hit {
		// fmt.Printf("Hit an object: %v\n", *hr)
		var N = (*hr).WithNormal.Normal.AddV(vec.Vec3{X: 1.0, Y: 1.0, Z: 1.0}).Div(2.0)
		return N
		//return colorf64(N.X, N.Y, N.Z, 1.0)
	}

	unitDirection := r.Direction.Normalize()
	t := 0.5 * (unitDirection.Y + 1.0)
	c1 := vec.Vec3{X: 1.0, Y: 1.0, Z: 1.0}
	c2 := vec.Vec3{X: 0.5, Y: 0.7, Z: 1.0}
	c3 := c1.Mul(1.0 - t).AddV(c2.Mul(t))
	return c3
	// return colorf64(c3.X, c3.Y, c3.Z, 1.0)
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
	var camera = scn.NewCamera()
	var samplesPerPixel = 100

	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))

	for j := 0; j < imageHeight; j++ {
		for i := 0; i < imageWidth; i++ {
			pixelColor := vec.Vec3{X: 0, Y: 0, Z: 0}
			for s := 0; s < samplesPerPixel; s++ {
				u := (float64(i) + vec.Random(0.0, 1.0)) / float64(imageWidth-1)
				v := 1.0 - (float64(j)+vec.Random(0.0, 1.0))/float64(imageHeight-1)
				ray := camera.GetRay(u, v)
				pixelColor = pixelColor.AddV(rayColor(ray, &world))
			}
			img.Set(i, j, multiSampleColorf64(pixelColor.X, pixelColor.Y, pixelColor.Z, 1.0, samplesPerPixel))
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
