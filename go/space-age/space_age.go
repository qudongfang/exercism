package space

// Planet planet type
type Planet string

const (
	// Mecury planet Mecury
	Mecury Planet = "Mercury"

	// Venus planet Venus
	Venus Planet = "Venus"

	// Earth planet Earch
	Earth Planet = "Earth"

	// Mars planet Mars
	Mars Planet = "Mars"

	// Jupiter planet Jupiter
	Jupiter Planet = "Jupiter"

	// Saturn planet Saturn
	Saturn Planet = "Saturn"

	// Uranus planet Uranus
	Uranus Planet = "Uranus"

	// Neptune planet Neptune
	Neptune Planet = "Neptune"
)

// SecPerYearOnEarth seconds per year on Earth
const SecPerYearOnEarth = 31557600

// hash to save times of year to Earth year
var hash = map[Planet]float64{
	Mecury:  0.2408467,
	Venus:   0.61519726,
	Earth:   1,
	Mars:    1.8808158,
	Jupiter: 11.862615,
	Saturn:  29.447498,
	Uranus:  84.016846,
	Neptune: 164.79132,
}

// Age calculate how old someone would be
func Age(sec float64, planet Planet) float64 {
	if times, ok := hash[planet]; ok {
		return sec / float64(SecPerYearOnEarth) / times
	}

	panic("invalid planet" + planet)
}
