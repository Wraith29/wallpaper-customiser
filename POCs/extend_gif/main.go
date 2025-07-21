package main

import (
	"bytes"
	"image/gif"
	_ "image/gif"
	"os"
)

func main() {
	buf, err := os.ReadFile("../assets/epilepsy.gif")
	if err != nil {
		panic(err)
	}

	anim, err := gif.DecodeAll(bytes.NewReader(buf))
	if err != nil {
		panic(err)
	}

	println(anim.LoopCount)

	println("Hello, World!")
}
