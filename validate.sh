#!/bin/bash

set -e

# VSCode
echo -e "\n=== VSCode Tests"
for file in vscode/themes/*.json; do
    echo -n "  ? $file: "
    cat $file | jq . > /dev/null
    echo -e "\r  ✓ $file "
done

# Sublime
echo -e "\n=== Sublime Tests"
for file in sublime/*.tmTheme; do
    echo -n "  ? $file: "
    xmllint --noout $file
    echo -e "\r  ✓ $file "
done

# Vim
echo -e "\n=== Vim Tests"
for file in vim/colors/*.vim; do
    rm -f vimout
    echo -n "  ? $file: "
    vim -c "let v:errmsg=''" -c "so $file" -c "call setline('.', v:errmsg)" -c "w vimout" -c "q"
    [[ $(wc -c <"vimout") -ge 2 ]] && cat vimout && rm -f vimout && exit 1
    echo -e "\r  ✓ $file "
done

# Atom
echo -e "\n=== Atom Tests"
for file in atom/themes/*-syntax/*.less; do 
    echo -n "  ? $file: "
    cp $file .
    lessc atom/styles/base.less out.css
    echo -e "\r  ✓ $file "
done

# (Cleanup)
echo ""
rm out.css
rm colors.less
rm -f vimout
