package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"image/png"
	"os"
)

const (
	width  = 64
	height = 64
)

func main() {
	buf, err := os.ReadFile("../assets/backdrop.png")
	if err != nil {
		panic(err)
	}

	bd, _, err := image.Decode(bytes.NewReader(buf))
	if err != nil {
		panic(err)
	}

	frameCount := bd.Bounds().Dx() - width
	frames := make([]*image.Paletted, frameCount)

	for idx := range frameCount {
		bounds := image.Rect(0, 0, width, height)

		snapshot := image.NewPaletted(bounds, palette.Plan9)
		draw.Draw(snapshot, bounds, bd, image.Pt(idx, 0), draw.Src)

		f, err := os.Create(fmt.Sprintf("../assets/anim/f_%d.png", idx))
		if err != nil {
			panic(err)
		}

		if err := png.Encode(f, snapshot); err != nil {
			panic(err)
		}

		frames[idx] = snapshot
	}

	delay := make([]int, frameCount)
	for idx := range frameCount {
		delay[idx] = 10
	}

	anim := gif.GIF{
		Image:    frames,
		Delay:    delay,
		Disposal: nil,
		Config: image.Config{
			ColorModel: nil,
			Width:      width,
			Height:     height,
		},
		BackgroundIndex: 0,
		LoopCount:       0,
	}

	f, err := os.Create("../assets/bd_gif.gif")
	if err != nil {
		panic(err)
	}

	if err := gif.EncodeAll(f, &anim); err != nil {
		panic(err)
	}
}
