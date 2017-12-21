package main

import (
	"image"
	"image/color"
	"image/png"
	"math/big"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax float64 = -2, -2, +2, +2
		width, height                  = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height*(ymax-ymin)) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width*(xmax-xmin)) + xmin
			z := complex128(complex(x, y))
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	var vr float64
	var vi float64
	rz := real(z)
	iz := imag(z)
	for n := uint8(0); n < iterations; n++ {
		vr = real(v*v) + rz
		vi = imag(v*v) + iz
		v = complex128(complex(vr, vi))
		if cmplx.Abs(v) > 2 {
			return color.RGBA{0, 255 - contrast*n, 128, 255 - contrast*n}
		}
	}
	return color.RGBA{123, 111, 255, 222}
}
