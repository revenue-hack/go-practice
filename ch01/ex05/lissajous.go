package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.Black, color.RGBA{0x0, 0xff, 0x0, 0xff}}

const (
	blackIndex = 0
	greenIndex = 1
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(calcRect(size))
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*math.Pi; t += res {
			x, y := calcXY(t, freq, phase)
			img.SetColorIndex(calcColorIndex(x, y, size))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func calcColorIndex(x float64, y float64, size int) (int, int, uint8) {
	return size + int(x*float64(size)+0.5), size + int(y*float64(size)+0.5), greenIndex
}

func calcXY(t float64, freq float64, phase float64) (float64, float64) {
	return math.Sin(t), math.Sin(t*freq + phase)
}

func calcRect(size int) (int, int, int, int) {
	return 0, 0, 2*size + 1, 2*size + 1
}
