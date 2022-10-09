package main

import (
	"math/rand"
	"time"
)

const Epsilon = 0.01

func main() {
	rand.Seed(time.Now().UnixNano())

	c := camera(vector3(10, 0, 10), vector3(0, 0, 0), 50)

	s := scene()

	s.Add(randomSphereInScene(s))
	s.Add(randomSphereInScene(s))
	s.Add(randomSphereInScene(s))
	s.Add(randomSphereInScene(s))
	s.Add(randomSphereInScene(s))
	s.Add(randomSphereInScene(s))

	img := createImage(1280, 720)

	c.Render(img, s, light(vector3(20, 0, 20), 5.0))

	err := saveImage(img, "test.png")
	if err != nil {
		panic(err)
	}
}
