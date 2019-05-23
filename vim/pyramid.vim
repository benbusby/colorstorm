" requires :set termguicolors
" author: Ben Busby

set termguicolors

set background=dark
hi clear
if exists("syntax_on")
  syntax reset
endif
let g:colors_name = 'pyramid'

hi Cursor	guifg=#3a553a guibg=#f6d56a
hi lCursor	guifg=#3a553a guibg=#f6d56a
hi CursorLine	guibg=#0e0900 gui=underline
hi CursorIM	guifg=#3a553a guibg=#f6d56a
hi Directory	guifg=#f6d56a guibg=#0e0900 gui=bold
hi DiffAdd	guifg=#f6d56a guibg=#3a553a gui=none
hi DiffChange	guifg=#f6d56a guibg=#3a553a gui=none
hi DiffDelete	guifg=#223322 guibg=#223322 gui=none
hi DiffText	guifg=#f6d56a guibg=#448844 gui=bold
hi ErrorMsg	guifg=#ee1111 guibg=#0e0900
hi VertSplit	guifg=#223322 guibg=#223322
hi Folded	guifg=#FFE17F guibg=#0e0900
hi FoldColumn	guifg=#557755 guibg=#102010
hi IncSearch	guifg=#3a553a guibg=#f6d56a gui=none
hi LineNr	guifg=#446644 guibg=#0e0900 " guifg=#446644 guibg=#0e0900 gui=none
hi ModeMsg	guifg=#FFE17F guibg=#0e0900
hi MoreMsg	guifg=#FFE17F guibg=#0e0900
hi Normal	guifg=#f9e4a1 guibg=#0e0900 " guifg=#d8bf6d guifg=#ead286
hi Question	guifg=#FFE17F guibg=#0e0900
hi Search	guifg=#755900 guibg=#ffd036 gui=none
hi NonText	guifg=#606060 gui=none
hi SpecialKey	guifg=#707070
"\n, \0, %d, %s, etc...
hi Special	guifg=#ffefbe " guibg=#223333 gui=bold
" status line
hi StatusLine	guifg=#ffd754 guibg=#447f55 gui=bold
hi StatusLineNC term=bold cterm=bold,underline ctermfg=green ctermbg=Black
hi StatusLineNC term=bold gui=bold,underline guifg=#3a553a  guibg=Black
hi Title	guifg=#f6d56a guibg=#223322 gui=bold
hi Visual	guifg=#0e0900 guibg=#ffe17f gui=none
hi VisualNOS	guifg=#FFE17F guibg=#0e0900
hi WarningMsg	guifg=#f6d56a guibg=#0e0900
hi WildMenu	guifg=#3a553a guibg=#f6d56a
hi Number	guifg=#f6d56a gui=underline " guibg=#5a5930
hi Char		guifg=#f6d56a guibg=#0e0900
hi String	guifg=#f6d56a guibg=#0e0900 gui=italic " guibg=#5a5930
hi Boolean	guifg=#f6d56a guibg=#0e0900
hi Comment	guifg=#446644
hi Constant	guifg=#ffebae gui=none
hi Identifier	guifg=#FFC300 "guifg=#f6d56a
hi Statement    guifg=#ffd03c  gui=bold "guifg=#ffebae

"Procedure name
hi Function     guifg=#f6d56a gui=bold

"Define, def
hi PreProc	guifg=#f6d56a gui=bold
hi Type		guifg=#f6d56a gui=bold
hi Underlined	guifg=#f6d56a gui=underline
hi Error	guifg=#ee1111 guibg=#0e0900
hi Todo		guifg=#223322 guibg=#FFE17F gui=none
hi SignColumn   guibg=#0e0900

if version >= 700
  " Pmenu
  hi Pmenu	guibg=#222222
  hi PmenuSel	guibg=#3a553a guifg=#f6d56a
  hi PmenuSbar	guibg=#222222

  " Tab
  hi TabLine	  guifg=#3a553a guibg=black gui=bold
  hi TabLineFill  guifg=black guibg=black gui=bold
  hi TabLineSel	  guifg=#ffd754 guibg=#447f55 gui=bold
endif

hi cCustomProp gui=italic

hi cCustomFunc guifg=#f6d56a guibg=#0e0900 gui=bold
