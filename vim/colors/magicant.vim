set background=light
hi clear
if exists("syntax_on")
  syntax reset
endif
let g:colors_name = 'magicant'

hi Cursor	guifg=#000000 guibg=#7100b1
hi lCursor	guifg=#000000 guibg=#7100b1
hi CursorLine	guifg=#000000 guibg=#F9F8B9 gui=underline
hi CursorIM	guifg=#000000 guibg=#7100b1
hi Directory	guifg=#7100b1 guibg=#F9F8B9 gui=bold
hi ErrorMsg	guifg=#ee1111 guibg=#F9F8B9
hi VertSplit	guifg=#F9F8B9 guibg=#543858 " fg is bg
hi Folded	guifg=#cccccc guibg=#F9F8B9
hi FoldColumn	guifg=#000000 guibg=#102010
hi IncSearch	guifg=#000000 guibg=#543858 gui=none
hi LineNr	guifg=#000000 guibg=#F9F8B9
hi ModeMsg	guifg=#cccccc guibg=#F9F8B9
hi MoreMsg	guifg=#cccccc guibg=#F9F8B9
hi Normal	guifg=#220088 guibg=#F9F8B9
hi Question	guifg=#cccccc guibg=#F9F8B9
hi Search	guifg=#000000 guibg=#543858 gui=none
hi NonText	guifg=#606060 gui=none
hi SpecialKey	guifg=#707070
"\n, \0, %d, %s, etc...
hi Special	guifg=#604633 guibg=#F9F8B9 gui=bold
" status line
"hi StatusLine	guifg=#88ee99 guibg=#447f55 gui=bold
"hi StatusLineNC term=bold cterm=bold,underline ctermfg=green ctermbg=Black
"hi StatusLineNC term=bold gui=bold,underline guifg=#220088  guibg=Black
hi Title	guifg=#604633 guibg=#F9F8B9 gui=bold " this is the executable one
hi Visual	guifg=#7100b1 guibg=#543858 gui=none
hi VisualNOS	guifg=#cccccc guibg=#F9F8B9
hi WarningMsg	guifg=#7100b1 guibg=#F9F8B9
hi WildMenu	guifg=#220088 guibg=#7100b1
hi Number	guifg=#604633 guibg=#F9F8B9 gui=underline
hi Char		guifg=#7100b1 guibg=#F9F8B9
hi String	guifg=#a31100 guibg=#F9F8B9 gui=italic
hi Boolean	guifg=#7100b1 guibg=#F9F8B9
hi Comment	guifg=#000000
hi Constant	guifg=#87000d gui=bold,underline
hi Identifier	guifg=#7100b1
hi Statement	guifg=#604633 gui=none
"hi MatchParen guifg=#F9F8B9 guibg=#d992ff

"Procedure name
hi Function     guifg=#7100b1 gui=bold

"Define, def
hi PreProc	guifg=#7100b1 gui=bold
hi Type		guifg=#7100b1 gui=bold
hi Underlined	guifg=#7100b1 gui=underline
hi Error	guifg=#ee1111 guibg=#F9F8B9
hi Todo		guifg=#212121 guibg=#87000d gui=none
hi SignColumn   guibg=#F9F8B9

if version >= 700
   "Pmenu
  hi Pmenu	guibg=#090510
  hi PmenuSel	guibg=#220088 guifg=#7100b1
  hi PmenuSbar	guibg=#000000

   "Tab
  hi TabLine	  guifg=black guibg=black gui=bold
  hi TabLineFill  guifg=black guibg=black gui=bold
  hi TabLineSel	  guifg=#7100b1 guibg=black gui=bold
endif

hi cCustomProp gui=italic
hi cCustomFunc guifg=#9D02F2 guibg=#F9F8B9 gui=bold "hi def cCustomFunc link cCustomFunc  Function
" hi def link cCustomClass Function
