package main

import (
	"bytes"
	"compress/gzip"
	"github.com/lucasb-eyer/go-colorful"
	"io"
	"log"
	"math"
	"regexp"
	"strings"
)

func v2ci(v uint8) uint8 {
	if v < 48 {
		return 0
	} else if v < 115 {
		return 1
	} else {
		return (v - 35) / 40
	}
}

func distSquare(A, B, C, a, b, c uint8) uint8 {
	return (A-a)*(A-a) + (B-b)*(B-b) + (C-c)*(C-c)
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func hexToRGB(hex string) (uint8, uint8, uint8) {
	c, _ := colorful.Hex(hex)
	r, g, b, _ := c.RGBA()
	return uint8(r), uint8(g), uint8(b)
}

// Adapted from https://github.com/tmux/tmux/pull/432/files
func rgbToX256(hex string) uint8 {
	r, g, b := hexToRGB(hex)
	ir := v2ci(r)
	ig := v2ci(g)
	ib := v2ci(b)
	colorIndex := 36*ir + 6*ig + ib

	var grayIndex uint8
	average := (r + g + b) / 3
	if average > 238 {
		grayIndex = 23
	} else {
		grayIndex = (average - 3) / 10
	}

	i2cv := []uint8{0, 0x5F, 0x87, 0xAF, 0xD7, 0xFF}
	cr := i2cv[ir]
	cg := i2cv[ig]
	cb := i2cv[ib]
	gv := 8 + 10*grayIndex

	colorErr := distSquare(cr, cg, cb, r, g, b)
	grayErr := distSquare(gv, gv, gv, r, g, b)

	if colorErr <= grayErr {
		return 16 + colorIndex
	} else {
		return 232 + grayIndex
	}
}

// sanitizeName removes whitespace and special characters to create a string
// value that can be used as a file name and a theme name for editors like Vim
func sanitizeName(name string) string {
	newName := strings.ReplaceAll(name, " - ", "_")
	newName = strings.ReplaceAll(newName, " ", "_")

	re := regexp.MustCompile(`[\\/:*?!#$%^&().,'"<>|]`)
	newName = re.ReplaceAllString(newName, "")

	return strings.ToLower(newName)
}

func changeColorBrightness(c colorful.Color, amount float64) colorful.Color {
	h, s, v := c.Hsv()
	return colorful.Hsv(h, s, min(v*amount, 1.0))
}

func isMovementKey(key string) bool {
	return strings.Contains(key, "up") || strings.Contains(key, "down") ||
		strings.Contains(key, "left") || strings.Contains(key, "right")
}

func isExitKey(key string) bool {
	return key == "esc" || key == "ctrl+c" || key == "q"
}

func compress(s []byte) []byte {
	buf := bytes.Buffer{}
	zipped := gzip.NewWriter(&buf)
	zipped.Write(s)
	zipped.Close()
	return buf.Bytes()
}

func decompress(s []byte) []byte {
	gz, _ := gzip.NewReader(bytes.NewReader(s))
	data, err := io.ReadAll(gz)
	if err != nil {
		log.Fatal(err)
	}
	gz.Close()
	return data
}
