// Package weightconv performs pounds and kilograms conversions.
package weightconv

import "fmt"

type Pound float64
type Kilogram float64

func (p Pound) String() string { return fmt.Sprintf("%glb", p) }
func (kg Kilogram) String() string { return fmt.Sprintf("%gkg", kg) }
