package main

import(
  "image"
  "image/color"
  "image/gif"
  "math"
  "math/rand"
  "log"
  "net/http"
  "io/ioutil"
  "os"
  "strconv"
)

var palette = []color.Color{color.Black, color.RGBA{0x0, 0xff, 0x0, 0xff}}

const (
  blackIndex = 0
  greenIndex = 1
)

func main() {
  http.HandleFunc("/", handler)
  log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
  cycles := getParameter(w, r)
  lissajous(w, cycles)
}

func getParameter(w http.ResponseWriter, r *http.Request) int {
  if err := r.ParseForm(); err != nil {
    panic(err)
  }
  var cycles = 5
  for k, v := range r.Form {
    if k == "cycles" {
      paramCycles, err := strconv.Atoi(v[0])
      if err != nil {
        panic(err)
      }
      cycles = paramCycles
    }
  }
  return cycles
}

func lissajous(w http.ResponseWriter, paramCycles int) {
  const (
    res = 0.001
    size = 100
    nframes = 64
    delay = 8
  )
  var cycles = float64(paramCycles)
  freq := rand.Float64() * 3.0
  anim := gif.GIF{LoopCount: nframes}
  phase := 0.0
  for i := 0; i < nframes; i++ {
    rect := image.Rect(0, 0, 2 * size + 1, 2 * size + 1)
    img := image.NewPaletted(rect, palette)
    for t := 0.0; t < cycles * math.Pi; t+= res {
      x := math.Sin(t)
      y := math.Sin(t * freq + phase)
      img.SetColorIndex(size + int(x * size + 0.5), size + int(y * size + 0.5), greenIndex)
    }
    phase += 0.1
    anim.Delay = append(anim.Delay, delay)
    anim.Image = append(anim.Image, img)
  }
  file, err := os.OpenFile("out.gif", os.O_WRONLY|os.O_CREATE, 0600)
  if err != nil {
    panic(err)
  }
  defer file.Close()
  gif.EncodeAll(file, &anim)
  data, err := ioutil.ReadFile("out.gif")
  if err != nil {
    panic(err)
  }
  w.Write([]byte(string(data)))
}

