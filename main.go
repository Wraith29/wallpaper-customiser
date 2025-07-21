package main

import (
	"fmt"
	"image"
	"image/gif"
	"os"
	"os/exec"
	"strings"
)

func loadDotenv(path string) error {
	buf, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	lines := strings.SplitSeq(strings.Trim(string(buf), "\n "), "\n")

	for line := range lines {
		keyIdx := strings.Index(line, "=")

		if err := os.Setenv(line[0:keyIdx], line[keyIdx+1:]); err != nil {
			return err
		}
	}

	return nil
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func turnGifToVideo(assetPath, outPath string) error {
	return exec.Command("ffmpeg", "-i", assetPath, outPath).Run()
}

func gifGetNextIndex(frameCount, frame int) int {
	if frame+1 >= frameCount {
		return 0
	}

	return frame + 1
}

func main() {
	handleError(loadDotenv(".env"))

	client, err := NewWeatherClient()
	if err != nil {
		handleError(err)
	}

	weather, err := client.GetWeather()
	if err != nil {
		handleError(err)
	}

	fmt.Printf("%+v\n", weather)
}
