<p align="center">
  <img width="700" src="https://raw.githubusercontent.com/benbusby/colorstorm/main/img/colorstorm.svg">
</p>

:art: *A command line tool to generate color themes for editors (Vim, VSCode, Sublime, Atom) and terminal emulators (iTerm2, Hyper).*

[![MIT License](https://img.shields.io/github/license/benbusby/colorstorm.svg)](http://opensource.org/licenses/MIT)
[![GitHub release](https://img.shields.io/github/release/benbusby/colorstorm.svg)](https://github.com/benbusby/colorstorm/releases/)
[![build](https://github.com/benbusby/colorstorm/workflows/build/badge.svg)](https://github.com/benbusby/colorstorm/actions)

___

Contents
1. [Install](#install)
1. [Usage](#usage)
1. [Creating Themes](#creating-themes)
1. [Screenshots](#screenshots)

## Install

### Binary Packages

In progress, check back soon!

### From Source

- Install [Zig](https://github.com/ziglang/zig/wiki/Install-Zig-from-a-Package-Manager)
- Run: `make release`
- Move `zig-out/bin/colorstorm` into your `PATH`

## Usage

```bash
$ colorstorm [-o outdir] [-g generator] input

-o|--outdir: The directory to output themes to (default: "./colorstorm-out")
-g|--gen:    Generator type (default: all)
             Available types: all, atom, vscode, vim, sublime, iterm, hyper
-i|--input:  The JSON input file to use for generating the themes
             See: https://github.com/benbusby/colorstorm#creating-themes
```

#### Supported Editors
- Vim
- VSCode
- Sublime
- Atom

#### Supported Terminal Emulators
- iTerm2
- Hyper

## Creating Themes

You can create themes for all available editors and terminal emulators using a
single JSON file to define the colors. The file should be an array (even for
one theme), with the following structure:

```json
[
    {
        "theme_name_full": "Moonside",
        "theme_name_safe": "moonside",
        "color_bg_main": "#000000",
        "color_bg_alt1": "#080808",
        "color_bg_alt2": "#131313",
        "color_fg": "#ffffff",
        "color_linenr": "#9e5dc8",
        "color_select": "#5a1359",
        "color_type": "#f6f929",
        "color_accent": "#fd35fa",
        "color_string": "#ff6693",
        "color_boolean": "#fd9935",
        "color_variable": "#c67ff4",
        "color_number": "#aaef64",
        "color_comment": "#7ca454",
        "color_function": "#5e9aff"
    },
    {
        ...
    }
]
```

Value names are mostly self-explanatory, but here is a breakdown of what each field means:

| Field | Explanation |
| -- | -- |
| `theme_name_full` | The full name of the theme that will appear in theme file documentation |
| `theme_name_safe` | The value to use as the filename for the theme |
| `color_bg_main`   | Primary background color |
| `color_bg_alt1`   | A separate background color to use for UI elements like file trees and tab bars |
| `color_bg_alt2`   | A separate background color to use for UI elements like line numbers and gutters |
| `color_fg`        | The foreground color (all generic text) |
| `color_linenr`    | The color used for line numbers |
| `color_select`    | The color used for selecting a word or lines of text |
| `color_type`      | The color used for variable types (int, float, etc) |
| `color_accent`    | An "accent" color -- typically used for special cases (like current line number highlight or badge backgrounds) |
| `color_string`    | The color used for strings |
| `color_boolean`   | The color used for boolean values |
| `color_variable`  | The color used for variable instances and constants |
| `color_number`    | The color used for numeric values |
| `color_comment`   | The color used for code comments |
| `color_function`  | The color used for function names |

## Screenshots

- [Earthbound Themes](https://github.com/benbusby/earthbound-themes)

[![Vim Installs](https://img.shields.io/static/v1?label=vim&message=a lot&color=green&logo=vim)](https://www.vim.org/scripts/script.php?script_id=5920)
[![VSCode Installs](https://img.shields.io/visual-studio-marketplace/i/benbusby.earthbound-themes?label=vscode&color=4444ff&logo=visual-studio-code)](https://marketplace.visualstudio.com/items?itemName=benbusby.earthbound-themes)
[![Package Control](https://img.shields.io/packagecontrol/dt/Earthbound%20Themes?color=ff4500&label=sublime&logo=sublime-text)](https://packagecontrol.io/packages/Earthbound%20Themes)
[![APM](https://img.shields.io/apm/dm/earthbound-themes-syntax?color=dark-green&label=atom&logo=atom)](https://atom.io/packages/earthbound-themes-syntax)

### Earthbound

![Earthbound Screenshot](img/screenshots/earthbound.png)

### Moonside

![Moonside Screenshot](img/screenshots/moonside.png)

### Zombie Threed

![Zombie Threed Screenshot](img/screenshots/threed.png)

### Fire Spring

![Fire Spring Screenshot](img/screenshots/fire_spring.png)

### Devil's Machine

![Devil's Machine Screenshot](img/screenshots/devils_machine.png)

### Dusty Dunes

![Dusty Dunes Screenshot](img/screenshots/dusty_dunes.png)

### Magicant (Light Theme)

![Magicant Screenshot](img/screenshots/magicant.png)

### Cave of the Past (Monochrome)

![Cave of the Past Screenshot](img/screenshots/cave_of_the_past.png)
