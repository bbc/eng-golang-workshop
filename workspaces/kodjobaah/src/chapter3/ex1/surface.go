// Surface computes an SVG rendering of a 3-D surface function
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            //Canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cose(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, ares := corner(i+1, j)
			bx, by, bres := corner(i, j)
			cx, cy, cres := corner(i, j+1)
			dx, dy, dres := corner(i+1, j+1)

			if ares && bres && cres && dres {

				fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' style='fill:#ff0000' />\n", ax, ay, bx, by, cx, cy, dx, dy)

				//fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' style='fill:#0000ff' />\n", ax, ay, bx, by, cx, cy, dx, dy)

			}
		}

	}
	fmt.Println("</svg>")
}

func All(vs []float64, f func(float64) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func corner(i, j int) (float64, float64, bool) {
	// Find point (x, y) at corner of cell (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	//	z := f(x, y)
	z := saddle(x, y)
	if math.IsInf(z, 0) {
		return 0, 0, false
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx, sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, true
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r

}

func saddle(x, y float64) float64 {
	return (x * x) - (y * y)
}
func eggFunction(x, y float64) float64 {

	sinx := math.Sin(x)
	siny := math.Sin(y)
	return (x * x) + (y * y) + 25*((sinx*sinx)+(siny*siny))

}
