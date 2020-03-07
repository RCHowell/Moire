package main

import (
	"flag"
	"image"
	"image/color"
	"image/gif"
	"math"
	"os"
)

const (
	rate  = 8
	size  = 128
	delay = 10
)

var cmyk = []color.Color{
	color.White, // background
	color.CMYK{C: math.MaxUint8},
	color.CMYK{M: math.MaxUint8},
	color.CMYK{Y: math.MaxUint8},
	color.CMYK{K: math.MaxUint8},
}

var frames = rate * 360

// https://i.imgur.com/k86QK.png
var outerSquareSize = int(math.Ceil(size * 2 / math.Sqrt2))
var radius = size / 2
var center = (outerSquareSize / 2) + 1

func main() {
	gaps := flag.Int("gaps", 4, "gap size between pixels")
	flag.Parse()

	anim := gif.GIF{LoopCount: frames}
	for frame := 0; frame < frames; frame++ {
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, generateImage(frame, *gaps))
	}
	gif.EncodeAll(os.Stdout, &anim)
}

func generateImage(frame, gaps int) *image.Paletted {
	bounds := image.Rect(0, 0, outerSquareSize, outerSquareSize)
	img := image.NewPaletted(bounds, cmyk)
	for c := 1; c <= 4; c++ {
		a := float64(c*frame) * math.Pi / 180 / rate // deg to rad
		for x := -radius; x < radius; x += gaps {
			fx := float64(x)
			for y := -radius; y < radius; y += gaps {
				fy := float64(y)
				rx := int(fx*math.Cos(a) - fy*math.Sin(a))
				ry := int(fx*math.Sin(a) + fy*math.Cos(a))
				img.SetColorIndex(rx+center, ry+center, uint8(c))
			}
		}
	}
	return img
}
