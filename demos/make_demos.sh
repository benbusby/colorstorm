#!/bin/sh

export PS1="$ "

rm -f ./colorstorm
ln -s ../colorstorm .

colorstorm() {
    ./colorstorm "$@"
}

export -f colorstorm

if [ "$#" -eq 1 ]; then
    rm -f ./"$1".gif
    vhs "$1" -o "$1".gif
else
    rm -f ./*.gif
    for i in *.tape; do
        [ -f "$i" ] || break
        vhs "$i" -o "$i".gif
    done
fi

rm ./colorstorm
