package main

import (
	"errors"
	"fmt"
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

	Foreground string
	Function   string
	Constant   string
	Keyword    string
	Comment    string
	Number     string
	String     string
	Type       string

	LightTheme bool
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

func newTheme() *Theme {
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
	} else if len(t.Name) > 20 {
		errorList = append(errorList, "name must be < 20 characters")
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

func (t *Theme) Finalize() FinalizedTheme {
	return FinalizedTheme{}
}
