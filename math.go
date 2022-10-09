package main

import "math"

type Vector3 struct {
	X float64
	Y float64
	Z float64
}

func EulerToForwardMatrix(pRotationX, pRotationY, pRotationZ float64) Vector3 {
	radX := Radians(pRotationX)
	radY := Radians(pRotationY)
	radZ := Radians(pRotationZ)

	sinX := math.Sin(radX)
	sinY := math.Sin(radY)
	sinZ := math.Sin(radZ)

	cosX := math.Cos(radX)
	cosY := math.Cos(radY)
	cosZ := math.Cos(radZ)

	return vector3(cosZ*sinX*sinY-cosX*sinZ, cosX*cosZ-sinX*sinY*sinZ, cosY*sinX)
}

func Radians(pDegrees float64) float64 {
	return pDegrees * (math.Pi / 180)
}

func vector3(pX, pY, pZ float64) Vector3 {
	return Vector3{
		X: pX,
		Y: pY,
		Z: pZ,
	}
}

func (v Vector3) Distance(pVector Vector3) float64 {
	return math.Abs(math.Sqrt(math.Pow(v.X-pVector.X, 2) + math.Pow(v.Y-pVector.Y, 2) + math.Pow(v.Z-pVector.Z, 2)))
}

func (v Vector3) Add(pVector Vector3) Vector3 {
	return vector3(v.X+pVector.X, v.Y+pVector.Y, v.Z+pVector.Z)
}

func (v Vector3) Sub(pVector Vector3) Vector3 {
	return vector3(v.X-pVector.X, v.Y-pVector.Y, v.Z-pVector.Z)
}

func (v Vector3) Scale(pFactor float64) Vector3 {
	return vector3(v.X*pFactor, v.Y*pFactor, v.Z*pFactor)
}

func (v Vector3) Dot(pVector Vector3) float64 {
	return v.X*pVector.X + v.Y*pVector.Y + v.Z*pVector.Z
}

func (v Vector3) Length() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2) + math.Pow(v.Z, 2))
}

func (v Vector3) Normalize() Vector3 {
	length := v.Length()

	return vector3(v.X/length, v.Y/length, v.Z/length)
}
