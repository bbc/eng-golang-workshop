// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
    "image"
    "image/color"
    "image/png"
    "math"
    "math/cmplx"
    "os"
)

func f(z complex128) complex128 {
    return cmplx.Pow(z, 4) - 1
}

func df(z complex128) complex128 {
    return 4 * cmplx.Pow(z, 3)
}

func main() {
    const (
        xmin, ymin, xmax, ymax = -2, -2, +2, +2
        width, height          = 1024, 1024
    )

    img := image.NewRGBA(image.Rect(0, 0, width, height))
    for py := 0; py < height; py++ {
        y := float64(py)/height*(ymax-ymin) + ymin
        for px := 0; px < width; px++ {
            x := float64(px)/width*(xmax-xmin) + xmin
            z := complex(x, y)
            // Image point (px, py) represents complex value x.
            img.Set(px, py, newton(z))
        }
    }
    png.Encode(os.Stdout, img)  // NOTE: ignoring errors
}

func supersample4(xw, yw float64, fn func(complex128) color.RGBA, z complex128) color.RGBA {
    c0 := fn(complex(real(z) - 0.25 * xw, imag(z) - 0.25 * yw))
    c1 := fn(complex(real(z) - 0.25 * xw, imag(z) + 0.25 * yw))
    c2 := fn(complex(real(z) + 0.25 * xw, imag(z) - 0.25 * yw))
    c3 := fn(complex(real(z) + 0.25 * xw, imag(z) + 0.25 * yw))
    return color.RGBA{(c0.R + c1.R + c2.R + c3.R) / 4,
                      (c0.G + c1.G + c2.G + c3.G) / 4,
                      (c0.B + c1.B + c2.B + c3.B) / 4,
                      255 }
}

func mandelbrot(z complex128) color.RGBA {
    const iterations = 200
    const contrast = 15

    var v complex128
    for n := uint8(0); n < iterations; n++ {
        v = v*v + z
        if cmplx.Abs(v) > 2 {
            return hslToRgb(float64(255 - contrast*n)/255.0, 0.5, 0.5)
        }
    }
    return color.RGBA{0, 0, 0, 255}
}

func newton(z complex128) color.RGBA {
    const iterations = 37
    const contrast = 7
    const tolerance = 0.000001

    var roots = [...]complex128{
        complex( 1,  0),
        complex(-1,  0),
        complex( 0,  1),
        complex( 0, -1),
    }

    var colours = [...]color.RGBA{
        color.RGBA{255,   0,   0, 255},
        color.RGBA{  0, 255,   0, 255},
        color.RGBA{  0,   0, 255, 255},
        color.RGBA{255, 255,   0, 255},
    }

    zx := z
    for i := uint8(0); i < 4; i++ {
        z = zx
        for n := uint8(0); n < iterations; n++ {
            z = z - f(z) / df(z)
            if cmplx.Abs(z - roots[i]) < tolerance {
                return color.RGBA{colours[i].R, colours[i].G, colours[i].B, max(0, 255 - contrast*n)}
            }
        }
    }
    return color.RGBA{0, 0, 0, 255}
}

func max(x, y uint8) uint8 {
    if x > y {
        return x
    }
    return y
}

func hslToRgb(h float64, s float64, l float64) color.RGBA {
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

    return color.RGBA{Round(r * 255.0), Round(g * 255.0), Round(b * 255.0), 0xff}
}

func Round(x float64) uint8 {
    t := math.Trunc(x)
    if math.Abs(x-t) >= 0.5 {
        return uint8(t + math.Copysign(1, x))
    }
    return uint8(t)
}
