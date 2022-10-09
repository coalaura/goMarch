package main

import (
	"image"
	"image/png"
	"math/rand"
	"os"
)

func createImage(pWidth, pHeight int) *image.RGBA {
	return image.NewRGBA(image.Rectangle{
		Min: image.Point{},
		Max: image.Point{
			X: pWidth,
			Y: pHeight,
		},
	})
}

func saveImage(img *image.RGBA, path string) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return err
	}

	err = png.Encode(f, img)

	_ = f.Close()

	return err
}

func random(min, max float64) float64 {
	return (rand.Float64() * (max - min)) + min
}

func randomSphereInScene(pScene Scene) Sphere {
	s := randomSphere()

	for pScene.SDF(s.Position) <= s.Radius {
		s = randomSphere()
	}

	return s
}

func randomSphere() Sphere {
	return sphere(vector3(random(0, 20), random(30, 50), random(0, 20)), random(2, 8))
}
