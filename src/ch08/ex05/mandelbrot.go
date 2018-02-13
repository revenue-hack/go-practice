package main

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"
	"sync"
	"time"
)

func main() {
	caseF := flag.String("case", "usual", "case type")
	flag.Parse()
	switch *caseF {
	case "usual":
		caseUsual()
	case "10":
		caseCh10()
	case "50":
		caseCh50()
	case "single":
		caseChSingle()
	case "wg":
		caseWg()
	}
}

func caseUsual() {
	start := time.Now()
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
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img)
	log.Printf("case usual done: %s\n", time.Since(start))
}

func caseWg() {
	start := time.Now()
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	var wg sync.WaitGroup
	for py := 0; py < height; py++ {
		go func(py int) {
			wg.Add(1)
			y := float64(py)/height*(ymax-ymin) + ymin
			for px := 0; px < width; px++ {
				x := float64(px)/width*(xmax-xmin) + xmin
				z := complex(x, y)
				img.Set(px, py, mandelbrot(z))
			}
			wg.Done()
		}(py)
	}
	wg.Wait()
	png.Encode(os.Stdout, img)
	log.Printf("case waitGroup done: %s\n", time.Since(start))
}

func caseCh10() {
	start := time.Now()
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	ch := make(chan struct{}, 10)
	for py := 0; py < height; py++ {
		go func(py int) {
			y := float64(py)/height*(ymax-ymin) + ymin
			for px := 0; px < width; px++ {
				x := float64(px)/width*(xmax-xmin) + xmin
				z := complex(x, y)
				img.Set(px, py, mandelbrot(z))
			}
			ch <- struct{}{}
		}(py)
	}
	for py := 0; py < height; py++ {
		<-ch
	}
	png.Encode(os.Stdout, img)
	log.Printf("case ch10 done: %s\n", time.Since(start))
}

func caseCh50() {
	start := time.Now()
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	ch := make(chan struct{}, 50)
	for py := 0; py < height; py++ {
		go func(py int) {
			y := float64(py)/height*(ymax-ymin) + ymin
			for px := 0; px < width; px++ {
				x := float64(px)/width*(xmax-xmin) + xmin
				z := complex(x, y)
				img.Set(px, py, mandelbrot(z))
			}
			ch <- struct{}{}
		}(py)
	}
	for py := 0; py < height; py++ {
		<-ch
	}
	png.Encode(os.Stdout, img)
	log.Printf("case ch50 done: %s\n", time.Since(start))
}

func caseChSingle() {
	start := time.Now()
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	ch := make(chan struct{})
	for py := 0; py < height; py++ {
		go func(py int) {
			y := float64(py)/height*(ymax-ymin) + ymin
			for px := 0; px < width; px++ {
				x := float64(px)/width*(xmax-xmin) + xmin
				z := complex(x, y)
				img.Set(px, py, mandelbrot(z))
			}
			ch <- struct{}{}
		}(py)
	}
	for py := 0; py < height; py++ {
		<-ch
	}
	png.Encode(os.Stdout, img)
	log.Printf("case ch single done: %s\n", time.Since(start))
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
