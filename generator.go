package main

import (
	"bytes"
	_ "embed"
	"github.com/benbusby/colorstorm/templates"
	"text/template"
)

type editorKey string

const (
	VimKey     editorKey = "vim"
	VSCodeKey  editorKey = "vscode"
	SublimeKey editorKey = "sublime"
)

type ThemeFile struct {
	FileName string
	Contents []byte
}

var generatorMap = map[editorKey]func(FinalizedTheme) (ThemeFile, error){
	VimKey:     generateVimTheme,
	VSCodeKey:  generateVSCodeTheme,
	SublimeKey: generateSublimeTheme,
}

func generateVimTheme(theme FinalizedTheme) (ThemeFile, error) {
	themeFileName := templates.GetVimThemeFileName(theme.ID)
	themeTemplate, err := templates.GetVimThemeTemplate()
	if err != nil {
		return ThemeFile{}, err
	}

	themeContents, err := generateTheme(themeTemplate, theme)
	if err != nil {
		return ThemeFile{}, err
	}
	return ThemeFile{FileName: themeFileName, Contents: themeContents}, nil
}

func generateVSCodeTheme(theme FinalizedTheme) (ThemeFile, error) {
	themeFileName := templates.GetVSCodeThemeFileName(theme.ID)
	themeTemplate, err := templates.GetVSCodeThemeTemplate()
	if err != nil {
		return ThemeFile{}, err
	}

	themeContents, err := generateTheme(themeTemplate, theme)
	if err != nil {
		return ThemeFile{}, err
	}

	return ThemeFile{FileName: themeFileName, Contents: themeContents}, nil
}

func generateSublimeTheme(theme FinalizedTheme) (ThemeFile, error) {
	themeFileName := templates.GetSublimeThemeFileName(theme.ID)
	themeTemplate, err := templates.GetSublimeThemeTemplate()
	if err != nil {
		return ThemeFile{}, err
	}

	themeContents, err := generateTheme(themeTemplate, theme)
	if err != nil {
		return ThemeFile{}, err
	}

	return ThemeFile{FileName: themeFileName, Contents: themeContents}, nil
}

func generateTheme(tmpl *template.Template, theme FinalizedTheme) ([]byte, error) {
	var buf bytes.Buffer
	err := tmpl.Execute(&buf, theme)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
