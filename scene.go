package main

import "math"

type Scene struct {
	Objects []Object
}

func scene() Scene {
	return Scene{
		Objects: make([]Object, 0),
	}
}

func (s *Scene) Add(pObject Object) {
	s.Objects = append(s.Objects, pObject)
}

// SDF Signed Distance Function for our scene. (Gives the minimum distance to any object in the scene for a point)
func (s *Scene) SDF(pPosition Vector3) float64 {
	minimum := math.MaxFloat64

	for _, object := range s.Objects {
		distance := object.Distance(pPosition)

		if distance < minimum {
			minimum = distance
		}
	}

	return minimum
}

func (s *Scene) EstimateNormal(pPosition Vector3, pEpsilon float64) Vector3 {
	return vector3(
		s.SDF(vector3(pPosition.X+pEpsilon, pPosition.Y, pPosition.Z))-s.SDF(vector3(pPosition.X-pEpsilon, pPosition.Y, pPosition.Z)),
		s.SDF(vector3(pPosition.X, pPosition.Y+pEpsilon, pPosition.Z))-s.SDF(vector3(pPosition.X, pPosition.Y-pEpsilon, pPosition.Z)),
		s.SDF(vector3(pPosition.X, pPosition.Y, pPosition.Z+pEpsilon))-s.SDF(vector3(pPosition.X, pPosition.Y, pPosition.Z-pEpsilon)),
	).Normalize()
}
