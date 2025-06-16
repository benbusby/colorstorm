package main

import "math"

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

// Adapted from https://github.com/tmux/tmux/pull/432/files
func rgbToX256(r, g, b uint8) uint8 {
	ir := v2ci(r)
	ig := v2ci(g)
	ib := v2ci(b)
	colorIndex := 36*ir + 6*ig + ib

	// Calculate the nearest 0-based gray index at 232 .. 255
	var grayIndex uint8
	average := (r + g + b) / 3
	if average > 238 {
		grayIndex = 23
	} else {
		grayIndex = (average - 3) / 10
	}

	// Calculate the represented colors back from the index
	i2cv := []uint8{0, 0x5F, 0x87, 0xAF, 0xD7, 0xFF}
	cr := i2cv[ir]
	cg := i2cv[ig]
	cb := i2cv[ib]
	gv := 8 + 10*grayIndex

	// Return the one which is nearer to the original input rgb value
	colorErr := distSquare(cr, cg, cb, r, g, b)
	grayErr := distSquare(gv, gv, gv, r, g, b)

	if colorErr <= grayErr {
		return 16 + colorIndex
	} else {
		return 232 + grayIndex
	}
}
