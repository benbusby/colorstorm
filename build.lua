#!/usr/bin/lua
-- Lua script for automatically building theme files
-- for vim, vscode, sublime, and atom.

USAGE = [[
./test.lua [vim, vscode, sublime]
]]

-- MISC UTILS
local random = math.random
local function uuid()
  local template ='xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'
  return string.gsub(template, '[xy]', function (c)
    local v = (c == 'x') and random(0, 0xf) or random(8, 0xb)
    return string.format('%x', v)
  end)
end

-- COLOR DEFINITIONS

darker_variants = {
  earthbound=1, threed=1, fire_spring=1, dusty_dunes=1
}

color_table = {
  earthbound={
    theme_id=uuid(),
    theme_name_full='Earthbound',
    theme_name_alt='earthbound',
    color_bg_main='#1c0037',
    color_bg_alt1='#2b044f',
    color_bg_alt2='#360a5f',
    color_fg='#ffffff',
    color_x1='#fafd51',
    color_x2='#fb967f',
    color_x3='#70caff',
    color_x4='#abdcdc',
    color_x5='#aaaaaa',
    color_x6='#70caff'
  },
  moonside={
    theme_id=uuid(),
    theme_name_full='Moonside',
    theme_name_alt='moonside',
    color_bg_main='#000000',
    color_bg_alt1='#080808',
    color_bg_alt2='#131313',
    color_fg='#ffffff',
    color_x1='#fd35fa',
    color_x2='#ff6693',
    color_x3='#f6f929',
    color_x4='#c67ff4',
    color_x5='#aaef64',
    color_x6='#5e9aff'
  },
  threed={
    theme_id=uuid(),
    theme_name_full='Threed',
    theme_name_alt='threed',
    color_bg_main='#303454',
    color_bg_alt1='#373c60',
    color_bg_alt2='#3a3f66',
    color_fg='#f0faff',
    color_x1='#d4cbff',
    color_x2='#c67ff4',
    color_x3='#ffcf32',
    color_x4='#2fff89',
    color_x5='#aaaaaa',
    color_x6='#f89070'
  },
  fire_spring={
    theme_id=uuid(),
    theme_name_full='Fire Spring',
    theme_name_alt='fire-spring',
    color_bg_main='#181020',
    color_bg_alt1='#21162c',
    color_bg_alt2='#261933',
    color_fg='#ffffca',
    color_x1='#74e4f3',
    color_x2='#d9c400',
    color_x3='#ff7e50',
    color_x4='#f0e500',
    color_x5='#a99ade',
    color_x6='#d992ff'
  },
  dusty_dunes={
    theme_id=uuid(),
    theme_name_full='Dusty Dunes',
    theme_name_alt='dusty-dunes',
    color_bg_main='#0e0900',
    color_bg_alt1='#150d00',
    color_bg_alt2='#1e1302',
    color_fg='#f9e4a1',
    color_x1='#f6d56a',
    color_x2='#ffebae',
    color_x3='#ffd03c',
    color_x4='#f6d56a',
    color_x5='#666644',
    color_x6='#f6d56a'
  },
  magicant={
    theme_id=uuid(),
    theme_name_full='Magicant (Light)',
    theme_name_alt='magicant-light',
    color_bg_main='#f9f8b9',
    color_bg_alt1='#efeeb2',
    color_bg_alt2='#e6e5ab',
    color_fg='#220088',
    color_x1='#604633',
    color_x2='#7100b1',
    color_x3='#a31100',
    color_x4='#87000d',
    color_x5='#604633',
    color_x6='#9d02f2'
  },
  cave_of_the_past={
    theme_id=uuid(),
    theme_name_full='Cave of the Past (Monochrome)',
    theme_name_alt='cave-of-the-past',
    color_bg_main='#b0d0b8',
    color_bg_alt1='#a5c4ad',
    color_bg_alt2='#9ebba6',
    color_fg='#394537',
    color_x1='#0e1e0e',
    color_x2='#2b342a',
    color_x3='#293d29',
    color_x4='#2b342a',
    color_x5='#445046',
    color_x6='#0e1e0e'
  }
}

color_files = {
  vim='vim/colors/template.vim',
  vscode='vscode/themes/template.json',
  sublime='sublime/earthbound_template.tmTheme',
  atom='atom/colors.less'
}

-- THEME GENERATION

local atom_path = 'atom/%s-syntax/colors.less'

function generate_theme(file, theme)
  local theme_file = io.open(file, 'r')

  if theme_file == nil then
    print('Unable to open ' .. file)
    return
  end

  local lines = {}
  for line in theme_file:lines() do
    for k,v in pairs(color_table[theme]) do
      line = string.gsub(line, k, v)
    end

    lines[#lines + 1] = line
  end
  theme_file:close()

  out_path = string.match(file, 'colors.less') and
    string.format(atom_path, color_table[theme]['theme_name_alt']) or
    string.gsub(file, 'template', theme)

  -- Ensure directory exists
  os.execute('mkdir -p ' .. string.gsub(out_path, 'colors.less', ''))

  theme_file = io.open(out_path, 'w')
  for i,line in ipairs(lines) do
    theme_file:write(line, "\n")
  end
  theme_file:close()
end

if arg[1] == nil then
  print('Error: Argument required')
  print(USAGE)
  os.exit(1)
elseif color_files[arg[1]] == nil then
  print('Error: Invalid argument')
  print(USAGE)
  os.exit(1)
else
  local filename = color_files[arg[1]]
  for theme,_ in pairs(color_table) do
    generate_theme(filename, theme)
    if darker_variants[theme] ~= nil then
      local darker_theme = theme .. '_darker'

      color_table[darker_theme] = color_table[theme]
      color_table[theme] = nil

      color_table[darker_theme]['color_bg'] = '#080808'
      color_table[darker_theme]['theme_name_full'] =
      color_table[darker_theme]['theme_name_full'] .. ' Darker'
      color_table[darker_theme]['theme_name_alt'] =
      color_table[darker_theme]['theme_name_alt'] .. '-darker'

      generate_theme(filename, darker_theme)
    end
  end
end
