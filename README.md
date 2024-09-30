# WhatWords

WhatWords is a simple CLI tool to parse documents or URL to find most used words and predefined words

## Pre-requisites

- Go 1.16 or higher

## Installation

```bash
git clone git@github.com/w9nz/whatwords.git
cd whatwords
go build -o build/whatwords whatwords.go
ln -s /xxxxx/build/whatwords /usr/local/bin/whatwords
```

## Usage

```bash
cat myFile.txt | whatwords
```
