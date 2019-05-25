" requires :set termguicolors
" author: Ben Busby

set background=dark
hi clear
if exists("syntax_on")
  syntax reset
endif
let g:colors_name = 'moonside'

hi Cursor	guifg=#000000 guibg=#F6F929
hi lCursor	guifg=#000000 guibg=#F6F929
hi CursorLine	guibg=#000000 gui=underline
hi CursorIM	guifg=#000000 guibg=#F6F929
hi Directory	guifg=#F6F929 guibg=#000000 gui=bold
hi DiffAdd	guifg=#F6F929 guibg=#3a553a gui=none
hi DiffChange	guifg=#F6F929 guibg=#3a553a gui=none
hi DiffDelete	guifg=#223322 guibg=#223322 gui=none
hi DiffText	guifg=#F6F929 guibg=#448844 gui=bold
hi ErrorMsg	guifg=#ee1111 guibg=#000000
hi VertSplit	guifg=#000000 guibg=#fd35fa
hi Folded	guifg=#cccccc guibg=#000000
hi FoldColumn	guifg=#557755 guibg=#102010
hi IncSearch	guifg=#3a553a guibg=#F6F929 gui=none
hi LineNr	guifg=#fD35fa guibg=#000000 " guifg=#446644 guibg=#000000 gui=none
hi ModeMsg	guifg=#cccccc guibg=#000000
hi MoreMsg	guifg=#cccccc guibg=#000000
hi Normal	guifg=#ffffff guibg=#000000
hi Question	guifg=#cccccc guibg=#000000
hi Search	guifg=#223322 guibg=#cccccc gui=none
hi NonText	guifg=#606060 gui=none
hi SpecialKey	guifg=#707070
"\n, \0, %d, %s, etc...
hi Special	guifg=#aaef64 guibg=#000000 gui=bold
" status line
hi StatusLine	guifg=#88ee99 guibg=#447f55 gui=bold
hi StatusLineNC term=bold cterm=bold,underline ctermfg=green ctermbg=Black
hi StatusLineNC term=bold gui=bold,underline guifg=#3a553a  guibg=Black
hi Title	guifg=#fd35fa guibg=#000000 gui=bold
hi Visual	guifg=#F6F929 guibg=#448844 gui=none
hi VisualNOS	guifg=#cccccc guibg=#000000
hi WarningMsg	guifg=#F6F929 guibg=#000000
hi WildMenu	guifg=#3a553a guibg=#F6F929
hi Number	guifg=#aaef64 guibg=#000000 gui=underline
hi Char		guifg=#F6F929 guibg=#000000
hi String	guifg=#ff6693 guibg=#000000 gui=italic
hi Boolean	guifg=#F6F929 guibg=#000000
hi Comment	guifg=#cccccc "guifg=#c67ff4
hi Constant	guifg=#c67ff4 gui=bold,underline "guifg=#fd35fa gui=none
hi Identifier	guifg=#F6F929
hi Statement	guifg=#aaef64 gui=none
"hi MatchParen guifg=#000000 guibg=#5e9aff

"Procedure name
hi Function     guifg=#F6F929 gui=bold

"Define, def
hi PreProc	guifg=#F6F929 gui=bold
hi Type		guifg=#F6F929 gui=bold
hi Underlined	guifg=#F6F929 gui=underline
hi Error	guifg=#ee1111 guibg=#000000
hi Todo		guifg=#223322 guibg=#cccccc gui=none
hi SignColumn   guibg=#000000

if version >= 700
  " Pmenu
  hi Pmenu	guibg=#222222
  hi PmenuSel	guibg=#3a553a guifg=#F6F929
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

hi cCustomFunc guifg=#5e9aff guibg=#000000 gui=bold "hi def cCustomFunc link cCustomFunc  Function
" hi def link cCustomClass Function
