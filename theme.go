package main

import (
	"errors"
	"fmt"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/muesli/gamut"
	"image/color"
	"math/rand/v2"
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

	blackHex = "#000000"
	whiteHex = "#ffffff"
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

func newNoColorTheme(light bool) *Theme {
	var (
		bg, fg string
	)

	if light {
		bg = whiteHex
		fg = blackHex
	} else {
		bg = blackHex
		fg = whiteHex
	}

	functionColor := fg
	constantColor := fg
	keywordColor := fg
	commentColor := fg
	numberColor := fg
	stringColor := fg
	typeColor := fg
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

func newRandomMonoTheme(light bool) *Theme {
	colors, _ := gamut.Generate(16, gamut.HappyGenerator{})
	monoColor := colors[rand.IntN(len(colors))]
	return newMonoTheme(light, monoColor)
}

func newMonoTheme(light bool, initColor color.Color) *Theme {
	var (
		palette         []color.Color
		bg, fg, comment color.Color
	)

	if light {
		bg = gamut.Tints(initColor, 4)[2]
		palette = gamut.Shades(initColor, 8)
		fg = palette[6]
		comment = gamut.Darker(bg, 0.7)
	} else {
		initColor = gamut.Lighter(initColor, 0.5)
		bg = gamut.Shades(initColor, 4)[2]
		palette = gamut.Tints(initColor, 8)
		fg = palette[6]
		comment = gamut.Lighter(bg, 0.7)
	}

	bgHex := gamut.ToHex(bg)
	fgHex := gamut.ToHex(fg)
	fnHex := gamut.ToHex(palette[0])
	constHex := gamut.ToHex(palette[1])
	keywordHex := gamut.ToHex(palette[2])
	commentHex := gamut.ToHex(comment)
	numberHex := gamut.ToHex(palette[3])
	stringHex := gamut.ToHex(palette[4])
	typeHex := gamut.ToHex(palette[5])

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

func newRandomTheme(light bool) *Theme {
	palette, err := gamut.Generate(7, gamut.PastelGenerator{})
	if err != nil {
		panic(err)
	}

	var bg, fg, comment color.Color
	anchor := palette[rand.IntN(len(palette))]

	if light {
		bg = gamut.Lighter(anchor, 0.25)
		fg = gamut.Darker(anchor, 0.45)
		comment = gamut.Darker(anchor, 0.2)
	} else {
		bg = gamut.Darker(anchor, 0.9)
		fg = gamut.Lighter(anchor, 0.2)
		comment = gamut.Darker(anchor, 0.4)
	}

	modFn := gamut.Lighter
	modVal := 0.2
	if light {
		modFn = gamut.Darker
		modVal = 0.6
	}

	bgHex := gamut.ToHex(bg)
	fgHex := gamut.ToHex(fg)
	fnHex := gamut.ToHex(modFn(palette[0], modVal))
	constHex := gamut.ToHex(modFn(palette[1], modVal))
	keywordHex := gamut.ToHex(modFn(palette[2], modVal))
	commentHex := gamut.ToHex(comment)
	numberHex := gamut.ToHex(modFn(palette[4], modVal))
	stringHex := gamut.ToHex(modFn(palette[5], modVal))
	typeHex := gamut.ToHex(modFn(palette[6], modVal))

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
		tColor := t.getColor(key)
		if len(tColor) == 0 || tColor[0] != '#' {
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

		BackgroundX256: hexToX256(*t.Background),
		ForegroundX256: hexToX256(*t.Foreground),
		FunctionX256:   hexToX256(*t.Function),
		ConstantX256:   hexToX256(*t.Constant),
		KeywordX256:    hexToX256(*t.Keyword),
		CommentX256:    hexToX256(*t.Comment),
		NumberX256:     hexToX256(*t.Number),
		StringX256:     hexToX256(*t.String),
		TypeX256:       hexToX256(*t.Type),
	}

	bgCol, _ := colorful.Hex(*t.Background)
	bgH, bgS, bgV := bgCol.Hsv()
	fgCol, _ := colorful.Hex(*t.Foreground)
	fgH, fgS, fgV := fgCol.Hsv()

	if values.IsLight {
		fgAlt = colorful.Hsv(fgH, fgS, fgV+0.1)
		bgAlt1 = colorful.Hsv(bgH, bgS, bgV-0.03)
		bgAlt2 = colorful.Hsv(bgH, bgS, bgV-0.06)
		final.DarkOrLight = "light"
	} else {
		fgAlt = colorful.Hsv(fgH, fgS, fgV-0.1)
		bgAlt1 = colorful.Hsv(bgH, bgS, bgV+0.03)
		bgAlt2 = colorful.Hsv(bgH, bgS, bgV+0.06)
		final.DarkOrLight = "dark"
	}

	final.ForegroundAlt = fgAlt.Hex()
	final.BackgroundAlt1 = bgAlt1.Hex()
	final.BackgroundAlt2 = bgAlt2.Hex()

	final.ForegroundAltX256 = hexToX256(fgAlt.Hex())
	final.BackgroundAlt1X256 = hexToX256(bgAlt1.Hex())
	final.BackgroundAlt2X256 = hexToX256(bgAlt2.Hex())

	return final
}
