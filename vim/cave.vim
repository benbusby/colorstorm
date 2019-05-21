set background=dark
hi clear
if exists("syntax_on")
  syntax reset
endif
let g:colors_name = 'cave'

hi Cursor	guifg=#9DB9A3 guibg=#2b342a
hi lCursor	guifg=#9DB9A3 guibg=#2b342a
hi CursorLine	guibg=#9DB9A3 gui=underline
hi CursorIM	guifg=#9DB9A3 guibg=#2b342a
hi Directory	guifg=#2b342a guibg=#9DB9A3 gui=bold
hi DiffAdd	guifg=#2b342a guibg=#9DB9A3 gui=none
hi DiffChange	guifg=#2b342a guibg=#9DB9A3 gui=none
hi DiffDelete	guifg=#223322 guibg=#223322 gui=none
hi DiffText	guifg=#2b342a guibg=#668876 gui=bold
hi ErrorMsg	guifg=#ee1111 guibg=#9DB9A3
hi VertSplit	guifg=#9DB9A3 guibg=#223322
hi Folded	guifg=#0e1e0e guibg=#9DB9A3
hi FoldColumn	guifg=#557755 guibg=#102010
hi IncSearch	guifg=#9DB9A3 guibg=#2b342a gui=none
hi LineNr	guifg=#0e1e0e guibg=#9DB9A3 " guifg=#446644 guibg=#9DB9A3 gui=none
hi ModeMsg	guifg=#0e1e0e guibg=#9DB9A3
hi MoreMsg	guifg=#0e1e0e guibg=#9DB9A3
hi Normal	guifg=#394537 guibg=#9DB9A3
hi Question	guifg=#0e1e0e guibg=#9DB9A3
hi Search	guifg=#9DB9A3 guibg=#0e1e0e gui=none
hi NonText	guifg=#606060 gui=none
hi SpecialKey	guifg=#707070
"\n, \0, %d, %s, etc...
hi Special	guifg=#394537 guibg=#9DB9A3 gui=bold
" status line
hi StatusLine	guifg=#293d29 guibg=#000000 gui=bold
hi StatusLineNC term=bold cterm=bold,underline ctermfg=green ctermbg=Black
hi StatusLineNC term=bold gui=bold,underline guifg=#9DB9A3  guibg=Black
hi Title	guifg=#2b342a guibg=#9DB9A3 gui=bold
hi Visual	guifg=#2b342a guibg=#778877 gui=none
hi VisualNOS	guifg=#0e1e0e guibg=#9DB9A3
hi WarningMsg	guifg=#2b342a guibg=#9DB9A3
hi WildMenu	guifg=#9DB9A3 guibg=#2b342a
hi Number	guifg=#2b342a guibg=#9DB9A3 gui=underline
hi Char		guifg=#2b342a guibg=#9DB9A3
hi String	guifg=#2b342a guibg=#9DB9A3 gui=italic
hi Boolean	guifg=#2b342a guibg=#9DB9A3
hi Comment	guifg=#555555
hi Constant	guifg=#293d29 gui=none
hi Identifier	guifg=#2b342a
hi Statement	guifg=#293d29 gui=none

"Procedure name
hi Function     guifg=#2b342a gui=bold

"Define, def
hi PreProc	guifg=#2b342a gui=bold
hi Type		guifg=#2b342a gui=bold
hi Underlined	guifg=#2b342a gui=underline
hi Error	guifg=#ee1111 guibg=#9DB9A3
hi Todo		guifg=#223322 guibg=#0e1e0e gui=none
hi SignColumn   guibg=#9DB9A3

if version >= 700
  " Pmenu
  hi Pmenu	guibg=#222222
  hi PmenuSel	guibg=#9DB9A3 guifg=#9DB9A3
  hi PmenuSbar	guibg=#222222

  " Tab
  hi TabLine	  guifg=#9DB9A3 guibg=black gui=bold
  hi TabLineFill  guifg=black guibg=black gui=bold
  hi TabLineSel	  guifg=#293d29 guibg=#447f55 gui=bold
endif

hi cCustomFunc guifg=#445046 guibg=#9DB9A3 gui=bold "hi def cCustomFunc link cCustomFunc  Function
