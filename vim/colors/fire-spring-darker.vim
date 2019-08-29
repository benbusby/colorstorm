set background=dark
hi clear
if exists("syntax_on")
  syntax reset
endif
let g:colors_name = 'fire-spring'

hi Cursor	guifg=#FFFFFF guibg=#d9c400
hi lCursor	guifg=#FFFFFF guibg=#d9c400
hi CursorLine	guifg=#FFFFFF guibg=#080808 gui=underline
hi CursorIM	guifg=#FFFFFF guibg=#d9c400
hi Directory	guifg=#d9c400 guibg=#080808 gui=bold
hi ErrorMsg	guifg=#ee1111 guibg=#080808
hi VertSplit	guifg=#080808 guibg=#543858 " fg is bg
hi Folded	guifg=#cccccc guibg=#080808
hi FoldColumn	guifg=#ffffff guibg=#102010
hi IncSearch	guifg=#ffffff guibg=#543858 gui=none
hi LineNr	guifg=#74e4f3 guibg=#080808
hi ModeMsg	guifg=#cccccc guibg=#080808
hi MoreMsg	guifg=#cccccc guibg=#080808
hi Normal	guifg=#ffffca guibg=#080808
hi Question	guifg=#cccccc guibg=#080808
hi Search	guifg=#ffffff guibg=#543858 gui=none
hi NonText	guifg=#606060 gui=none
hi SpecialKey	guifg=#707070
"\n, \0, %d, %s, etc...
hi Special	guifg=#f0e500 guibg=#080808 gui=bold
" status line
"hi StatusLine	guifg=#88ee99 guibg=#447f55 gui=bold
"hi StatusLineNC term=bold cterm=bold,underline ctermfg=green ctermbg=Black
"hi StatusLineNC term=bold gui=bold,underline guifg=#ffffca  guibg=Black
hi Title	guifg=#f0e500 guibg=#080808 gui=bold " this is the executable one
hi Visual	guifg=#d9c400 guibg=#543858 gui=none
hi VisualNOS	guifg=#cccccc guibg=#080808
hi WarningMsg	guifg=#d9c400 guibg=#080808
hi WildMenu	guifg=#ffffca guibg=#d9c400
hi Number	guifg=#f0e500 guibg=#080808 gui=underline
hi Char		guifg=#d9c400 guibg=#080808
hi String	guifg=#ff7e50 guibg=#080808 gui=italic
hi Boolean	guifg=#d9c400 guibg=#080808
hi Comment	guifg=#74e4f3
hi Constant	guifg=#a99ade gui=bold,underline
hi Identifier	guifg=#d9c400
hi Statement	guifg=#f0e500 gui=none
"hi MatchParen guifg=#080808 guibg=#d992ff

"Procedure name
hi Function     guifg=#d9c400 gui=bold

"Define, def
hi PreProc	guifg=#d9c400 gui=bold
hi Type		guifg=#d9c400 gui=bold
hi Underlined	guifg=#d9c400 gui=underline
hi Error	guifg=#ee1111 guibg=#080808
hi Todo		guifg=#212121 guibg=#a99ade gui=none
hi SignColumn   guibg=#080808

if version >= 700
   "Pmenu
  hi Pmenu	guibg=#090510
  hi PmenuSel	guibg=#ffffca guifg=#d9c400
  hi PmenuSbar	guibg=#000000

   "Tab
  hi TabLine	  guifg=black guibg=black gui=bold
  hi TabLineFill  guifg=black guibg=black gui=bold
  hi TabLineSel	  guifg=#d9c400 guibg=black gui=bold
endif

hi cCustomProp gui=italic
hi cCustomFunc guifg=#d992ff guibg=#080808 gui=bold "hi def cCustomFunc link cCustomFunc  Function
" hi def link cCustomClass Function
