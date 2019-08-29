set background=dark
hi clear
if exists("syntax_on")
  syntax reset
endif
let g:colors_name = 'threed'

hi Cursor	guifg=#080808 guibg=#d4cbff
hi lCursor	guifg=#080808 guibg=#d4cbff
hi CursorLine	guibg=#080808 gui=underline
hi CursorIM	guifg=#080808 guibg=#d4cbff
hi Directory	guifg=#d4cbff guibg=#080808 gui=bold
hi ErrorMsg	guifg=#ee1111 guibg=#080808
hi VertSplit	guifg=#080808 guibg=#aaaaaa
hi Folded	guifg=#cccccc guibg=#080808
hi FoldColumn	guifg=#557755 guibg=#102010
hi IncSearch	guifg=#3a553a guibg=#d4cbff gui=none
hi LineNr	guifg=#aaaaaa guibg=#080808 " guifg=#446644 guibg=#080808 gui=none
hi ModeMsg	guifg=#cccccc guibg=#080808
hi MoreMsg	guifg=#cccccc guibg=#080808
hi Normal	guifg=#F0FAFF guibg=#080808
hi Question	guifg=#cccccc guibg=#080808
hi Search	guifg=#223322 guibg=#cccccc gui=none
hi NonText	guifg=#606060 gui=none
hi SpecialKey	guifg=#707070
"\n, \0, %d, %s, etc...
hi Special	guifg=#2fff89 guibg=#080808 gui=bold
" status line
hi StatusLine	guifg=#88ee99 guibg=#447f55 gui=bold
hi StatusLineNC term=bold cterm=bold,underline ctermfg=green ctermbg=Black
hi StatusLineNC term=bold gui=bold,underline guifg=#3a553a  guibg=Black
hi Title	guifg=#d4cbff guibg=#080808 gui=bold
hi Visual	guifg=#d4cbff guibg=#448844 gui=none
hi VisualNOS	guifg=#cccccc guibg=#080808
hi WarningMsg	guifg=#d4cbff guibg=#080808
hi WildMenu	guifg=#3a553a guibg=#d4cbff
hi Number	guifg=#2fff89 guibg=#080808 gui=underline
hi Char		guifg=#d4cbff guibg=#080808
hi String	guifg=#ffcf32 guibg=#080808 gui=italic
hi Boolean	guifg=#d4cbff guibg=#080808
hi Comment	guifg=#cccccc "guifg=#c67ff4
hi Constant	guifg=#c67ff4 gui=bold,underline "guifg=#fd35fa gui=none
hi Identifier	guifg=#d4cbff
hi Statement	guifg=#2fff89 gui=none
"hi MatchParen guifg=#080808 guibg=#5e9aff

"Procedure name
hi Function     guifg=#d4cbff gui=bold

"Define, def
hi PreProc	guifg=#d4cbff gui=bold
hi Type		guifg=#d4cbff gui=bold
hi Underlined	guifg=#d4cbff gui=underline
hi Error	guifg=#ee1111 guibg=#080808
hi Todo		guifg=#223322 guibg=#cccccc gui=none
hi SignColumn   guibg=#080808

if version >= 700
  " Pmenu
  hi Pmenu	guibg=#222222
  hi PmenuSel	guibg=#3a553a guifg=#d4cbff
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

hi cCustomFunc guifg=#F89070 guibg=#080808 gui=bold "hi def cCustomFunc link cCustomFunc  Function
" hi def link cCustomClass Function
