// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

// Package weightconv performs weight conversions.
package weightconv

import "fmt"

// Kilo weight type
type Kilo float64

// Stone weight type
type Stone float64

// Pound weight type
type Pound float64

// Conversion constants
const (
	KiloP Kilo = 2.2046226218
	KiloS Kilo = 6.35029318

	PoundK Pound = 0.453592
	PoundS Pound = 0.0714286

	StoneK Stone = 0.15747
	StoneP Stone = 14
)

// Adds custom String() formatting to Kilo
func (f Kilo) String() string { return fmt.Sprintf("%gkg", f) }

// Adds custom String() formatting to Stone
func (f Stone) String() string { return fmt.Sprintf("%gst", f) }

// Adds custom String() formatting to Pound
func (f Pound) String() string { return fmt.Sprintf("%glb", f) }

//!-
