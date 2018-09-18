// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math"
	"math/cmplx"
	"net/http"
	"strconv"
	"strings"
)

func main() {

	//!+http
	handler := func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		xMin, xMinErr := strconv.Atoi(strings.Join(r.Form["xmin"], ""))
		if xMinErr != nil {
			xMin = -2
		}
		yMin, yMinErr := strconv.Atoi(strings.Join(r.Form["xmax"], ""))
		if yMinErr != nil {
			yMin = -2
		}
		xMax, xMaxErr := strconv.Atoi(strings.Join(r.Form["ymin"], ""))
		if xMaxErr != nil {
			xMax = +2
		}
		yMax, yMaxErr := strconv.Atoi(strings.Join(r.Form["ymax"], ""))
		if yMaxErr != nil {
			yMax = +2
		}
		z, zErr := strconv.Atoi(strings.Join(r.Form["zoom"], ""))
		if zErr != nil {
			z = 1
		}

		draw(w, xMin, yMin, xMax, yMax, z)
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return
}

func draw(w io.Writer, xMin int, yMin int, xMax int, yMax int, zoom int) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	// TODO: not finished yet...

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(w, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + complex64(z)
		if cmplx.Abs(complex128(v)) > 2 {
			switch {
			case n > 50:
				return color.RGBA{0, 255, 0, 255}
			default:
				logScale := math.Log(float64(n)) / math.Log(float64(iterations))
				return color.RGBA{0, 0, 255 - uint8(logScale*255), 255}
			}
		}
	}
	return color.Black
}
