// Package lengthconv performs Feet and Meters conversions.
package lengthconv

import "fmt"

type Foot float64
type Meter float64

func (f Foot) String() string { return fmt.Sprintf("%g'", f) }
func (m Meter) String() string { return fmt.Sprintf("%gm", m) }
