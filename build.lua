#!/usr/bin/lua
-- Lua script for automatically building theme files
-- for vim, vscode, sublime, and atom.

USAGE = [[
./test.lua [vim, vscode, sublime, atom]
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
  earthbound=1,
  threed=1,
  fire_spring=1,
  dusty_dunes=1
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
    color_type='#bf3000',
    color_accent='#fafd51',
    color_string='#fb967f',
    color_number='#bcfaaa',
    color_boolean='#70caff',
    color_comment='#aaaaaa',
    color_variable='#abdcdc',
    color_function='#70caff'
  },
  moonside={
    theme_id=uuid(),
    theme_name_full='Moonside',
    theme_name_alt='moonside',
    color_bg_main='#000000',
    color_bg_alt1='#080808',
    color_bg_alt2='#131313',
    color_fg='#ffffff',
    color_type='#fd9935',
    color_accent='#fd35fa',
    color_string='#ff6693',
    color_boolean='#f6f929',
    color_variable='#c67ff4',
    color_number='#aaef64',
    color_comment='#aaaaaa',
    color_function='#5e9aff'
  },
  threed={
    theme_id=uuid(),
    theme_name_full='Threed',
    theme_name_alt='threed',
    color_bg_main='#303454',
    color_bg_alt1='#373c60',
    color_bg_alt2='#3a3f66',
    color_fg='#f0faff',
    color_type='#ffcfcb',
    color_accent='#d4cbff',
    color_string='#c67ff4',
    color_boolean='#ffcf32',
    color_variable='#2fff89',
    color_number='#d4cbff',
    color_comment='#aaaaaa',
    color_function='#f89070'
  },
  fire_spring={
    theme_id=uuid(),
    theme_name_full='Fire Spring',
    theme_name_alt='fire-spring',
    color_bg_main='#181020',
    color_bg_alt1='#21162c',
    color_bg_alt2='#261933',
    color_fg='#ffffca',
    color_type='#e5caff',
    color_accent='#74e4f3',
    color_string='#d9c400',
    color_boolean='#ff7e50',
    color_variable='#f0e500',
    color_number='#a99ade',
    color_comment='#aaaaaa',
    color_function='#d992ff'
  },
  dusty_dunes={
    theme_id=uuid(),
    theme_name_full='Dusty Dunes',
    theme_name_alt='dusty-dunes',
    color_bg_main='#0e0900',
    color_bg_alt1='#150d00',
    color_bg_alt2='#1e1302',
    color_fg='#f9e4a1',
    color_type='#e0c364',
    color_accent='#f6d56a',
    color_string='#ffebae',
    color_boolean='#ffd03c',
    color_variable='#f6d56a',
    color_number='#f6d56a',
    color_comment='#666644',
    color_function='#f6d56a'
  },
  magicant={
    theme_id=uuid(),
    theme_name_full='Magicant (Light)',
    theme_name_alt='magicant',
    color_bg_main='#f9f8b9',
    color_bg_alt1='#efeeb2',
    color_bg_alt2='#e6e5ab',
    color_fg='#220088',
    color_type='#881200',
    color_accent='#604633',
    color_string='#7100b1',
    color_boolean='#a31100',
    color_variable='#87000d',
    color_number='#604633',
    color_comment='#888888',
    color_function='#9d02f2'
  },
  cave_of_the_past={
    theme_id=uuid(),
    theme_name_full='Cave of the Past (Monochrome)',
    theme_name_alt='cave-of-the-past',
    color_bg_main='#b0d0b8',
    color_bg_alt1='#a5c4ad',
    color_bg_alt2='#9ebba6',
    color_fg='#394537',
    color_type='#3e5f39',
    color_accent='#0e1e0e',
    color_string='#2b342a',
    color_boolean='#293d29',
    color_variable='#2b342a',
    color_number='#445046',
    color_comment='#93a096',
    color_function='#0e1e0e'
  }
}

color_files = {
  vim='templates/template.vim',
  vscode='templates/template.json',
  sublime='templates/earthbound_template.tmTheme',
  atom='templates/colors.less'
}

out_paths = {
  vim='vim/colors/%s.vim',
  vscode='vscode/themes/%s.json',
  sublime='sublime/earthbound_%s.tmTheme',
  atom='atom/themes/%s-syntax/colors.less'
}

-- THEME GENERATION

--- Generates a formatted theme file
-- A string replacement is performed for each line of (file)
-- using the values in color_table.
-- @param file: The template file to use
-- @param theme: The theme to generate a file for
function generate_theme(file, name, theme, out_path)
  local theme_file = io.open(file, 'r')

  if theme_file == nil then
    print('Unable to open ' .. file)
    return
  end

  local lines = {}
  for line in theme_file:lines() do
    for k,v in pairs(theme) do
      line = string.gsub(line, k, v)
    end

    lines[#lines + 1] = line
  end
  theme_file:close()

  out_path = string.format(out_path, theme['theme_name_alt'])
  print(out_path)

  theme_file = io.open(out_path, 'w')
  for i,line in ipairs(lines) do
    theme_file:write(line, "\n")
  end
  theme_file:close()
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
  local editor = arg[1]
  local filename = color_files[editor]
  print('=== Generating theme files for ' .. arg[1])

  for theme, value in pairs(color_table) do
    generate_theme(filename, theme, color_table[theme], out_paths[editor])

    -- A few themes can use darker variants, which replaces the background
    -- with #080808
    if darker_variants[theme] ~= nil then
      local darker_name = theme .. '_darker'

      dark_theme = {}
      for key, val in pairs(color_table[theme]) do
        dark_theme[key] = val
      end

      dark_theme['color_bg_main'] = '#080808'
      dark_theme['theme_name_full'] = dark_theme['theme_name_full'] .. ' Darker'
      dark_theme['theme_name_alt'] = dark_theme['theme_name_alt'] .. '-darker'

      generate_theme(filename, darker_name, dark_theme, out_paths[editor])
    end
  end
end
