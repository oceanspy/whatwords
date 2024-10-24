#!/bin/bash

if command -v go &>/dev/null; then
    echo "Go is installed"
else
    echo "Go is not installed"
    echo "Please install Go and try again"
    exit 1
fi

go build -o build/whatwords whatwords.go

echo "Do you want to copy the default configuration to ~/.oceanspy/whatwords ? (y/n)"
read -r answerCopy
if [ "$answerCopy" != "${answerLn#[Yy]}" ]; then
    mkdir -p $HOME/.oceanspy/whatwords
    cp data/*.csv $HOME/.oceanspy/whatwords
fi

# Ask to create a symlink
echo "Do you want to create a symlink to /usr/local/bin/whatwords? (y/n)"
read -r answerLn
if [ "$answerLn" != "${answerLn#[Yy]}" ]; then
    if command -v sudo &>/dev/null; then
        sudo ln -s $PWD/build/whatwords /usr/local/bin/whatwords
    else
        ln -s $PWD/build/whatwords /usr/local/bin/whatwords
    fi
fi
