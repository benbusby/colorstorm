set background=dark
hi clear
if exists("syntax_on")
  syntax reset
endif
let g:colors_name = 'fire-spring'

"Background: #181020
"Foreground: #ffffca
"Color 1: #d9c400
"Color 2: #d992ff
"Color 3: #eeeeee (for less common uses like numbers, possibly)
"Color 4: #f0e500
"Color 5: #ff7e7e (strings? maybe a really light grey would be better)
"Color 6: #a99ade
"Color 7: #ff7f3d (need to check this against the other colors)


hi Cursor	guifg=#FFFFFF guibg=#d9c400
hi lCursor	guifg=#FFFFFF guibg=#d9c400
hi CursorLine	guifg=#FFFFFF guibg=#181020 gui=underline
hi CursorIM	guifg=#FFFFFF guibg=#d9c400
hi Directory	guifg=#d9c400 guibg=#181020 gui=bold
"hi DiffAdd	guifg=#d9c400 guibg=#ffffca gui=none
"hi DiffChange	guifg=#d9c400 guibg=#ffffca gui=none
"hi DiffDelete	guifg=#223322 guibg=#223322 gui=none
"hi DiffText	guifg=#d9c400 guibg=#448844 gui=bold
hi ErrorMsg	guifg=#ee1111 guibg=#181020
hi VertSplit	guifg=#543858 guibg=#543858 " fg is bg
hi Folded	guifg=#cccccc guibg=#181020
hi FoldColumn	guifg=#ffffff guibg=#102010
hi IncSearch	guifg=#ffffff guibg=#543858 gui=none
hi LineNr	guifg=#74e4f3 guibg=#181020
hi ModeMsg	guifg=#cccccc guibg=#181020
hi MoreMsg	guifg=#cccccc guibg=#181020
hi Normal	guifg=#ffffca guibg=#181020
hi Question	guifg=#cccccc guibg=#181020
hi Search	guifg=#ffffff guibg=#543858 gui=none
hi NonText	guifg=#606060 gui=none
hi SpecialKey	guifg=#707070
"\n, \0, %d, %s, etc...
hi Special	guifg=#f0e500 guibg=#181020 gui=bold
" status line
"hi StatusLine	guifg=#88ee99 guibg=#447f55 gui=bold
"hi StatusLineNC term=bold cterm=bold,underline ctermfg=green ctermbg=Black
"hi StatusLineNC term=bold gui=bold,underline guifg=#ffffca  guibg=Black
hi Title	guifg=#f0e500 guibg=#181020 gui=bold " this is the executable one
hi Visual	guifg=#d9c400 guibg=#543858 gui=none
hi VisualNOS	guifg=#cccccc guibg=#181020
hi WarningMsg	guifg=#d9c400 guibg=#181020
hi WildMenu	guifg=#ffffca guibg=#d9c400
hi Number	guifg=#f0e500 guibg=#181020 gui=underline
hi Char		guifg=#d9c400 guibg=#181020
hi String	guifg=#ff7e50 guibg=#181020 gui=italic
hi Boolean	guifg=#d9c400 guibg=#181020
hi Comment	guifg=#74e4f3
hi Constant	guifg=#a99ade gui=bold,underline
hi Identifier	guifg=#d9c400
hi Statement	guifg=#f0e500 gui=none
"hi MatchParen guifg=#181020 guibg=#d992ff

"Procedure name
hi Function     guifg=#d9c400 gui=bold

"Define, def
hi PreProc	guifg=#d9c400 gui=bold
hi Type		guifg=#d9c400 gui=bold
hi Underlined	guifg=#d9c400 gui=underline
hi Error	guifg=#ee1111 guibg=#181020
hi Todo		guifg=#212121 guibg=#a99ade gui=none
hi SignColumn   guibg=#181020

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
hi cCustomFunc guifg=#d992ff guibg=#181020 gui=bold "hi def cCustomFunc link cCustomFunc  Function
" hi def link cCustomClass Function
