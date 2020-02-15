


<img src="images/logo.png" alt=")" width="50%" >

[![Go Report Card](https://goreportcard.com/badge/github.com/pr4k/howto)](https://goreportcard.com/report/github.com/pr4k/howto)
![build](https://travis-ci.com/pr4k/howto.svg?branch=master)
![Issues](https://img.shields.io/github/issues/pr4k/howto)
![Forks](https://img.shields.io/github/forks/pr4k/howto)
![Stars](https://img.shields.io/github/stars/pr4k/howto)

---
howto is a terminal client for getting stackoverflow Answers for those who are constantly googling for doing basic programming tasks.
Now it uses both google as well as Stackoverflow to get the results, because lets agree, google's search algorithm is way better than stackoverflow's search algorithm.

It is inspired by [Howdoi](https://github.com/gleitz/howdoi) and is purely written in Go

![Terminal](images/terminal.gif)

---

# Installation

```
go get github.com/pr4k/howto 
```

For installing the package Use:
```
go install github.com/pr4k/howto 
```
then just do:

```howto <Your-Query>:mode <google/stackoverflow>```
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
