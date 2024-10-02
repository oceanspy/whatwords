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

## Customization

You can customize the predefined words by modifying the files in `~/.oceanspy/whatwords` directory

## Usage

```bash
cat myFile.txt | whatwords
```
