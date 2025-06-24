package main

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
	"strings"
	"testing"
)

var (
	keys = []interface{}{
		"function",
		"constant",
		"keyword",
		"comment",
		"number",
		"string",
		"type",
	}

	hlTest = fmt.Sprintf("normal"+
		"{%[1]s}%[1]s{/%[1]s}"+
		"{%[2]s}%[2]s{/%[2]s}"+
		"{%[3]s}%[3]s{/%[3]s}"+
		"{%[4]s}%[4]s{/%[4]s}"+
		"{%[5]s}%[5]s{/%[5]s}"+
		"{%[6]s}%[6]s{/%[6]s}"+
		"{%[7]s}%[7]s{/%[7]s}", keys...)
)

func TestHighlighting(t *testing.T) {
	r := lipgloss.Renderer{}
	r.SetColorProfile(termenv.TrueColor)

	testTheme := newRandomTheme(false)

	var expected string
	output := parseSyntaxHighlighting(hlTest, testTheme)
	for _, c := range "normal" {
		expected += lipgloss.NewStyle().Renderer(&r).
			Foreground(lipgloss.Color(*testTheme.Foreground)).
			Background(lipgloss.Color(*testTheme.Background)).
			Render(string(c))
	}

	totalLength := len("normal")

	for _, key := range keys {
		totalLength += len(key.(string))
		expected += lipgloss.NewStyle().Renderer(&r).
			Foreground(lipgloss.Color(testTheme.getColor(key.(string)))).
			Background(lipgloss.Color(*testTheme.Background)).
			Render(key.(string))
	}

	expected += lipgloss.NewStyle().Renderer(&r).
		Foreground(lipgloss.Color(*testTheme.Foreground)).
		Background(lipgloss.Color(*testTheme.Background)).
		Render(strings.Repeat(spaceChar, maxPreviewWidth-totalLength))

	t.Logf("Expected: %s", expected)
	t.Logf("Output:   %s", output)

	if strings.TrimSpace(output) != strings.TrimSpace(expected) {
		t.Error("incorrect output")
	}
}
