package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

type Parameters struct {
	xmin, ymin, xmax, ymax float64
	x, y                   int
	magnification          float64
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}

func setParameters(r *http.Request) *Parameters {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	typeParams := Parameters{}
	typeParams.xmin = -2
	typeParams.xmax = 2
	typeParams.ymin = -2
	typeParams.ymax = 2
	for k, v := range r.Form {
		switch k {
		case "x":
			x, err := strconv.Atoi(v[0])
			if err != nil {
				panic(err)
			}
			typeParams.x = x
		case "y":
			y, err := strconv.Atoi(v[0])
			if err != nil {
				panic(err)
			}
			typeParams.y = y
		case "magnification":
			magnification, err := strconv.ParseFloat(v[0], 64)
			if err != nil {
				panic(err)
			}
			typeParams.xmin *= 1 / magnification
			typeParams.xmax *= 1 / magnification
			typeParams.ymin *= 1 / magnification
			typeParams.ymax *= 1 / magnification
		}
	}
	return &typeParams
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	const (
		width, height = 1024, 1024
	)
	parameters := setParameters(r)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py+parameters.y/2)/height*(parameters.ymax-parameters.ymin) + parameters.ymin
		for px := 0; px < width; px++ {
			x := float64(px+parameters.x/2)/width*(parameters.xmax-parameters.xmin) + parameters.xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(w, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{0, 255 - contrast*n, 128, 255 - contrast*n}
		}
	}
	return color.RGBA{123, 111, 255, 222}
}
