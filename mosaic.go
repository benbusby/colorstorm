package main

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/image/draw"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"os"
	"strings"
)

const colorPreview = "   "

type ReferenceMosaic struct {
	Image      string
	ColorTable string
}

type Pixel struct {
	R uint8
	G uint8
	B uint8
	A uint8

	null bool
}

func (p Pixel) toHex() string {
	return fmt.Sprintf("#%02X%02X%02X", p.R, p.G, p.B)
}

func generateColorTable(colors []string) string {
	output := ""
	maxColorsPerRow := 6
	numColors := 0

	for _, color := range colors {
		style := lipgloss.NewStyle().
			Foreground(lipgloss.Color(color)).
			Background(lipgloss.Color(color))

		output += fmt.Sprintf("%s: %s  ", color, style.Render(colorPreview))
		numColors += 1
		if numColors >= maxColorsPerRow {
			numColors = 0
			output += "\n"
		}
	}

	return output
}

// getPixelPair uses the top-half block character to create a (~)1:1 image render,
// where the "foreground color" of the character fills the top "square" and the
// "background color" fills the bottom "square".
//
// In terms of coordinates, getPixelPair automatically adjusts the specified
// coordinates to account for the fact that a single cell can hold 2 rows:
// 0, 0 --> 0, 0 (top half)
// 0, 1 --> 0, 0 (bottom half)
// 0, 2 --> 0, 1 (top half)
// 0, 3 --> 0, 1 (bottom half)
// etc.
func getPixelPair(top, bottom Pixel) (string, string, string) {
	if top.A == 0 && top.B == 0 {
		return " ", "", ""
	}

	topHex := top.toHex()
	topColor := lipgloss.Color(topHex)
	style := lipgloss.NewStyle()

	if top.A > 0 {
		style = style.Foreground(topColor)
	}

	bottomHex := bottom.toHex()
	if !bottom.null && bottom.A > 0 {
		bottomColor := lipgloss.Color(bottomHex)
		style = style.Background(bottomColor)
	}

	result := style.Render("▀")
	return result, topHex, bottomHex
}

// scaleImage scales the source image to a given pre-sized rectangle
func scaleImage(
	src image.Image,
	rect image.Rectangle,
	scale draw.Scaler,
) image.Image {
	dst := image.NewRGBA(rect)
	scale.Scale(dst, rect, src, src.Bounds(), draw.Over, nil)
	return dst
}

// openImage returns an image given a path to an image file
func openImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

// getTargetDimensions determines the sizing of the output image, restricted by
// the given max width and height.
func getTargetDimensions(src image.Image, maxWidth, maxHeight int) (int, int) {
	srcWidth := src.Bounds().Max.X
	srcHeight := src.Bounds().Max.Y

	xScale := float64(srcWidth) / float64(maxWidth)
	yScale := float64(srcHeight) / float64(maxHeight)

	if xScale > yScale {
		return maxWidth, int(float64(srcHeight) / xScale)
	} else {
		return int(float64(srcWidth) / yScale), maxHeight
	}
}

func getPixels(raw []uint8, width int) [][]Pixel {
	var (
		result [][]Pixel
		row    []Pixel
	)

	idx := 0
	for idx < len(raw) {
		if len(row) == width {
			result = append(result, row)
			row = []Pixel{}
		}

		px := Pixel{
			R: roundPixelRGBA(raw[idx]),
			G: roundPixelRGBA(raw[idx+1]),
			B: roundPixelRGBA(raw[idx+2]),
			A: roundPixelRGBA(raw[idx+3]),
		}

		row = append(row, px)
		idx += 4
	}

	return result
}

func roundPixelRGBA(val uint8) uint8 {
	return uint8(math.Round(float64(val)/float64(50)) * float64(50))
}

// generateReferenceMosaic returns an image that has been rendered into a series
// of colored ASCII half-blocks, as well as the list of colors used in the image.
func generateReferenceMosaic(imgPath string, width, height int) (ReferenceMosaic, error) {
	var result []string

	img, err := openImage(imgPath)
	if err != nil {
		return ReferenceMosaic{}, err
	}

	// Height can be doubled here since the output uses two "pixels" per row
	w, h := getTargetDimensions(img, width, height*2)
	dr := image.Rect(0, 0, w, h)
	res := scaleImage(img, dr, draw.NearestNeighbor)
	rgba := image.NewRGBA(dr)
	draw.Draw(rgba, dr, res, dr.Min, draw.Src)

	pixels := getPixels(rgba.Pix, w)
	row := 0
	colors := make(map[string]struct{})

	for row < len(pixels) {
		line := ""
		for i, p := range pixels[row] {
			top := p
			bottom := Pixel{null: true}

			if len(pixels) > row+1 {
				bottom = pixels[row+1][i]
			}

			imageLine, topHex, bottomHex := getPixelPair(top, bottom)
			line += imageLine
			colors[topHex] = struct{}{}
			colors[bottomHex] = struct{}{}
		}
		result = append(result, line)
		row += 2
	}

	keyColors := make([]string, len(colors))

	i := 0
	for k := range colors {
		keyColors[i] = k
		i++
	}

	output := strings.Join(result, "\n")
	return ReferenceMosaic{
		Image:      output,
		ColorTable: generateColorTable(keyColors),
	}, nil
}
