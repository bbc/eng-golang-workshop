package tempconv

// CToF converts a Celsius temperature to Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// CToK converts a Celsius temperature to Kelvin
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

// FToF converts a Fahrenheit temperature to Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// FToK converts a Fahrenheit temperature to Kelvin
func FToK(f Fahrenheit) Kelvin { return Kelvin((f + 459.67) * 5 / 9) }

// KToF converts Kelvins to Fahrenheit
func KToF(k Kelvin) Fahrenheit { return Fahrenheit(k*9/6 - 459.67) }

// KToC converts Kelvins to Celsius
func KToC(k Kelvin) Celsius { return Celsius((k - 273.15)) }
