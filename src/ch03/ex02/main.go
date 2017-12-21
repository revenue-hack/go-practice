package main

import (
	"fmt"
	"math"
)

const (
	width, height = 400, 320
	cells = 2*math.Pi
	xyrange = 0.5
	xyscale = width / 10 / xyrange
	zscale = height * 0.4
	angle = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)


func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := float64(0); i < cells; i+=0.1 {
		for j := float64(0); j < cells; j+=0.1 {
			ax, ay := corner(i+0.1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+0.1)
			dx, dy := corner(i+0.1, j+0.1)
			if !isInf(ax, ay) &&
				!isInf(bx, by) &&
				!isInf(cx, cy) &&
				!isInf(dx, dy) {
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	fmt.Println("</svg>")
}

func isInf(value1 float64, value2 float64) bool {
	return math.IsInf(value1, 0) || math.IsInf(value2, 0)
}

func corner(i, j float64) (float64, float64) {
	x := math.Cos(float64(i))*math.Cos(float64(j))
	y := math.Cos(float64(i))*math.Sin(float64(j))
	z := f(float64(i))

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x float64) float64 {
	return math.Sin(x)
}


