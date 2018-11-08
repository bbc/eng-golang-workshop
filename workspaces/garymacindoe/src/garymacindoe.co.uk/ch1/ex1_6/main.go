// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Run with "web" command-line argument for web server.
// See page 13.
//!+main

// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

//!-main
// Packages not needed by version in book.
import (
	"log"
	"net/http"
	"time"
)

//!+main

func main() {
	//!-main
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe(":8000", nil))
		return
	}
	//!+main
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	var palette = make([]color.Color, 255, 255)
	palette[0] = color.RGBA{0x00, 0x00, 0x00, 0xff}
	for i:= 1; i < 255; i++ {
	    palette[i] = hslToRgb(float64(i) / 255.0, 1.0, 0.5)
	}
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				Round((y + 1.0) * 127.0))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
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

//!-main
