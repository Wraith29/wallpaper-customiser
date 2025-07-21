package main

import (
	"bytes"
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"os"
)

func getNextFrameIndex(count, idx int) int {
	idx++
	if idx >= count {
		return 0
	}

	return idx
}

func main() {
	gifBuf, err := os.ReadFile("../assets/rain.gif")
	if err != nil {
		panic(err)
	}

	anim, err := gif.DecodeAll(bytes.NewReader(gifBuf))
	if err != nil {
		panic(err)
	}

	backdropBuf, err := os.ReadFile("../assets/rain_backdrop.gif")
	if err != nil {
		panic(err)
	}

	backdrop, err := gif.DecodeAll(bytes.NewReader(backdropBuf))
	if err != nil {
		panic(err)
	}

	bdFc := len(backdrop.Image)
	if bdFc <= 0 {
		panic("invalid gif")
	}

	animFc := len(anim.Image)

	bdIdx := 0
	animIdx := 0

	imgBounds := backdrop.Image[0].Bounds()
	frameCount := max(bdFc, animFc)
	frames := make([]*image.Paletted, frameCount)

	for idx := range frameCount {
		snapshot := image.NewPaletted(backdrop.Image[0].Bounds(), palette.Plan9)

		draw.Draw(snapshot, imgBounds, backdrop.Image[bdIdx], image.Pt(0, 0), draw.Over)
		draw.Draw(snapshot, imgBounds, anim.Image[animIdx], image.Pt(0, 0), draw.Over)

		frames[idx] = snapshot

		bdIdx = getNextFrameIndex(bdFc, bdIdx)
		animIdx = getNextFrameIndex(animFc, animIdx)
	}

	delay := make([]int, frameCount)
	for idx := range frameCount {
		delay[idx] = 10
	}

	final := gif.GIF{
		Image:    frames,
		Delay:    delay,
		Disposal: nil,
		Config: image.Config{
			ColorModel: nil,
			Width:      imgBounds.Dx(),
			Height:     imgBounds.Dy(),
		},
		BackgroundIndex: 0,
		LoopCount:       0,
	}

	f, err := os.Create("../assets/layered_gifs.gif")
	if err != nil {
		panic(err)
	}

	if err := gif.EncodeAll(f, &final); err != nil {
		panic(err)
	}
}
