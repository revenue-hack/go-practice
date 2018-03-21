#!/bin/bash
go build gopl.io/ch3/mandelbrot
go build image.go
./mandelbrot | ./image -type png > mandelbrot.png
./mandelbrot | ./image -type jpeg > mandelbrot.jpg
./mandelbrot | ./image -type gif > mandelbrot.gif
