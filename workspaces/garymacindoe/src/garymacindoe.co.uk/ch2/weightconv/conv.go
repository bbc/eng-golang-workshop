package weightconv

// PToK converts a weight in pounds to kilograms.
func PToK(p Pound) Kilogram { return Kilogram(p * 0.453592) }

// KToP converts a weight in kilograms to pounds.
func KToP(kg Kilogram) Pound { return Pound(kg * 2.204622) }
