## Installation
#### vim-plug
```
Plug 'benbusby/earthbound-themes', {'rtp': 'vim/', 'branch': 'vim' }
```

#### Vundle
Note: Since Vundle doesn't allow checking out plugins by branch, you'll need to navigate to the plugin's directory (`~/.vim/bundle/earthbound-themes/`) and check out the `vim` branch to view the color schemes.
```
Plugin 'benbusby/earthbound-themes', {'rtp': 'vim/'}
```

#### Manually
You can either:
1. Clone the repo
2. Check out the `vim` branch
3. Extract the contents of `vim/colors` to your syntax folder

or [download the latest zip release](https://github.com/benbusby/earthbound-themes/releases) and extract the contents to your colors folder.

## NOTE
These vim color schemes make use of extended syntax highlighting that vim does not provide by default (in particular, function highlighting).

If you want this extended highlighting, copy the ```extended-syntax.vim``` file to your home directory, for example, and include the following snippet somewhere in your .vimrc:

```vim
au BufEnter * :source ~/extend-syntax.vim
```

This will allow you to see the updated syntax highlighting in each theme without having to reload the theme each time or make any modifictions to the existing vim syntax files.
