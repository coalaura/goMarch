package main

import (
	"image"
	"image/color"
)

type Camera struct {
	Position Vector3
	Rotation Vector3
	FOV      float64
}

// Rotation:
//   X: Pitch
//   Y: Roll
//   Z: Yaw
// Position:
//   X: Left & Right
//   Y: Forward & Backward
//   Z: Up & Down
func camera(pPosition Vector3, pRotation Vector3, FOV float64) Camera {
	return Camera{
		Position: pPosition,
		Rotation: pRotation,
		FOV:      FOV,
	}
}

func (c Camera) Render(pImage *image.RGBA, pScene Scene, pLight Light) {
	startDistance := pScene.SDF(c.Position)

	width := float64(pImage.Rect.Max.X)
	height := float64(pImage.Rect.Max.Y)

	yFOV := c.FOV
	xFOV := c.FOV

	if width > height {
		xFOV *= width / height
	} else if width < height {
		yFOV *= height / width
	}

	yHalfFOV := yFOV / 2
	xHalfFOV := xFOV / 2

	for y := float64(0); y < height; y++ {
		for x := float64(0); x < width; x++ {
			direction := vector3(c.Rotation.X+((x/width)*xFOV-xHalfFOV), -1.0, c.Rotation.Z+((y/height)*yFOV-yHalfFOV))

			hit, _, marchedPosition := march(c.Position, direction, startDistance, Epsilon, pScene.SDF)

			if hit {
				// Calculate the diffuse lighting
				lightVector := pLight.Position.Sub(marchedPosition).Normalize()
				normalVector := pScene.EstimateNormal(marchedPosition, Epsilon)

				diffuse := normalVector.Dot(lightVector)
				if diffuse < 0 {
					diffuse = 0
				}

				c := uint8(diffuse * 255)
				pImage.Set(int(x), int(y), color.RGBA{R: c, G: c, B: c, A: 255})
			} else {
				pImage.Set(int(x), int(y), color.Black)
			}
		}
	}
}

func march(pPosition, pDirection Vector3, pStartDistance, pEpsilon float64, pSDF func(Vector3) float64) (bool, float64, Vector3) {
	vecY := EulerToForwardMatrix(pDirection.X, pDirection.Y, pDirection.Z)

	marchedDistance := pStartDistance
	marchedPosition := pPosition.Add(vecY.Scale(pStartDistance))

	for marchedDistance < 100 {
		distance := pSDF(marchedPosition)

		marchedDistance += distance

		if distance <= pEpsilon {
			return true, marchedDistance, marchedPosition
		}

		marchedPosition = marchedPosition.Add(vecY.Scale(distance))
	}

	return false, marchedDistance, marchedPosition
}
