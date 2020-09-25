package space

// Planet represents a Planet
type Planet string

// Age is aged cured for Planet
func Age(seconds float64, planet Planet) float64 {
	planetRatioMap := map[Planet]float64{
		"Earth":   1.0,
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}
	return seconds / planetRatioMap[planet] / 31557600
}
