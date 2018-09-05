// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

// Package tempconv performs temperature conversions.
package tempconv

import "fmt"

// Celsius temperature type
type Celsius float64

// Fahrenheit temperature type
type Fahrenheit float64

// Kelvin temperature type
type Kelvin float64

// Conversion constants (some could be more generic to support two types)
const (
	AbsoluteZeroC Celsius    = -273.15
	FreezingC     Celsius    = 0
	BoilingC      Celsius    = 100
	CelsiusK      Celsius    = 273.15
	KelvinC       Kelvin     = 273.15
	KelvinF       Kelvin     = 459.67
	FahrenheitK   Fahrenheit = 459.67
)

// Adds custom String() formatting to Celsius
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

// Adds custom String() formatting to Fahrenheit
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

// Adds custom String() formatting to Kelvin
func (f Kelvin) String() string { return fmt.Sprintf("%g K", f) }

//!-
