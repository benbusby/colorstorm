#!/usr/bin/lua
-- Lua script for automatically building theme files
-- for vim, vscode, sublime, and atom.

USAGE = [[
./test.lua [vim, vscode, sublime, atom, iterm2]
]]

-- COLOR DEFINITIONS

color_table = {
  earthbound={
    theme_name_full='Earthbound',
    theme_name_alt='earthbound',
    color_bg_main='#360a5f',
    color_bg_alt1='#2b044f',
    color_bg_alt2='#1c0037',
    color_fg='#ffffff',
    color_linenr='#94b2b2',
    color_select='#6a6c23',
    color_type='#f56b3c',
    color_accent='#fafd51',
    color_string='#84fc60',
    color_number='#fb967f',
    color_boolean='#70caff',
    color_comment='#acaf6f',
    color_variable='#abdcdc',
    color_function='#70caff'
  },
  moonside={
    theme_name_full='Moonside',
    theme_name_alt='moonside',
    color_bg_main='#000000',
    color_bg_alt1='#080808',
    color_bg_alt2='#131313',
    color_fg='#ffffff',
    color_linenr='#9e5dc8',
    color_select='#5a1359',
    color_type='#f6f929',
    color_accent='#fd35fa',
    color_string='#ff6693',
    color_boolean='#fd9935',
    color_variable='#c67ff4',
    color_number='#aaef64',
    color_comment='#7ca454',
    color_function='#5e9aff'
  },
  threed={
    theme_name_full='Threed',
    theme_name_alt='threed',
    color_bg_main='#303454',
    color_bg_alt1='#373c60',
    color_bg_alt2='#2a2e4a',
    color_fg='#f0faff',
    color_linenr='#9590b2',
    color_select='#534f63',
    color_type='#ffcfcb',
    color_accent='#d4cbff',
    color_string='#ffcf32',
    color_boolean='#c67ff4',
    color_variable='#2fff89',
    color_number='#d4cbff',
    color_comment='#bdb7db',
    color_function='#f89070'
  },
  fire_spring={
    theme_name_full='Fire Spring',
    theme_name_alt='fire-spring',
    color_bg_main='#261933',
    color_bg_alt1='#21162c',
    color_bg_alt2='#181020',
    color_fg='#ffffca',
    color_linenr='#b49a19',
    color_select='#632611',
    color_type='#ff7e50',
    color_accent='#f0e500',
    color_string='#74e4f3',
    color_boolean='#d9c400',
    color_variable='#e5caff',
    color_number='#a99ade',
    color_comment='#bb8673',
    color_function='#d992ff'
  },
  dusty_dunes={
    theme_name_full='Dusty Dunes',
    theme_name_alt='dusty-dunes',
    color_bg_main='#1e1b07',
    color_bg_alt1='#150d00',
    color_bg_alt2='#140f00',
    color_fg='#f9e4a1',
    color_linenr='#f9e4a1',
    color_select='#6b5e33',
    color_type='#e0c364',
    color_accent='#f6d56a',
    color_string='#ffebae',
    color_boolean='#ffd03c',
    color_variable='#f6d56a',
    color_number='#f6d56a',
    color_comment='#aaaa88',
    color_function='#f6d56a'
  },
  magicant={
    theme_name_full='Magicant (Light)',
    theme_name_alt='magicant',
    color_bg_main='#f9f8b9',
    color_bg_alt1='#efeeb2',
    color_bg_alt2='#e6e5ab',
    color_fg='#220088',
    color_linenr='#876a55',
    color_select='#b299ff',
    color_type='#881200',
    color_accent='#604633',
    color_string='#7100b1',
    color_boolean='#a31100',
    color_variable='#87000d',
    color_number='#604633',
    color_comment='#525252',
    color_function='#9d02f2'
  },
  cave_of_the_past={
    theme_name_full='Cave of the Past (Monochrome)',
    theme_name_alt='cave-of-the-past',
    color_bg_main='#b0d0b8',
    color_bg_alt1='#a5c4ad',
    color_bg_alt2='#9ab5a2',
    color_fg='#262e25',
    color_linenr='#315b31',
    color_select='#7c9283',
    color_type='#3e5f39',
    color_accent='#0e1e0e',
    color_string='#2b342a',
    color_boolean='#293d29',
    color_variable='#2b342a',
    color_number='#445046',
    color_comment='#5b5f59',
    color_function='#0e1e0e'
  },
  devils_machine={
    theme_name_full='Devils Machine',
    theme_name_alt='devils-machine',
    color_bg_main='#040001',
    color_bg_alt1='#260000',
    color_bg_alt2='#170000',
    color_fg='#ffcccc',
    color_linenr='#a45a52',
    color_select='#481200',
    color_type='#c15bac',
    color_accent='#ffcc99',
    color_string='#ddaaaa',
    color_boolean='#938198',
    color_variable='#ff8c69',
    color_number='#b4a0dc',
    color_comment='#a59e85',
    color_function='#e6817e'
  }
}

darker_variants = {
  earthbound=1,
  threed=1,
  fire_spring=1,
  dusty_dunes=1
}

use_percentages = {
  iterm2=1
}

color_files = {
  vim='templates/template.vim',
  vscode='templates/template.json',
  sublime='templates/earthbound_template.tmTheme',
  atom='templates/colors.less',
  iterm2='templates/template.itermcolors',
  all=''
}

out_paths = {
  vim='vim/colors/%s.vim',
  vscode='vscode/themes/%s.json',
  sublime='sublime/earthbound_%s.tmTheme',
  atom='atom/themes/%s-syntax/colors.less',
  iterm2='iterm2/%s.itermcolors'
}

-- MISC UTILS

local random = math.random
local function uuid()
  local template ='xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'
  return string.gsub(template, '[xy]', function (c)
    local v = (c == 'x') and random(0, 0xf) or random(8, 0xb)
    return string.format('%x', v)
  end)
end

local function hex_to_percent(color)
  local colorR = string.sub(color, 1, 2)
  local colorG = string.sub(color, 3, 4)
  local colorB = string.sub(color, 5, 6)

  return {
    R=tonumber(string.format('0x%s', colorR)) / 255,
    G=tonumber(string.format('0x%s', colorG)) / 255,
    B=tonumber(string.format('0x%s', colorB)) / 255
  }
end

-- THEME GENERATION

--- Generates a formatted theme file
-- A string replacement is performed for each line of (file)
-- using the values in color_table.
-- @param file: The template file to use
-- @param theme: The theme to generate a file for
local function generate_theme(file, name, theme, editor)
  local theme_file = io.open(file, 'r')

  if theme_file == nil then
    print('Unable to open ' .. file)
    return
  end

  local lines = {}
  for line in theme_file:lines() do
    for k, v in pairs(theme) do
      if use_percentages[editor] ~= nil and string.find(v, '^#') then
        local colorRGB = hex_to_percent(string.sub(v, 2))
        line = string.gsub(line, 'R' .. k, colorRGB['R'])
        line = string.gsub(line, 'G' .. k, colorRGB['G'])
        line = string.gsub(line, 'B' .. k, colorRGB['B'])
      else
        line = string.gsub(line, k, v)
      end
    end

    lines[#lines + 1] = line
  end
  theme_file:close()

  out_path = string.format(out_paths[editor], theme['theme_name_alt'])
  print(out_path)

  theme_file = io.open(out_path, 'w')
  for i, line in ipairs(lines) do
    theme_file:write(line, "\n")
  end
  theme_file:close()
end

--- Runs all necessary theme generation functions for a specific editor
-- @param editor: The editor to use for theme generation
local function create_editor_themes(editor)
  local filename = color_files[editor]
  print('=== Generating theme files for ' .. editor)

  for theme, value in pairs(color_table) do
    color_table[theme]['uuid'] = uuid()
    generate_theme(filename, theme, color_table[theme], editor)

    -- A few themes can use "-darker" variants, which swaps the default
    -- background (color_bg_main) and the darker background (color_bg_alt2)
    if darker_variants[theme] ~= nil then
      local darker_name = theme .. '_darker'

      dark_theme = {}
      for key, val in pairs(color_table[theme]) do
        dark_theme[key] = val
      end

      dark_theme['uuid'] = uuid()
      dark_theme['color_bg_main'] = color_table[theme]['color_bg_alt2']
      dark_theme['color_bg_alt2'] = color_table[theme]['color_bg_main']
      dark_theme['theme_name_full'] = dark_theme['theme_name_full'] .. ' Darker'
      dark_theme['theme_name_alt'] = dark_theme['theme_name_alt'] .. '-darker'

      generate_theme(filename, darker_name, dark_theme, editor)
    end
  end
end

-- CLI

if arg[1] == nil then
  print('Error: Argument required')
  print(USAGE)
  os.exit(1)
elseif color_files[arg[1]] == nil then
  print('Error: Invalid argument')
  print(USAGE)
  os.exit(1)
else
  -- Get template file for theme generation
  if arg[1] == 'all' then
    for editor, _ in pairs(out_paths) do
      create_editor_themes(editor)
    end
  else
    create_editor_themes(arg[1])
  end
end
