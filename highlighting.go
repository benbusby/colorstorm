package main

import (
	"github.com/charmbracelet/lipgloss"
	"strings"
)

const (
	spaceChar       = " " // lipgloss seems to strip out regular trailing whitespace
	maxPreviewWidth = 58
)

func parseSyntaxHighlighting(code string, theme Theme) string {
	var result strings.Builder
	lines := strings.Split(code, "\n")
	baseStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Foreground)).
		Background(lipgloss.Color(theme.Background))

	for _, line := range lines {
		i := 0
		previewLine := strings.Builder{}
		rawLine := strings.Builder{}
		for i < len(line) {
			if i+1 < len(line) && line[i:i+1] == "{" {
				tagEnd := strings.Index(line[i:], "}")
				if tagEnd == -1 {
					previewLine.WriteString(line[i:])
					rawLine.WriteString(line[i:])
					break
				}
				tagEnd += i

				tag := line[i+1 : tagEnd]

				closingTag := "{/" + tag + "}"
				closingTagPos := strings.Index(line[tagEnd+1:], closingTag)
				if closingTagPos == -1 {
					previewLine.WriteString(line[i:])
					rawLine.WriteString(line[i:])
					break
				}
				closingTagPos += tagEnd + 1

				content := line[tagEnd+1 : closingTagPos]

				style := lipgloss.NewStyle().
					Foreground(lipgloss.Color(theme.getColor(tag))).
					Background(lipgloss.Color(theme.Background))
				previewLine.WriteString(style.Render(content))
				rawLine.WriteString(content)

				i = closingTagPos + len(closingTag)
			} else {
				previewLine.WriteString(baseStyle.Render(line[i : i+1]))
				rawLine.WriteString(line[i : i+1])
				i++
			}
		}
		if len(rawLine.String()) < maxPreviewWidth {
			padding := strings.Repeat(spaceChar, maxPreviewWidth-len(rawLine.String()))
			previewLine.WriteString(baseStyle.Render(padding))
		}
		result.WriteString(previewLine.String())
		result.WriteString("\n")
	}

	return result.String()
}
