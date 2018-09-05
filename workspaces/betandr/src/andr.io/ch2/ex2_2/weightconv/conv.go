// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 41.

//!+

package weightconv

// KToP converts a Kilo weight to Pounds.
func KToP(k Kilo) Pound { return Pound(k * KiloP) }

// KToS converts a Kilo temperature to Stone.
func KToS(k Kilo) Stone { return Stone(k / KiloS) }

// PToK converts a Pound weight to Kilo.
func PToK(p Pound) Kilo { return Kilo(p * PoundK) }

// PToS converts a Pound temperature to Stone.
func PToS(p Pound) Stone { return Stone(p * PoundS) }

// SToK converts a Stone weight to Kilo.
func SToK(s Stone) Kilo { return Kilo(s / StoneK) }

// SToP converts a Stone temperature to Pound.
func SToP(s Stone) Pound { return Pound(s * StoneP) }

//!-
