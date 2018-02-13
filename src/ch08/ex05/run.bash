#!/bin/bash

go build mandelbrot.go
./mandelbrot -case usual > outUsual.png
./mandelbrot -case 10 > out10.png
./mandelbrot -case 50 > out50.png
./mandelbrot -case single > outSingle.png
./mandelbrot -case wg > outWg.png
