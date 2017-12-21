package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

var (
	width, height = 600, 320
	colorCode     = "00FFBF"
)

const (
	cells   = 100
	xyrange = 30.0
	angle   = math.Pi / 6
	red     = 0x00ff0000
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	setParameters(r)
	w.Write([]byte(surface()))
}

func setParameters(r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	for k, v := range r.Form {
		switch k {
		case "height":
			tmpH, err := strconv.Atoi(v[0])
			if err != nil {
				panic(err)
			}
			height = tmpH
		case "width":
			tmpW, err := strconv.Atoi(v[0])
			if err != nil {
				panic(err)
			}
			width = tmpW
		case "colorCode":
			colorCode = v[0]
		}
	}
}

func surface() string {
	var result string
	result = fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			if isFinite(ax) && isFinite(ay) &&
				isFinite(bx) && isFinite(by) &&
				isFinite(cx) && isFinite(cy) &&
				isFinite(dx) && isFinite(dy) {
				result += fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='#%s' />\n",
					ax, ay, bx, by, cx, cy, dx, dy, colorCode)
				result += fmt.Sprintf("color: %v\n", colorCode)
			}
		}
	}
	result += fmt.Sprintln("</svg>")
	return result
}

func isFinite(f float64) bool {
	return !math.IsInf(f, 0)
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	widthF := float64(width)
	heightF := float64(height)
	sx := widthF/2 + (x-y)*cos30*widthF/2/xyrange
	sy := heightF/2 + (x+y)*sin30*widthF/2/xyrange - z*heightF*0.4
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
