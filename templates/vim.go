package templates

import (
	"fmt"
	"text/template"
)

func GetVimThemeFileName(name string) string {
	return fmt.Sprintf("%s.vim", name)
}

func GetVimThemeTemplate() (*template.Template, error) {
	return template.New("vim_template").Parse(VimTemplate)
}

const VimTemplate = `" Theme:   {{ .Name }}
" File:    {{ .ID }}.vim
" Author:  {{ .Author }}
" requires :set termguicolors
set background={{ .DarkOrLight }}
hi clear
if exists("syntax_on")
  syntax reset
endif
let g:colors_name = '{{ .ID }}'

hi Cursor        guifg={{ .Background }}     guibg={{ .Foreground }}
hi lCursor       guifg={{ .Background }}     guibg={{ .Foreground }}
hi CursorIM      guifg={{ .Background }}     guibg={{ .Foreground }}
hi Directory     guifg={{ .Type }}           guibg={{ .Background }}     gui=bold
hi ErrorMsg      guifg=#ee1111               guibg={{ .Background }}
hi VertSplit     guifg={{ .Background }}     guibg={{ .Keyword }}
hi LineNr        guifg={{ .ForegroundAlt }}  guibg={{ .BackgroundAlt1 }}
hi ModeMsg       guifg={{ .Comment }}        guibg={{ .Background }}
hi MoreMsg       guifg={{ .Comment }}        guibg={{ .Background }}
hi Normal        guifg={{ .Foreground }}     guibg={{ .Background }}
hi Question      guifg={{ .Comment }}        guibg={{ .Background }}
hi Search        guifg=#223322               guibg={{ .Comment }}        gui=none
hi NonText       guifg=#606060                                           gui=none
hi SpecialKey    guifg=#707070
"\n, \0, %d, %s, etc...
hi Special       guifg={{ .Keyword }}                                    gui=bold
" status line
hi Title         guifg={{ .Keyword }}        guibg={{ .Background }}     gui=bold
hi Visual                                    guibg={{ .BackgroundAlt2 }} gui=none
hi VisualNOS                                 guibg={{ .Background }}
hi WarningMsg    guifg={{ .Type }}           guibg={{ .Background }}
hi Number        guifg={{ .Number }}                                     gui=underline
hi Char          guifg={{ .String }}
hi String        guifg={{ .String }}                                     gui=italic
hi Boolean       guifg={{ .Constant }}
hi Comment       guifg={{ .Comment }}
hi Constant      guifg={{ .Constant }}                                   gui=bold
hi Identifier    guifg={{ .Type }}
hi Statement     guifg={{ .Keyword }}                                    gui=none
hi CursorLine                                guibg={{ .BackgroundAlt2 }}
hi CursorLineNR  guifg={{ .Keyword }}                                    gui=bold

"Procedure name
hi Function      guifg={{ .Function }}                                   gui=bold

"Define, def
hi PreProc       guifg={{ .Type }}                                       gui=bold
hi Type          guifg={{ .Keyword }}                                    gui=bold
hi Underlined    guifg={{ .Type }}                                       gui=underline
hi Error         guifg=#ee1111               guibg={{ .Background }}
hi Todo          guifg={{ .Background }}     guibg={{ .Comment }}        gui=none
hi SignColumn                                guibg={{ .Background }}

if version >= 700
  " Pmenu
  hi Pmenu        guifg={{ .Foreground }} guibg={{ .BackgroundAlt2 }}
  hi PmenuSel     guifg={{ .Keyword }}    guibg={{ .Background }}
  hi PmenuSbar                            guibg={{ .BackgroundAlt1 }}

  " Tab
  hi TabLine      guifg={{ .Foreground }} guibg={{ .BackgroundAlt1 }}   gui=bold
  hi TabLineFill  guifg={{ .Background }} guibg={{ .Background }}       gui=bold
  hi TabLineSel   guifg={{ .Foreground }} guibg={{ .BackgroundAlt2 }}   gui=bold
endif

" Highlight Class and Function names
syn match    cCustomParen    "(" contains=cParen,cCppParen
syn match    cCustomFunc     "\w\+\s*(" contains=cCustomParen
syn match    cCustomScope    "::"
syn match    cCustomClass    "\w\+\s*::" contains=cCustomScope
syn match    cCustomProp     "\.\w\+\s*."

"hi cCustomProp
hi cCustomFunc    guifg={{ .Function }}                      gui=bold

hi diffAdded ctermfg=green guifg=#00FF00
hi diffRemoved ctermfg=red guifg=#FF0000`
