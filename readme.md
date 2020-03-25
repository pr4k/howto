

<p align="center">
    <img src="images/logo.png" width="490" alt"howto" >
    <br>
    <img src="https://goreportcard.com/badge/github.com/pr4k/howto"
        alt="GoReport">
    <img src="https://travis-ci.com/pr4k/howto.svg?branch=master" alt="Build">
    <img src="https://img.shields.io/github/stars/pr4k/howto" alt="Stars">
    <img src="https://img.shields.io/github/issues/pr4k/howto" alt="Issues">
    <img src="https://img.shields.io/github/forks/pr4k/howto" alt="Forks">
</p>
<p align="center">
    <a href="#installation">Installation</a> •
    <a href="#usage">Usage</a> •
    <a href="#features">Features</a> •
    <a href="#to-do">To-Do</a> •
    <a href="#license">License</a>
</p>

---
howto is a terminal client for getting stackoverflow Answers for those who are constantly googling for doing basic programming tasks.
Now it uses both google as well as Stackoverflow to get the results, because lets agree, google's search algorithm is way better than stackoverflow's search algorithm.

It is inspired by [Howdoi](https://github.com/gleitz/howdoi) and is purely written in Go.

Read About it at *[Medium](https://medium.com/better-programming/how-i-use-stackoverflow-with-just-a-terminal-go-17548716ab61?source=---------2------------------)*

[![asciicast](https://asciinema.org/a/Fh5xrpejzh2miP88NZtLED5gm.svg)](https://asciinema.org/a/Fh5xrpejzh2miP88NZtLED5gm)

---

# Installation

Its simply go get to install
```
go get github.com/pr4k/howto 

```

For installing the package Use:

```bash
go install github.com/pr4k/howto 
```
---
# Usage


```howto <Your-Query>:<google/stackoverflow>```
Note:- if mode is not specified then the default mode is google.

---
# Features
 - Uses Go Routines to parallely scrape answers so time complexity is independent of number of solution
 - Provides Terminal ui to navigate and access answers
 - Uses google's result along with stackoverflow's result.

# To-Do
- Add google results along with stackoverflow results (Done)
- Implement Syntax Highlighting for code parts

---

## License

[![License](https://img.shields.io/github/license/pr4k/howto)](http://badges.mit-license.org)

- **[MIT license](http://opensource.org/licenses/mit-license.php)**
- Copyright 2020 © pr4k
---
