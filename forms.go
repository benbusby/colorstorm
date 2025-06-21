package main

import (
	"errors"
	"fmt"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/lucasb-eyer/go-colorful"
)

const (
	saveDraftAction     = -1
	generateThemeAction = -2

	colorDetailHexKey   = "color_hex"
	colorDetailFieldKey = "color_detail"

	editorSelectKey = "editor_select"
)

type GeneratorFormValues struct {
	Author    string
	Editors   []editorKey
	IsLight   bool
	Confirmed bool
}

func generateColorLabel(lg *lipgloss.Renderer, hex, label string) string {
	style := lg.NewStyle()
	color, err := colorful.Hex(hex)
	if err != nil {
		return "error!"
	}

	foreground := lipgloss.Color("#fff")
	background := lipgloss.Color(hex)
	_, _, lightness := color.Hsv()
	if lightness > 0.5 {
		foreground = "#000"
	}

	return style.Foreground(foreground).Background(background).Render(label)
}

func createForm(lg *lipgloss.Renderer) *huh.Form {
	themeName := huh.NewInput().Title("Theme Name").Placeholder("My Theme").Value(&theme.Name)

	bgLabel := generateColorLabel(lg, *theme.Background,
		fmt.Sprintf("Background [%s]", *theme.Background))
	fgLabel := generateColorLabel(lg, *theme.Foreground,
		fmt.Sprintf("Foreground [%s]", *theme.Foreground))
	funcLabel := generateColorLabel(lg, *theme.Function,
		fmt.Sprintf("Function   [%s]", *theme.Function))
	constLabel := generateColorLabel(lg, *theme.Constant,
		fmt.Sprintf("Constant   [%s]", *theme.Constant))
	keywordLabel := generateColorLabel(lg, *theme.Keyword,
		fmt.Sprintf("Keyword    [%s]", *theme.Keyword))
	commentLabel := generateColorLabel(lg, *theme.Comment,
		fmt.Sprintf("Comment    [%s]", *theme.Comment))
	numberLabel := generateColorLabel(lg, *theme.Number,
		fmt.Sprintf("Number     [%s]", *theme.Number))
	stringLabel := generateColorLabel(lg, *theme.String,
		fmt.Sprintf("String     [%s]", *theme.String))
	typeLabel := generateColorLabel(lg, *theme.Type,
		fmt.Sprintf("Type       [%s]", *theme.Type))

	themeSelect := huh.NewSelect[int]().
		Options(
			huh.NewOption(bgLabel, BackgroundIndex),
			huh.NewOption(fgLabel, ForegroundIndex),
			huh.NewOption(funcLabel, FunctionIndex),
			huh.NewOption(constLabel, ConstantIndex),
			huh.NewOption(keywordLabel, KeywordIndex),
			huh.NewOption(commentLabel, CommentIndex),
			huh.NewOption(numberLabel, NumberIndex),
			huh.NewOption(stringLabel, StringIndex),
			huh.NewOption(typeLabel, TypeIndex),
			huh.NewOption("Save Draft", saveDraftAction),
			huh.NewOption("Generate Theme", generateThemeAction),
		).Value(&mainAction)

	form := huh.NewForm(huh.NewGroup(themeName, themeSelect)).
		WithWidth(25).
		WithShowHelp(false).
		WithShowErrors(false).
		WithTheme(huh.ThemeCatppuccin())

	if mainAction > 0 {
		themeSelect.Focus()
	} else {
		themeName.Focus()
	}

	return form
}

func createSaveForm() (*huh.Form, *bool) {
	saveFile := true
	return huh.NewForm(huh.NewGroup(
		huh.NewInput().Title("Save As...").Placeholder("my_theme.json").Value(&fileName),
		huh.NewSelect[bool]().
			Options(
				huh.NewOption("Cancel", false),
				huh.NewOption("Save File", true)).
			Value(&saveFile))).
		WithWidth(25).
		WithShowHelp(false).
		WithShowErrors(false).
		WithTheme(huh.ThemeCatppuccin()), &saveFile
}

func createColorForm(colorName string, hexColor *string) (*huh.Form, *bool) {
	saveColor := true
	color, err := colorful.Hex(*hexColor)
	if err == nil {
		h, s, v := color.Hsv()
		if !colorEdit.init || updateSV {
			colorEdit.S = uint8(roundFloat(s, 2) * 100)
			colorEdit.V = uint8(roundFloat(v, 2) * 100)
			updateSV = false
		}

		colorEdit.Hex = hexColor
		colorEdit.H = uint16(roundFloat(h, 2))
		colorEdit.R = uint8(color.R * 255)
		colorEdit.G = uint8(color.G * 255)
		colorEdit.B = uint8(color.B * 255)
		colorEdit.init = true
	}

	rLabel := fmt.Sprintf("R: [%03d / 255]", colorEdit.R)
	gLabel := fmt.Sprintf("G: [%03d / 255]", colorEdit.G)
	bLabel := fmt.Sprintf("B: [%03d / 255]", colorEdit.B)
	hLabel := fmt.Sprintf("H: [%03d / 360]", colorEdit.H)
	sLabel := fmt.Sprintf("S: [%03d / 100]", colorEdit.S)
	vLabel := fmt.Sprintf("V: [%03d / 100]", colorEdit.V)

	input := huh.NewInput().Title(colorName).Value(hexColor).Key(colorDetailHexKey)
	colorDetails := huh.NewSelect[int]().
		Title("RGB/HSV").
		Description("→ = +1\n← = -1\nshift + → = +10\nshift + ← = -10").
		Options(
			huh.NewOption(rLabel, 0),
			huh.NewOption(gLabel, 1),
			huh.NewOption(bLabel, 2),
			huh.NewOption(hLabel, 3),
			huh.NewOption(sLabel, 4),
			huh.NewOption(vLabel, 5)).
		Value(&colorFormField).
		Key(colorDetailFieldKey)

	form := huh.NewForm(
		huh.NewGroup(
			input,
			colorDetails,
			huh.NewSelect[bool]().
				Options(
					huh.NewOption("Cancel", false),
					huh.NewOption("Save Changes", true)).
				Value(&saveColor),
		)).
		WithWidth(25).
		WithShowHelp(false).
		WithShowErrors(false).
		WithTheme(huh.ThemeCatppuccin())

	if colorFormSelected == 0 {
		for form.GetFocusedField().GetKey() != colorDetailHexKey {
			form.NextField()
		}
	} else if colorFormSelected == 1 {
		for form.GetFocusedField().GetKey() != colorDetailFieldKey {
			form.NextField()
		}
	}

	return form, &saveColor
}

func createGeneratorForm() (*huh.Form, *GeneratorFormValues) {
	final := GeneratorFormValues{Confirmed: true}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Author Name").
				Value(&final.Author),
			huh.NewSelect[bool]().
				Options(
					huh.NewOption("Dark Theme", false),
					huh.NewOption("Light Theme", true)).
				Value(&final.IsLight),
			huh.NewMultiSelect[editorKey]().
				Title("Editors").
				Options(
					huh.NewOption("Vim / NeoVim", VimKey),
					huh.NewOption("VSCode", VSCodeKey),
					huh.NewOption("Sublime", SublimeKey),
				).
				Key(editorSelectKey).
				Validate(func(keys []editorKey) error {
					if len(keys) == 0 {
						return errors.New("must select at least one editor")
					}

					return nil
				}).
				Value(&final.Editors),
			huh.NewSelect[bool]().
				Options(
					huh.NewOption("Cancel", false),
					huh.NewOption("Generate", true),
				).
				Value(&final.Confirmed),
		)).
		WithWidth(25).
		WithShowHelp(false).
		WithShowErrors(false).
		WithTheme(huh.ThemeCatppuccin())

	return form, &final
}
