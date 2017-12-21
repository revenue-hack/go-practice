package main

import (
	"fmt"
	"math"
	"strconv"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
	red           = 0x00ff0000
	blue          = "0x000000ff"
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j+1)
			dx, dy, dz := corner(i+1, j+1)
			if !isInf(ax, ay) &&
				!isInf(bx, by) &&
				!isInf(cx, cy) &&
				!isInf(dx, dy) {
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='#%s' />\n",
					ax, ay, bx, by, cx, cy, dx, dy, color(az, bz, cz, dz))
				fmt.Printf("color: %v\n", color(az, bz, cz, dz))
			}
		}
	}
	fmt.Println("</svg>")
}

func isInf(value1 float64, value2 float64) bool {
	return math.IsInf(value1, 0) || math.IsInf(value2, 0)
}

func color(az, bz, cz, dz float64) string {
	height := (az + bz + cz + dz) / 4
	maxH := 0.986
	minH := -0.218
	blue10, err := strconv.ParseInt(blue, 0, 64)
	if err != nil {
		panic(err)
	}

	ratio := uint32((maxH - height) / (maxH - minH) * float64(blue10))
	return fmt.Sprintf("%06x", (red-ratio<<16)+ratio)
}

func corner(i, j int) (float64, float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
