[![Build Status](https://travis-ci.org/zhexuany/wordGenerator.png?branch=master)](https://travis-ci.org/zhexuany/wordGenerator)
[![GoDoc](https://godoc.org/github.com/zhexuany/wordGenerator?status.svg)](https://godoc.org/github.com/zhexuany/wordGenerator)
#Word Generator
It is just an implementation for generating randomized character, word and words. 
The goal of creating such project is that I need generate a lot of random hash key 
in order to test the correctness of anohter project that I am working on. 

## Using

### import

`import "github.com/zhexuany/wordGenerator"`

### API

~~~go
func foo() {
    ch := wordGenerator.GetChar()
    // 5 is the length of word
    str := wordGenerator.GetWord(5)
    // 5 is the number of word, 20 is the maximum length of word
    strs := wordGenerator.GetWords(5, 20)
}
~~~

## Dependency
This project is built and tested under `Go` 1.7. Other `Go` version  may not supported yet. Other than `GO` itself,
this project does not have any deoendency.

## Future work
- [ ] support Non-ascii such Chinese and Japanese
