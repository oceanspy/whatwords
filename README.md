# WhatWords

WhatWords is a simple CLI tool to parse documents or URL to find most used words and predefined words

![alt text](https://github.com/oceanspy/whatwords/blob/main/whatwords-screenshot0.png?raw=true)
![alt text](https://github.com/oceanspy/whatwords/blob/main/whatwords-screenshot1.png?raw=true)

## Pre-requisites

- Go 1.21 or higher

## Usage

```bash
cat myFile.txt | whatwords 
cat myFile.txt | whatwords 50 
```

__50__ is the number of words to display. Default is 20.

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

### Custom Words

You can also have a custom list of words to be searched for in the document. You can add the words in `customwords.csv` file in the `~/.oceanspy/whatwords` directory.

All words need to be in lowercase, one word per line.

Example:

```csv
word1
word2
...
```

Words should not have any special characters or spaces.

If you want to add a word with special characters, you can transform them first in `similarwords.csv` file in the `~/.oceanspy/whatwords` directory. Cf next section.

### Similar Words

You can also have a list of words that are similar to the predefined words. You can add the words in `similarwords.csv` file in the `~/.oceanspy/whatwords` directory.

All words need to be in lowercase, one word per line. The `+` character is used to separate the words.

Example:

```csv
js,javasscript
c++,cpp
...
```

All occurrences of the first word will be replaced by the second word.

### Excluded Words

You can also have a list of words that should be excluded from the search. You can add the words in `excludedwords.csv` file in the `~/.oceanspy/whatwords` directory.

All words need to be in lowercase, one word per line

Example:

```csv
and
the
...
```

### Short Words to keep

By default, all words with less than 3 characters are excluded from the search. You can change this by modifying the `shortwords.csv` file in the `~/.oceanspy/whatwords` directory.

All words need to be in lowercase, one word per line.

Example:

```csv
a
an
...
```

### Multiple Words

You can also have a list of multiple words that should be considered as one word. You can add the words in `multiplewords.csv` file in the `~/.oceanspy/whatwords` directory.

All words need to be in lowercase, one word per line. The words should be separated by a `+` character.

Example:

```csv
new york,new+york
...
```

All occurrences of the first word will be replaced by the second word. The `+` character is used to separate the words.
