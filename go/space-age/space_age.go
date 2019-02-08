package space

// Planet contains just the name - learning to use types
type Planet string

// The number of seconds in an Earth year
var seconds float64 = 31557600

// The number of seconds in planets oribtal period vs. Earth
var planets = map[Planet]float64{
	"Earth":   1.0 * seconds,
	"Mercury": 0.2408467 * seconds,
	"Venus":   0.61519726 * seconds,
	"Mars":    1.8808158 * seconds,
	"Jupiter": 11.862615 * seconds,
	"Saturn":  29.447498 * seconds,
	"Uranus":  84.016846 * seconds,
	"Neptune": 164.79132 * seconds,
}

// Age when given an age in seconds calculate how many years
// old a person would be for a given planet based on
// the planets orbital period.
func Age(age float64, planet Planet) float64 {
	return age / planets[planet]
}
