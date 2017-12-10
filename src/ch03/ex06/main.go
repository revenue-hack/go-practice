package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
	deltaX, deltaY         = (xmax - xmin) / width, (ymax - ymin) / height
)

func main() {

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			img.Set(px, py, superSampling(x, y))
		}
	}
	png.Encode(os.Stdout, img)
}

func superSampling(x, y float64) color.Color {
	z1 := complex(x-deltaX, y-deltaY)
	z2 := complex(x+deltaX, y-deltaY)
	z3 := complex(x-deltaX, y+deltaY)
	z4 := complex(x+deltaX, y+deltaY)
	colors := []color.Color{mandelbrot(z1), mandelbrot(z2), mandelbrot(z3), mandelbrot(z4)}
	num := uint8(len(colors))
	var sumR, sumG, sumB, sumA uint32
	for _, color := range colors {
		r, g, b, a := color.RGBA()
		sumR += r
		sumG += g
		sumB += b
		sumA += a
	}
	return color.RGBA{uint8(sumR) / num, uint8(sumG) / num, uint8(sumB) / num, uint8(sumA) / num}
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
