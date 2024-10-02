# WhatWords

WhatWords is a simple CLI tool to parse documents or URL to find most used words and predefined words

## Pre-requisites

- Go 1.21 or higher

## Installation

```bash
git clone git@github.com/oceanspy/whatwords.git
cd whatwords
./install.sh
```

## Update

```bash
cd whatwords
git pull origin main
./update.sh
```

## Uninstall

```bash
rm -rf whatwords
rm -f /usr/local/bin/whatwords
rm -rf ~/.oceanspy/whatwords
```

## Customization

You can customize the predefined words by modifying the files in `~/.oceanspy/whatwords` directory

## Usage

```bash
cat myFile.txt | whatwords
```
