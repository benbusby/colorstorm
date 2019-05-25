set background=dark
hi clear
if exists("syntax_on")
  syntax reset
endif
let g:colors_name = 'threed'

hi Cursor	guifg=#303454 guibg=#ded7ff
hi lCursor	guifg=#303454 guibg=#ded7ff
hi CursorLine	guibg=#303454 gui=underline
hi CursorIM	guifg=#303454 guibg=#ded7ff
hi Directory	guifg=#ded7ff guibg=#303454 gui=bold
hi ErrorMsg	guifg=#ee1111 guibg=#303454
hi VertSplit	guifg=#303454 guibg=#aaaaaa
hi Folded	guifg=#cccccc guibg=#303454
hi FoldColumn	guifg=#557755 guibg=#102010
hi IncSearch	guifg=#3a553a guibg=#ded7ff gui=none
hi LineNr	guifg=#aaaaaa guibg=#303454 " guifg=#446644 guibg=#303454 gui=none
hi ModeMsg	guifg=#cccccc guibg=#303454
hi MoreMsg	guifg=#cccccc guibg=#303454
hi Normal	guifg=#F0FAFF guibg=#303454
hi Question	guifg=#cccccc guibg=#303454
hi Search	guifg=#223322 guibg=#cccccc gui=none
hi NonText	guifg=#606060 gui=none
hi SpecialKey	guifg=#707070
"\n, \0, %d, %s, etc...
hi Special	guifg=#2fff89 guibg=#303454 gui=bold
" status line
hi StatusLine	guifg=#88ee99 guibg=#447f55 gui=bold
hi StatusLineNC term=bold cterm=bold,underline ctermfg=green ctermbg=Black
hi StatusLineNC term=bold gui=bold,underline guifg=#3a553a  guibg=Black
hi Title	guifg=#ded7ff guibg=#303454 gui=bold
hi Visual	guifg=#ded7ff guibg=#448844 gui=none
hi VisualNOS	guifg=#cccccc guibg=#303454
hi WarningMsg	guifg=#ded7ff guibg=#303454
hi WildMenu	guifg=#3a553a guibg=#ded7ff
hi Number	guifg=#2fff89 guibg=#303454 gui=underline
hi Char		guifg=#ded7ff guibg=#303454
hi String	guifg=#ffcf32 guibg=#303454 gui=italic
hi Boolean	guifg=#ded7ff guibg=#303454
hi Comment	guifg=#cccccc "guifg=#c67ff4
hi Constant	guifg=#c67ff4 gui=bold,underline "guifg=#fd35fa gui=none
hi Identifier	guifg=#ded7ff
hi Statement	guifg=#2fff89 gui=none
"hi MatchParen guifg=#303454 guibg=#5e9aff

"Procedure name
hi Function     guifg=#ded7ff gui=bold

"Define, def
hi PreProc	guifg=#ded7ff gui=bold
hi Type		guifg=#ded7ff gui=bold
hi Underlined	guifg=#ded7ff gui=underline
hi Error	guifg=#ee1111 guibg=#303454
hi Todo		guifg=#223322 guibg=#cccccc gui=none
hi SignColumn   guibg=#303454

if version >= 700
  " Pmenu
  hi Pmenu	guibg=#222222
  hi PmenuSel	guibg=#3a553a guifg=#ded7ff
  hi PmenuSbar	guibg=#222222

  " Tab
  hi TabLine	  guifg=#3a553a guibg=black gui=bold
  hi TabLineFill  guifg=black guibg=black gui=bold
  hi TabLineSel	  guifg=#88ee99 guibg=#447f55 gui=bold
endif

" Highlight Class and Function names
syn match    cCustomParen    "(" contains=cParen,cCppParen
syn match    cCustomFunc     "\w\+\s*(" contains=cCustomParen
syn match    cCustomScope    "::"
syn match    cCustomClass    "\w\+\s*::" contains=cCustomScope
syn match    cCustomProp     "\.\w\+\s*."

hi cCustomProp gui=italic

hi cCustomFunc guifg=#F89070 guibg=#303454 gui=bold "hi def cCustomFunc link cCustomFunc  Function
" hi def link cCustomClass Function
