package lengthconv

// FToM converts a length in feet to meters.
func FToM(f Foot) Meter { return Meter(f * 0.3048) }

// MToF converts a length in meters to feet.
func MToF(m Meter) Foot { return Foot(m * 3.280839) }
