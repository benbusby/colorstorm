## Installation
#### Vundle
```
Plugin 'benbusby/earthbound-themes', {'rtp': 'vim/'}
```

## NOTE
These vim color schemes make use of extended syntax highlighting that vim does not provide by default (in particular, function highlighting).

If you want this extended highlighting, copy the ```extended-syntax.vim``` file to your home directory, for example, and include the following snippet somewhere in your .vimrc:

```vim
au BufEnter * :source ~/extend-syntax.vim
```

This will allow you to see the updated syntax highlighting in each theme without having to reload the theme each time or make any modifictions to the existing vim syntax files.
