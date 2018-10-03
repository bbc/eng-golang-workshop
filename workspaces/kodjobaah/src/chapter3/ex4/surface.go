// Server2 is a minimal "echo" and counter server
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
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
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "image/svg+xml")
	header := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	w.Write([]byte(header))
	var poly = ""
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, ares := corner(i+1, j)
			bx, by, bres := corner(i, j)
			cx, cy, cres := corner(i, j+1)
			dx, dy, dres := corner(i+1, j+1)

			if ares && bres && cres && dres {
				poly = fmt.Sprintf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' style='stroke:#ff0000' />\n", ax, ay, bx, by, cx, cy, dx, dy)
				w.Write([]byte(poly))

			}
		}

	}
	footer := fmt.Sprintf("</svg>")
	w.Write([]byte(footer))
}

func corner(i, j int) (float64, float64, bool) {
	// Find point (x, y) at corner of cell (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

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
