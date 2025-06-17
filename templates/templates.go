package templates

import _ "embed"

const (
	VimKey     = "vim"
	VSCodeKey  = "vscode"
	SublimeKey = "sublime"
)

//go:embed sublime_theme.tmTheme
var SublimeThemeTemplate string

//go:embed vscode_theme.json
var VSCodeThemeTemplate string

//go:embed vim_theme.vim
var VimThemeTemplate string
