package main

import (
	"bytes"
	"os"

	"image"
	"image/draw"
	"image/png"
)

func main() {
	btmBytes, _ := os.ReadFile("assets/bottom.png")
	topBytes, _ := os.ReadFile("assets/top.png")

	btm, _, _ := image.Decode(bytes.NewReader(btmBytes))
	top, _, _ := image.Decode(bytes.NewReader(topBytes))

	_ = btm
	_ = top

	base := image.NewRGBA(btm.Bounds())
	draw.Draw(base, btm.Bounds(), btm, image.Point{0, 0}, draw.Over)
	draw.Draw(base, top.Bounds(), top, image.Point{0, 0}, draw.Over)

	f, _ := os.Create("output.png")

	png.Encode(f, base)

	f.Close()
}
