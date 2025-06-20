package main

import (
	"errors"
	"fmt"
	"github.com/lucasb-eyer/go-colorful"
	"strings"
)

const (
	BackgroundIndex int = iota
	ForegroundIndex
	FunctionIndex
	ConstantIndex
	KeywordIndex
	CommentIndex
	NumberIndex
	StringIndex
	TypeIndex

	BackgroundKey = "background"
	ForegroundKey = "foreground"
	FunctionKey   = "function"
	ConstantKey   = "constant"
	KeywordKey    = "keyword"
	CommentKey    = "comment"
	NumberKey     = "number"
	StringKey     = "string"
	TypeKey       = "type"
)

var keyList = []string{
	BackgroundKey,
	ForegroundKey,
	FunctionKey,
	ConstantKey,
	KeywordKey,
	CommentKey,
	NumberKey,
	StringKey,
	TypeKey,
}

type FinalizedTheme struct {
	ID     string
	Name   string
	Author string

	Background     string
	BackgroundAlt1 string
	BackgroundAlt2 string

	Foreground    string
	ForegroundAlt string

	Function string
	Constant string
	Keyword  string
	Comment  string
	Number   string
	String   string
	Type     string

	BackgroundX256     uint8
	BackgroundAlt1X256 uint8
	BackgroundAlt2X256 uint8
	ForegroundX256     uint8
	ForegroundAltX256  uint8
	FunctionX256       uint8
	ConstantX256       uint8
	KeywordX256        uint8
	CommentX256        uint8
	NumberX256         uint8
	StringX256         uint8
	TypeX256           uint8

	DarkOrLight string
}

type Theme struct {
	Name string `json:"name"`

	Background *string `json:"background"`
	Foreground *string `json:"foreground"`
	Function   *string `json:"function"`
	Constant   *string `json:"constant"`
	Keyword    *string `json:"keyword"`
	Comment    *string `json:"comment"`
	Number     *string `json:"number"`
	String     *string `json:"string"`
	Type       *string `json:"type"`

	Reference []byte `json:"reference"`
}

func (t *Theme) getColor(key string) string {
	switch key {
	case BackgroundKey:
		return *t.Background
	case ForegroundKey:
		return *t.Foreground
	case FunctionKey:
		return *t.Function
	case ConstantKey:
		return *t.Constant
	case KeywordKey:
		return *t.Keyword
	case CommentKey:
		return *t.Comment
	case NumberKey:
		return *t.Number
	case StringKey:
		return *t.String
	case TypeKey:
		return *t.Type
	}

	// Invalid key
	return "#ff0000"
}

func newDefaultTheme() *Theme {
	bg := "#1e001e"
	fg := "#d4d4d4"
	functionColor := "#dcdcaa"
	constantColor := "#50fa7b"
	keywordColor := "#569cd6"
	commentColor := "#6a9955"
	numberColor := "#b5cea8"
	stringColor := "#ce9178"
	typeColor := "#50fa7b"
	return &Theme{
		Background: &bg,
		Foreground: &fg,
		Function:   &functionColor,
		Constant:   &constantColor,
		Keyword:    &keywordColor,
		Comment:    &commentColor,
		Number:     &numberColor,
		String:     &stringColor,
		Type:       &typeColor,
	}
}

func newRandomTheme() *Theme {
	palette, err := colorful.WarmPalette(7)
	if err != nil {
		panic(err)
	}

	var bg, fg colorful.Color
	maxV := 0.0

	// Find the brightest color in the palette
	for _, color := range palette {
		_, _, v := color.Hsv()
		if v > maxV {
			maxV = v
			bg = color
		}
	}

	h, s, v := bg.Hsv()
	bg = colorful.Hsv(h, s, v*0.15)
	fg = colorful.Hsv(h, s*0.5, 1.0)

	bgHex := bg.Hex()
	fgHex := fg.Hex()
	fnHex := changeColorBrightness(palette[0], 2.0).Hex()
	constHex := changeColorBrightness(palette[1], 2.0).Hex()
	keywordHex := changeColorBrightness(palette[2], 2.0).Hex()
	commentHex := changeColorBrightness(palette[3], 0.9).Hex()
	numberHex := changeColorBrightness(palette[4], 2.0).Hex()
	stringHex := changeColorBrightness(palette[5], 2.0).Hex()
	typeHex := changeColorBrightness(palette[6], 2.0).Hex()

	return &Theme{
		Background: &bgHex,
		Foreground: &fgHex,
		Function:   &fnHex,
		Constant:   &constHex,
		Keyword:    &keywordHex,
		Comment:    &commentHex,
		Number:     &numberHex,
		String:     &stringHex,
		Type:       &typeHex,
	}
}

func (t *Theme) GetColorName(idx int) string {
	switch idx {
	case BackgroundIndex:
		return "Background"
	case ForegroundIndex:
		return "Foreground"
	case FunctionIndex:
		return "Functions"
	case ConstantIndex:
		return "Constants"
	case KeywordIndex:
		return "Keywords"
	case CommentIndex:
		return "Comments"
	case NumberIndex:
		return "Numbers"
	case StringIndex:
		return "Strings"
	case TypeIndex:
		return "Types"
	}

	return "?"
}

func (t *Theme) GetHexColor(idx int) *string {
	switch idx {
	case BackgroundIndex:
		return t.Background
	case ForegroundIndex:
		return t.Foreground
	case FunctionIndex:
		return t.Function
	case ConstantIndex:
		return t.Constant
	case KeywordIndex:
		return t.Keyword
	case CommentIndex:
		return t.Comment
	case NumberIndex:
		return t.Number
	case StringIndex:
		return t.String
	case TypeIndex:
		return t.Type
	}

	return nil
}

func (t *Theme) SetHexColor(idx int, newVal string) {
	switch idx {
	case BackgroundIndex:
		t.Background = &newVal
	case ForegroundIndex:
		t.Foreground = &newVal
	case FunctionIndex:
		t.Function = &newVal
	case ConstantIndex:
		t.Constant = &newVal
	case KeywordIndex:
		t.Keyword = &newVal
	case CommentIndex:
		t.Comment = &newVal
	case NumberIndex:
		t.Number = &newVal
	case StringIndex:
		t.String = &newVal
	case TypeIndex:
		t.Type = &newVal
	}
}

func (t *Theme) Validate() error {
	var errorList []string
	if len(t.Name) == 0 {
		errorList = append(errorList, "name cannot be empty")
	}

	for _, key := range keyList {
		color := t.getColor(key)
		if len(color) == 0 || color[0] != '#' {
			msg := fmt.Sprintf("missing or invalid '%s' color", key)
			errorList = append(errorList, msg)
		}
	}

	if len(errorList) > 0 {
		errorMsg := strings.Join(errorList, " * ")
		return errors.New(errorMsg)
	}

	return nil
}

func (t *Theme) Finalize(values GeneratorFormValues) FinalizedTheme {
	var (
		fgAlt,
		bgAlt1,
		bgAlt2 colorful.Color
	)

	final := FinalizedTheme{
		ID:         sanitizeName(t.Name),
		Name:       t.Name,
		Author:     values.Author,
		Background: *t.Background,
		Foreground: *t.Foreground,
		Function:   *t.Function,
		Constant:   *t.Constant,
		Keyword:    *t.Keyword,
		Comment:    *t.Comment,
		Number:     *t.Number,
		String:     *t.String,
		Type:       *t.Type,

		BackgroundX256: rgbToX256(*t.Background),
		ForegroundX256: rgbToX256(*t.Foreground),
		FunctionX256:   rgbToX256(*t.Function),
		ConstantX256:   rgbToX256(*t.Constant),
		KeywordX256:    rgbToX256(*t.Keyword),
		CommentX256:    rgbToX256(*t.Comment),
		NumberX256:     rgbToX256(*t.Number),
		StringX256:     rgbToX256(*t.String),
		TypeX256:       rgbToX256(*t.Type),
	}

	bgCol, _ := colorful.Hex(*t.Background)
	bgH, bgS, bgV := bgCol.Hsv()
	fgCol, _ := colorful.Hex(*t.Foreground)
	fgH, fgS, fgV := fgCol.Hsv()

	if values.IsLight {
		fgAlt = colorful.Hsv(fgH, fgS, fgV+0.1)
		bgAlt1 = colorful.Hsv(bgH, bgS, bgV-0.05)
		bgAlt2 = colorful.Hsv(bgH, bgS, bgV-0.1)
		final.DarkOrLight = "light"
	} else {
		fgAlt = colorful.Hsv(fgH, fgS, fgV-0.1)
		bgAlt1 = colorful.Hsv(bgH, bgS, bgV+0.05)
		bgAlt2 = colorful.Hsv(bgH, bgS, bgV+0.1)
		final.DarkOrLight = "dark"
	}

	final.ForegroundAlt = fgAlt.Hex()
	final.BackgroundAlt1 = bgAlt1.Hex()
	final.BackgroundAlt2 = bgAlt2.Hex()

	final.ForegroundAltX256 = rgbToX256(fgAlt.Hex())
	final.BackgroundAlt1X256 = rgbToX256(bgAlt1.Hex())
	final.BackgroundAlt2X256 = rgbToX256(bgAlt2.Hex())

	return final
}
