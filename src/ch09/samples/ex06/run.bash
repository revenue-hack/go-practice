#!/bin/bash

go build mandelbrot.go
echo "usual"
export GOMAXPROCS=1
echo "GOMAXPROCS=${GOMAXPROCS}"
./mandelbrot -case usual > outUsual.png
echo ""
export GOMAXPROCS=2
echo "GOMAXPROCS=${GOMAXPROCS}"
./mandelbrot -case usual > outUsual.png
echo ""
export GOMAXPROCS=3
echo "GOMAXPROCS=${GOMAXPROCS}"
./mandelbrot -case usual > outUsual.png
echo ""
export GOMAXPROCS=4
echo "GOMAXPROCS=${GOMAXPROCS}"
./mandelbrot -case usual > outUsual.png

echo ""
echo "goroutine 10"
export GOMAXPROCS=1
echo "GOMAXPROCS=${GOMAXPROCS}"
./mandelbrot -case 10 > out10.png
echo ""
export GOMAXPROCS=2
echo "GOMAXPROCS=${GOMAXPROCS}"
./mandelbrot -case 10 > out10.png
echo ""
export GOMAXPROCS=3
echo "GOMAXPROCS=${GOMAXPROCS}"
./mandelbrot -case 10 > out10.png
echo ""
export GOMAXPROCS=4
echo "GOMAXPROCS=${GOMAXPROCS}"
./mandelbrot -case 10 > out10.png

echo ""
echo "goroutine 50"
export GOMAXPROCS=1
echo "GOMAXPROCS=${GOMAXPROCS}"
./mandelbrot -case 50 > out50.png
echo ""
export GOMAXPROCS=2
echo "GOMAXPROCS=${GOMAXPROCS}"
./mandelbrot -case 50 > out50.png
echo ""
export GOMAXPROCS=3
echo "GOMAXPROCS=${GOMAXPROCS}"
./mandelbrot -case 50 > out50.png
echo ""
export GOMAXPROCS=4
echo "GOMAXPROCS=${GOMAXPROCS}"
./mandelbrot -case 50 > out50.png

echo ""
echo "goroutine single"
export GOMAXPROCS=1
echo "GOMAXPROCS=${GOMAXPROCS}"
./mandelbrot -case single > outSingle.png
echo ""
export GOMAXPROCS=2
echo "GOMAXPROCS=${GOMAXPROCS}"
./mandelbrot -case single > outSingle.png
echo ""
export GOMAXPROCS=3
echo "GOMAXPROCS=${GOMAXPROCS}"
./mandelbrot -case single > outSingle.png
echo ""
export GOMAXPROCS=4
echo "GOMAXPROCS=${GOMAXPROCS}"
./mandelbrot -case single > outSingle.png

echo ""
echo "sync waitgroup"
export GOMAXPROCS=1
echo "GOMAXPROCS=${GOMAXPROCS}"
./mandelbrot -case wg > outWg.png
echo ""
export GOMAXPROCS=2
echo "GOMAXPROCS=${GOMAXPROCS}"
./mandelbrot -case wg > outWg.png
echo ""
export GOMAXPROCS=3
echo "GOMAXPROCS=${GOMAXPROCS}"
./mandelbrot -case wg > outWg.png
echo ""
export GOMAXPROCS=4
echo "GOMAXPROCS=${GOMAXPROCS}"
./mandelbrot -case wg > outWg.png
