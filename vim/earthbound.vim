set background=dark
hi clear
if exists("syntax_on")
  syntax reset
endif
let g:colors_name = 'earthbound'

hi Cursor	guifg=#1c0037 guibg=#fafd51
hi lCursor	guifg=#1c0037 guibg=#fafd51
hi CursorLine	guibg=#1c0037 gui=underline
hi CursorIM	guifg=#1c0037 guibg=#fafd51
hi Directory	guifg=#fafd51 guibg=#1c0037 gui=bold
hi ErrorMsg	guifg=#ee1111 guibg=#1c0037
hi VertSplit	guifg=#1c0037 guibg=#aaaaaa
hi Folded	guifg=#cccccc guibg=#1c0037
hi FoldColumn	guifg=#557755 guibg=#102010
hi IncSearch	guifg=#3a553a guibg=#fafd51 gui=none
hi LineNr	guifg=#aaaaaa guibg=#1c0037 " guifg=#446644 guibg=#1c0037 gui=none
hi ModeMsg	guifg=#cccccc guibg=#1c0037
hi MoreMsg	guifg=#cccccc guibg=#1c0037
hi Normal	guifg=#ffffff guibg=#1c0037
hi Question	guifg=#cccccc guibg=#1c0037
hi Search	guifg=#223322 guibg=#cccccc gui=none
hi NonText	guifg=#606060 gui=none
hi SpecialKey	guifg=#707070
"\n, \0, %d, %s, etc...
hi Special	guifg=#ffd69b guibg=#1c0037 gui=bold
" status line
hi StatusLine	guifg=#88ee99 guibg=#447f55 gui=bold
hi StatusLineNC term=bold cterm=bold,underline ctermfg=green ctermbg=Black
hi StatusLineNC term=bold gui=bold,underline guifg=#3a553a  guibg=Black
hi Title	guifg=#fafd51 guibg=#1c0037 gui=bold
hi Visual	guifg=#fafd51 guibg=#448844 gui=none
hi VisualNOS	guifg=#cccccc guibg=#1c0037
hi WarningMsg	guifg=#fafd51 guibg=#1c0037
hi WildMenu	guifg=#3a553a guibg=#fafd51
hi Number	guifg=#25ff81 guibg=#1c0037 gui=underline
hi Char		guifg=#fafd51 guibg=#1c0037
hi String	guifg=#fb967f guibg=#1c0037 gui=italic
hi Boolean	guifg=#fafd51 guibg=#1c0037
hi Comment	guifg=#aaaaaa "guifg=#c67ff4
hi Constant	guifg=#70caff gui=bold,underline "guifg=#fd35fa gui=none
hi Identifier	guifg=#fafd51
hi Statement	guifg=#abdcdc gui=none
"hi MatchParen guifg=#1c0037 guibg=#5e9aff

"Procedure name
hi Function     guifg=#fafd51 gui=bold

"Define, def
hi PreProc	guifg=#fafd51 gui=bold
hi Type		guifg=#fafd51 gui=bold
hi Underlined	guifg=#fafd51 gui=underline
hi Error	guifg=#ee1111 guibg=#1c0037
hi Todo		guifg=#223322 guibg=#cccccc gui=none
hi SignColumn   guibg=#1c0037

if version >= 700
  " Pmenu
  hi Pmenu	guibg=#222222
  hi PmenuSel	guibg=#3a553a guifg=#fafd51
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

hi cCustomFunc guifg=#70caff guibg=#1c0037 gui=bold "hi def cCustomFunc link cCustomFunc  Function
" hi def link cCustomClass Function
