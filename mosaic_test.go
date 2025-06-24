package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"
)

func TestMosaic(t *testing.T) {
	baseImage := image.NewRGBA(image.Rect(0, 0, 10, 10))
	baseImage.Set(1, 1, color.RGBA{R: 255, A: 255})
	baseImage.Set(5, 5, color.RGBA{G: 255, A: 255})
	baseImage.Set(9, 9, color.RGBA{B: 255, A: 255})

	file, err := os.Create("test.png")
	if err != nil {
		t.Errorf("Error creating file: %v", err)
	}
	defer file.Close()

	if err := png.Encode(file, baseImage); err != nil {
		t.Errorf("Error encoding image: %v", err)
	}

	m, err := GenerateMosaic("test.png", 1000, 1000)
	if err != nil {
		t.Error("Error generating mosaic", err)
	}

	serialized := SerializeMosaic(m)
	deserialized := DeserializeMosaic(serialized)

	if deserialized.Image != m.Image {
		t.Error("deserialized image does not match original")
	} else if len(deserialized.Pixels) != len(m.Pixels) {
		t.Error("deserialized pixel count different than original")
	} else if len(deserialized.Colors) != len(m.Colors) {
		t.Error("deserialized color count different than original")
	}

	for idx, col := range m.Colors {
		if deserialized.Colors[idx] != col {
			t.Error("deserialized colors don't match original")
		}
	}

	for idxX, pixels := range m.Pixels {
		for idxY := range pixels {
			if deserialized.Pixels[idxX][idxY] != m.Pixels[idxX][idxY] {
				t.Error("deserialized pixels don't match original")
			}
		}
	}
}
