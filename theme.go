package main

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
)

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
	case "background":
		return *t.Background
	case "foreground":
		return *t.Foreground
	case "function":
		return *t.Function
	case "constant":
		return *t.Constant
	case "keyword":
		return *t.Keyword
	case "comment":
		return *t.Comment
	case "number":
		return *t.Number
	case "string":
		return *t.String
	case "type":
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
