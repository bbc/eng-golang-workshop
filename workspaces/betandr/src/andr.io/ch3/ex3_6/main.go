// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax  = -2, -2, +2, +2
		width, height           = 1024, 1024
		superwidth, superheight = width * 2, height * 2
	)

	var pixels [superwidth][superheight]color.Color

	for py := 0; py < superheight; py++ {
		y := float64(py)/superheight*(ymax-ymin) + ymin
		for px := 0; px < superwidth; px++ {
			x := float64(px)/superwidth*(xmax-xmin) + xmin
			z := complex(x, y)
			pixels[px][py] = mandelbrot(z)
		}
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {

			x2, y2 := 2*x, 2*y

			// get four surrounding pixels
			r1, g1, b1, a1 := pixels[x2][y2].RGBA()
			r2, g2, b2, a2 := pixels[x2+1][y2].RGBA()
			r3, g3, b3, a3 := pixels[x2][y2+1].RGBA()
			r4, g4, b4, a4 := pixels[x2+1][y2+1].RGBA()

			// average the colour
			averaged := color.RGBA{
				uint8((r1 + r2 + r3 + r4) / 1028),
				uint8((g1 + g2 + g3 + g4) / 1028),
				uint8((b1 + b2 + b3 + b4) / 1028),
				uint8((a1 + a2 + a3 + a4) / 1028)}

			img.Set(x, y, averaged)
		}
	}

	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return palette.Plan9[255-contrast*n]
		}
	}
	return color.Black
}

//!-

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return palette.Plan9[255-contrast*i]
		}
	}
	return color.Black
}
