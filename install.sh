#!/bin/bash

go build -o build/whatwords whatwords.go

mkdir -p $HOME/.oceanspy/whatwords
cp data/*.csv $HOME/.oceanspy/whatwords

# Ask to create a symlink
echo "Do you want to create a symlink to /usr/local/bin/whatwords? (y/n)"
read -r answer
if [ "$answer" != "${answer#[Yy]}" ]; then
    sudo ln -s $(PWD)/build/whatwords /usr/local/bin/whatwords
fi
