// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "image/svg+xml")
			surface(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe(":8000", nil))
		return
	}
	surface(os.Stdout)
}

func surface(out io.Writer) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, erra := corner(i+1, j)
			bx, by, errb := corner(i, j)
			cx, cy, errc := corner(i, j+1)
			dx, dy, errd := corner(i+1, j+1)
			if erra == nil && errb == nil && errc == nil && errd == nil {
				x := xyrange * (float64(i)/cells - 0.5)
				y := xyrange * (float64(j)/cells - 0.5)
				z := f(x, y)
				r, g, b := hslToRgb(z, 0.5, 0.5)
				fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='#%X%X%X'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, r, g, b)
			}
		}
	}
	fmt.Fprintln(out, "</svg>")
}

func corner(i, j int) (float64, float64, error) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)
	if math.IsInf(z, 0) {
		return 0.0, 0.0, errors.New("function returned +/- Infinity")
        }

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, nil
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

// Egg box - http://benchmarkfcns.xyz/benchmarkfcns/eggcratefcn.html
func g(x, y float64) float64 {
	r := x * x + y * y + 25 * (math.Sin(x) * math.Sin(x) + math.Sin(y) * math.Sin(y))
	return r * 0.001
}

// Moguls
func h(x, y float64) float64 {
	r := g(x, y)
	return math.Max(r, 0.0)
}

// Saddle plot
func k(x, y float64) float64 {
	var a, b, c = 25.0, 25.0, 1.0
	z := (x * x) / (a * a) - (y * y) / (b * b)
	return z * c
}

func hslToRgb(h float64, s float64, l float64) (uint8, uint8, uint8) {
    var r, g, b float64

    if s == 0 {
        r = l
        g = l
        b = l
    } else {
        var hue2rgb = func(p float64, q float64, t float64) float64 {
            if t < 0.0 { t += 1.0 }
            if t > 1.0 { t -= 1.0 }
            if t < 1.0/6.0 { return p + (q - p) * 6.0 * t }
            if t < 1.0/2.0 { return q }
            if t < 2.0/3.0 { return p + (q - p) * (2.0/3.0 - t) * 6.0 }
            return p
        }
        var q float64
        if l < 0.5 {
            q = l * (1.0 + s)
        } else {
            q = l + s - l * s
        }
        var p = 2.0 * l - q
        r = hue2rgb(p, q, h + 1.0/3.0)
        g = hue2rgb(p, q, h)
        b = hue2rgb(p, q, h - 1.0/3.0)
    }

    return Round(r * 255.0), Round(g * 255.0), Round(b * 255.0)
}

func Round(x float64) uint8 {
    t := math.Trunc(x)
    if math.Abs(x-t) >= 0.5 {
        return uint8(t + math.Copysign(1, x))
    }
    return uint8(t)
}

//!-
