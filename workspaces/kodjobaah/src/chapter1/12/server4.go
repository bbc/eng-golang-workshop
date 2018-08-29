// Generates GIF animations of random Lissajous figures. via a browser. It accepts cycles as a query parameter
package main

import (
	"log"
	"net/http"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"strconv"
	"fmt"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	cycles, ok := r.URL.Query()["cycles"]
	cyc := 5
	if (ok) {
		c, err := strconv.Atoi(cycles[0])
		if err != nil {
			fmt.Printf("errr %s", err)
		}
		cyc = c
	} 
	lissajous(w, cyc)

}

func lissajous(out io.Writer, cycles int) {
	const(
		res     = 0.001    // angular resolution
		size    = 1000     // image canvas convers [-size..+size]
		nframes = 64       // number of animation frames
		delay   = 8        // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img  := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),blackIndex)
		}
		phase += 0.1
		anim.Delay  = append(anim.Delay, delay)
		anim.Image  = append(anim.Image, img)

	}
	gif.EncodeAll(out, &anim)  // NOTE: ignoring encoding errors

}

