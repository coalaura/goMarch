package main

type Light struct {
	Position     Vector3
	LightFalloff float64
}

func light(pPosition Vector3, pLightFalloff float64) Light {
	return Light{
		Position:     pPosition,
		LightFalloff: pLightFalloff,
	}
}
