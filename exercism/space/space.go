package space

// Planet planet type
type Planet string

// Age find the age of someone on other planets
func Age(s float64, planet Planet) float64 {
	earthSeconds := 31557600.00
	planetMap := map[Planet]float64{
		"Earth": 1,
		"Mercury":  0.2408467,
		"Venus": 0.61519726,
		"Mars": 1.8808158,
		"Jupiter": 11.862615,
		"Saturn": 29.447498,
		"Uranus": 84.016846,
		"Neptune": 164.79132,
	}

	earthAge := s/earthSeconds
	return earthAge/planetMap[planet]
}