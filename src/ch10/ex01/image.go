package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

func main() {
	var f = flag.String("type", "png", "select png or jpeg or gif")
	flag.Parse()
	img, err := readImg(os.Stdin)
	if err != nil {
		fmt.Errorf("can't read image")
	}
	switch *f {
	case "png":
		toPNG(img, os.Stdout)
	case "jpeg":
		toJPEG(img, os.Stdout)
	case "gif":
		toGIF(img, os.Stdout)
	default:
		fmt.Errorf("select type")
	}

}

func toJPEG(img image.Image, out io.Writer) {
	jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}

func toPNG(img image.Image, out io.Writer) {
	png.Encode(out, img)
}

func toGIF(img image.Image, out io.Writer) {
	gif.Encode(out, img, &gif.Options{NumColors: 256, Quantizer: nil, Drawer: nil})
}

func readImg(in io.Reader) (image.Image, error) {
	img, kind, err := image.Decode(in)
	if err != nil {
		return nil, err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return img, nil
}
