" Theme:   theme_name_full
" File:    theme_name_alt.vim

" requires :set termguicolors
set background=dark
hi clear
if exists("syntax_on")
  syntax reset
endif
let g:colors_name = 'theme_name_alt'

hi Cursor        guifg=color_bg_main   guibg=color_fg
hi lCursor       guifg=color_bg_main   guibg=color_fg
hi CursorIM      guifg=color_bg_main   guibg=color_fg
hi Directory     guifg=color_type      guibg=color_bg_main  gui=bold
hi ErrorMsg      guifg=#ee1111         guibg=color_bg_main
hi VertSplit     guifg=color_bg_main   guibg=color_accent
hi LineNr        guifg=color_linenr    guibg=color_bg_alt2
hi ModeMsg       guifg=color_comment   guibg=color_bg_main
hi MoreMsg       guifg=color_comment   guibg=color_bg_main
hi Normal        guifg=color_fg        guibg=color_bg_main
hi Question      guifg=color_comment   guibg=color_bg_main
hi Search        guifg=#223322         guibg=color_comment  gui=none
hi NonText       guifg=#606060                              gui=none
hi SpecialKey    guifg=#707070
"\n, \0, %d, %s, etc...
hi Special       guifg=color_accent                         gui=bold
" status line
hi Title         guifg=color_accent    guibg=color_bg_main  gui=bold
hi Visual                              guibg=color_select   gui=none
hi VisualNOS                           guibg=color_bg_main
hi WarningMsg    guifg=color_type      guibg=color_bg_main
hi Number        guifg=color_number                         gui=underline
hi Char          guifg=color_string                       
hi String        guifg=color_string                         gui=italic         
hi Boolean       guifg=color_boolean                      
hi Comment       guifg=color_comment
hi Constant      guifg=color_variable                       gui=bold
hi Identifier    guifg=color_type
hi Statement     guifg=color_accent                         gui=none
hi CursorLine                          guibg=color_bg_alt2
hi CursorLineNR  guifg=color_accent                         gui=bold

"Procedure name
hi Function      guifg=color_function                       gui=bold

"Define, def
hi PreProc       guifg=color_type                           gui=bold
hi Type          guifg=color_accent                         gui=bold
hi Underlined    guifg=color_type                           gui=underline
hi Error         guifg=#ee1111         guibg=color_bg_main
hi Todo          guifg=color_bg_main   guibg=color_comment  gui=none
hi SignColumn                          guibg=color_bg_main

if version >= 700
  " Pmenu
  hi Pmenu                             guibg=#222222
  hi PmenuSel     guifg=color_type     guibg=#3a553a
  hi PmenuSbar                         guibg=#222222

  " Tab
  hi TabLine      guifg=#3a553a        guibg=black          gui=bold
  hi TabLineFill  guifg=black          guibg=black          gui=bold
  hi TabLineSel   guifg=#88ee99        guibg=#447f55        gui=bold
endif

" Highlight Class and Function names
syn match    cCustomParen    "(" contains=cParen,cCppParen
syn match    cCustomFunc     "\w\+\s*(" contains=cCustomParen
syn match    cCustomScope    "::"
syn match    cCustomClass    "\w\+\s*::" contains=cCustomScope
syn match    cCustomProp     "\.\w\+\s*."

"hi cCustomProp                                                        
hi cCustomFunc    guifg=color_function                      gui=bold 

hi diffAdded ctermfg=green guifg=#00FF00
hi diffRemoved ctermfg=red guifg=#FF0000
