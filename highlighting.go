package main

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
	"strings"
)

const (
	spaceChar       = " " // lipgloss seems to strip out regular trailing whitespace
	maxPreviewWidth = 58
	sampleText      = `
 === Python ===
 {keyword}def{/keyword} {function}return_pi_or_nan{/function}(return_pi: {type}bool{/type}) -> {type}float{/type}:
     {function}print{/function}({string}"Hello world!"{/string})
     {comment}# This is a comment{/comment}
     pi: {type}float{/type} = {number}3.14159{/number}
     {keyword}if{/keyword} return_pi:
         {keyword}return{/keyword} pi
     {keyword}return{/keyword} {constant}numpy{/constant}.nan

 === Go ===
 {keyword}func{/keyword} {function}returnPiOrNaN{/function}(returnPi {type}bool{/type}) {type}float64{/type} {
         {constant}fmt{/constant}.{function}Println{/function}({string}"Hello world!"{/string})
         {comment}// This is a comment{/comment}
         pi := {number}3.14159{/number}
         {keyword}if{/keyword} returnPi {
                 {keyword}return{/keyword} pi
         }

         {keyword}return{/keyword} {constant}math{/constant}.{function}NaN(){/function}    
 }`
)

func parseSyntaxHighlighting(code string, theme *Theme) string {
	var result strings.Builder
	lines := strings.Split(code, "\n")

	r := lipgloss.Renderer{}
	r.SetColorProfile(termenv.TrueColor)

	baseStyle := lipgloss.NewStyle().
		Renderer(&r).
		Foreground(lipgloss.Color(*theme.Foreground)).
		Background(lipgloss.Color(*theme.Background))

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
					Renderer(&r).
					Foreground(lipgloss.Color(theme.getColor(tag))).
					Background(lipgloss.Color(*theme.Background))
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
