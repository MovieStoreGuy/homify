package data

import "fmt"

type (
	Unit int

	suffix struct {
		name   string
		suffix string
	}
)

const (
	_ Unit = iota
	// Time Units
	StatisticNanosecond
	StatisticMicrosecond
	StatisticMillisecond
	StatisticSecond
	StatisticMinute
	StatisticHour
	// Length Units
	StatisticMillimetre
	StatisticCentimetre
	StatisticMetre
	StatisticKilometre
	// Mass Units
	StatisticMilligram
	StatisticGrams
	StatisticKilograms
	// Electric Currents
	StatisticAmpere
	StatisticWatts
	StatisticKiloWatts
	// Temperature
	StatisticKelvin
	StatisticCelsius
	StatisticFahrenheit
	// Luminous intensity
	StatisticCandela
	StatisticLuminousFlux
	StatisticLuminance
	// Various Units
	StatisticCount
)

func (s suffix) String() string {
	return fmt.Sprintf("%s [%s]", s.name, s.suffix)
}

var units = [...]suffix{
	{name: "unknown", suffix: ""},
	{name: "nanosecond", suffix: "ns"},
	{name: "microsecond", suffix: "µs"},
	{name: "millisecond", suffix: "ms"},
	{name: "second", suffix: "s"},
	{name: "minute", suffix: "m"},
	{name: "hour", suffix: "h"},
	{name: "millimetre", suffix: "mm"},
	{name: "centimetre", suffix: "cm"},
	{name: "metre", suffix: "m"},
	{name: "kilometre", suffix: "km"},
	{name: "milligrams", suffix: "mg"},
	{name: "grams", suffix: "g"},
	{name: "kilograms", suffix: "g"},
	{name: "ampere", suffix: "amp"},
	{name: "watt", suffix: "w"},
	{name: "kilowatt", suffix: "kw"},
	{name: "kelvin", suffix: "K"},
	{name: "celsius", suffix: "°C"},
	{name: "fahrenheit", suffix: "°F"},
	{name: "candela", suffix: "cd"},
	{name: "luminous flux", suffix: "lm"},
	{name: "luminance", suffix: "cd/m2"},
	{name: "count", suffix: ""},
}

func (s Unit) String() string {
	if i := int(s); i < len(units) {
		return units[i].String()
	}
	return "unknown"
}
