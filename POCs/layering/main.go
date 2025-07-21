package main

import (
	"bytes"
	"image"
	"image/draw"
	"image/png"
	"os"
)

func main() {
	btmBuf, err := os.ReadFile("../assets/bottom.png")
	if err != nil {
		panic(err)
	}

	topBuf, err := os.ReadFile("../assets/top.png")
	if err != nil {
		panic(err)
	}

	btm, _, err := image.Decode(bytes.NewReader(btmBuf))
	if err != nil {
		panic(err)
	}

	top, _, err := image.Decode(bytes.NewReader(topBuf))
	if err != nil {
		panic(err)
	}

	base := image.NewRGBA(btm.Bounds())

	draw.Draw(base, base.Bounds(), btm, image.Point{0, 0}, draw.Over)
	draw.Draw(base, base.Bounds(), top, image.Point{0, 0}, draw.Over)

	f, err := os.Create("../assets/stacked.png")
	if err != nil {
		panic(err)
	}

	if err := png.Encode(f, base); err != nil {
		panic(err)
	}
}
