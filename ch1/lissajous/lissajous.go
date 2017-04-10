// Package lissajous ...
package lissajous

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

var green = color.RGBA{0x00, 0x55, 0x00, 0xff}
var red = color.RGBA{0x55, 0x00, 0x00, 0xff}
var blue = color.RGBA{0x00, 0x00, 0x55, 0xff}
var black = color.Black
var white = color.White

var palette = []color.Color{green, red, blue, black, white}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

// LissaJous generates GIF animations of random Lissajous figures.
func LissaJous(out io.Writer, cycles int) {
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			var colorIndex uint8

			if i%4 == 0 {
				colorIndex = 4
			} else if i%3 == 0 {
				colorIndex = 3
			} else if i%2 == 0 {
				colorIndex = 2
			} else {
				colorIndex = 1
			}

			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
