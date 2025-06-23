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

hi Cursor        guifg={{ .Background }} guibg={{ .Foreground }}  ctermfg={{ .BackgroundX256 }} ctermbg={{ .ForegroundX256 }}
hi lCursor       guifg={{ .Background }} guibg={{ .Foreground }}  ctermfg={{ .BackgroundX256 }} ctermbg={{ .ForegroundX256 }}
hi CursorIM      guifg={{ .Background }} guibg={{ .Foreground }}  ctermfg={{ .BackgroundX256 }} ctermbg={{ .ForegroundX256 }}
hi Directory     guifg={{ .Type }} guibg={{ .Background }}  ctermfg={{ .TypeX256 }} ctermbg={{ .BackgroundX256 }} gui=bold
hi ErrorMsg      guifg=#ee1111 guibg={{ .Background }}  ctermbg={{ .BackgroundX256 }}
hi VertSplit     guifg={{ .Background }} guibg={{ .Keyword }}  ctermfg={{ .BackgroundX256 }} ctermbg={{ .KeywordX256 }}
hi LineNr        guifg={{ .ForegroundAlt }} guibg={{ .BackgroundAlt1 }}  ctermfg={{ .ForegroundAltX256 }} ctermbg={{ .BackgroundAlt1X256 }}
hi ModeMsg       guifg={{ .Comment }} guibg={{ .Background }}  ctermfg={{ .CommentX256 }} ctermbg={{ .BackgroundX256 }}
hi MoreMsg       guifg={{ .Comment }} guibg={{ .Background }}  ctermfg={{ .CommentX256 }} ctermbg={{ .BackgroundX256 }} 
hi Normal        guifg={{ .Foreground }} guibg={{ .Background }}  ctermfg={{ .ForegroundX256 }} ctermbg={{ .BackgroundX256 }} 
hi Question      guifg={{ .Comment }} guibg={{ .Background }}  ctermfg={{ .CommentX256 }} ctermbg={{ .BackgroundX256 }} 
hi Search        guifg={{ .ForegroundAlt }} guibg={{ .BackgroundAlt2 }} gui=none  ctermfg={{ .ForegroundX256 }} ctermbg={{ .BackgroundAlt2X256 }}
hi NonText       gui=none
hi SpecialKey    guifg={{ .Comment }} ctermfg={{ .CommentX256 }}
"\n, \0, %d, %s, etc...
hi Special       guifg={{ .Keyword }} ctermfg={{ .KeywordX256 }} gui=bold
" status line
hi Title         guifg={{ .Keyword }} guibg={{ .Background }} gui=bold  ctermfg={{ .KeywordX256 }} ctermbg={{ .BackgroundX256 }}
hi Visual        guibg={{ .BackgroundAlt2 }} gui=none  ctermbg={{ .BackgroundAlt2X256 }} 
hi VisualNOS     guibg={{ .Background }} ctermbg={{ .BackgroundX256 }}
hi WarningMsg    guifg={{ .Type }} guibg={{ .Background }}  ctermfg={{ .TypeX256 }} ctermbg={{ .BackgroundX256 }}
hi Number        guifg={{ .Number }} ctermfg={{ .NumberX256 }} 
hi Char          guifg={{ .String }} gui=italic ctermfg={{ .StringX256 }}
hi String        guifg={{ .String }} ctermfg={{ .StringX256 }}
hi Boolean       guifg={{ .Constant }}  ctermfg={{ .ConstantX256 }}
hi Comment       guifg={{ .Comment }}  ctermfg={{ .CommentX256 }}
hi Constant      guifg={{ .Constant }} ctermfg={{ .ConstantX256 }} 
hi Identifier    guifg={{ .Type }}  ctermfg={{ .TypeX256 }}
hi Statement     guifg={{ .Keyword }} gui=none  ctermfg={{ .KeywordX256 }} 
hi CursorLine    guibg={{ .BackgroundAlt2 }}  ctermbg={{ .BackgroundAlt2X256 }}
hi CursorLineNR  guifg={{ .Keyword }} gui=bold  ctermfg={{ .KeywordX256 }} 

"Procedure name
hi Function      guifg={{ .Function }}  ctermfg={{ .FunctionX256 }}

"Define, def
hi PreProc       guifg={{ .Type }} gui=bold  ctermfg={{ .TypeX256 }} 
hi Type          guifg={{ .Keyword }} gui=bold  ctermfg={{ .KeywordX256 }}
hi Underlined    guifg={{ .Type }} gui=underline  ctermfg={{ .TypeX256 }} 
hi Error         guifg=#ee1111 guibg={{ .Background }}  ctermbg={{ .BackgroundX256 }}
hi Todo          guifg={{ .Background }} guibg={{ .Comment }} gui=none  ctermfg={{ .BackgroundX256 }} ctermbg={{ .CommentX256 }}
hi SignColumn    guibg={{ .Background }}  ctermbg={{ .BackgroundX256 }}

if version >= 700
  " Pmenu
  hi Pmenu        guifg={{ .Foreground }} guibg={{ .BackgroundAlt2 }}  ctermfg={{ .ForegroundX256 }} ctermbg={{ .BackgroundAlt2X256 }}
  hi PmenuSel     guifg={{ .Keyword }} guibg={{ .Background }}  ctermfg={{ .KeywordX256 }} ctermbg={{ .BackgroundX256 }}
  hi PmenuSbar    guibg={{ .BackgroundAlt1 }}  ctermbg={{ .BackgroundAlt1X256 }}

  " Tab
  hi TabLine      guifg={{ .Foreground }} guibg={{ .BackgroundAlt1 }} gui=bold  ctermfg={{ .ForegroundX256 }} ctermbg={{ .BackgroundAlt1X256 }} 
  hi TabLineFill  guifg={{ .Background }} guibg={{ .Background }} gui=bold  ctermfg={{ .BackgroundX256 }} ctermbg={{ .BackgroundX256 }} 
  hi TabLineSel   guifg={{ .Foreground }} guibg={{ .BackgroundAlt2 }} gui=bold  ctermfg={{ .ForegroundX256 }} ctermbg={{ .BackgroundAlt2X256 }} 
endif

hi diffAdded ctermfg=green guifg=#00FF00
hi diffRemoved ctermfg=red guifg=#FF0000

if has('nvim')
  if has('nvim-0.9')
    hi! link  @lsp.type.class Type
    hi! link  @lsp.type.decorator Function
    hi! link  @lsp.type.enum Type
    hi! link  @lsp.type.enumMember Constant
    hi! link  @lsp.type.function Function
    hi! link  @lsp.type.interface Type
    hi! link  @lsp.type.macro Type
    hi! link  @lsp.type.method Function
    hi! link  @lsp.type.namespace Type
    hi! link  @lsp.type.parameter Keyword
    hi! link  @lsp.type.property Constant
    hi! link  @lsp.type.struct Type
    hi! link  @lsp.type.type Type
    hi! link  @lsp.type.typeParameter Constant
    hi! link  @lsp.type.variable Constant
  endif

  if has('nvim-0.8.1')
    hi! link @type Type
    hi! link @function Function
    hi! link @number Number
    hi! link @string String
    hi! link @comment Comment
    hi! link @keyword Keyword
    hi! link @variable  Constant
    hi! link @constant  Constant
    hi! link @module    Constant
    hi! link @namespace Constant
  endif
else
  " Highlight Class and Function names
  syn match    cCustomParen    "(" contains=cParen,cCppParen
  syn match    cCustomFunc     "\w\+\s*(" contains=cCustomParen
  syn match    cCustomScope    "::"
  syn match    cCustomClass    "\w\+\s*::" contains=cCustomScope
  syn match    cCustomProp     "\.\w\+\s*."

  "hi cCustomProp
  hi cCustomFunc guifg={{ .Function }} gui=bold  ctermfg={{ .FunctionX256 }}
endif
`
