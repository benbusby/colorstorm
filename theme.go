package main

type Theme struct {
	Name string `json:"name"`

	Background string `json:"background"`
	Foreground string `json:"foreground"`
	Function   string `json:"function"`
	Constant   string `json:"constant"`
	Keyword    string `json:"keyword"`
	Comment    string `json:"comment"`
	Number     string `json:"number"`
	String     string `json:"string"`
	Type       string `json:"type"`
}

func (t Theme) getColor(key string) string {
	switch key {
	case "background":
		return t.Background
	case "foreground":
		return t.Foreground
	case "function":
		return t.Function
	case "constant":
		return t.Constant
	case "keyword":
		return t.Keyword
	case "comment":
		return t.Comment
	case "number":
		return t.Number
	case "string":
		return t.String
	case "type":
		return t.Type
	}

	// Invalid key
	return "#ff0000"
}

func newTheme() Theme {
	return Theme{
		Background: "#1e1e1e",
		Foreground: "#d4d4d4",
		Function:   "#dcdcaa",
		Constant:   "#50fa7b",
		Keyword:    "#569cd6",
		Comment:    "#6a9955",
		Number:     "#b5cea8",
		String:     "#ce9178",
		Type:       "#50fa7b",
	}
}
