package main

import "math"

type Object interface {
	Distance(Vector3) float64
}

type Sphere struct {
	Position Vector3
	Radius   float64
}

type Cube struct {
	Position Vector3
	Size     Vector3
}

func sphere(pPosition Vector3, pRadius float64) Sphere {
	return Sphere{
		Position: pPosition,
		Radius:   pRadius,
	}
}

func (s Sphere) Distance(pPosition Vector3) float64 {
	return s.Position.Distance(pPosition) - s.Radius
}

func cube(pPosition, pSize Vector3) Cube {
	return Cube{
		Position: pPosition,
		Size:     pSize,
	}
}

func (c Cube) Distance(pPosition Vector3) float64 {
	x := math.Max(
		pPosition.X-c.Position.X-(c.Size.X/2),
		c.Position.X-pPosition.X-(c.Size.X/2),
	)

	y := math.Max(
		pPosition.Y-c.Position.Y-(c.Size.Y/2),
		c.Position.Y-pPosition.Y-(c.Size.Y/2),
	)

	z := math.Max(
		pPosition.Z-c.Position.Z-(c.Size.Z/2),
		c.Position.Z-pPosition.Z-(c.Size.Z/2),
	)

	return math.Max(math.Max(x, y), z)
}
