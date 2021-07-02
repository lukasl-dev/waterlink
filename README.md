# waterlink

<div align="center">
  <a href="https://golang.org/">
    <img
      src="https://img.shields.io/badge/Written%20in-Go-%23EF4041?style=for-the-badge"
      height="30"
    />
  </a>
  <a href="https://pkg.go.dev/github.com/lukasl-dev/waterlink">
    <img
      src="https://img.shields.io/badge/godoc-reference-5272B4.svg?style=for-the-badge"
      height="30"
    />
  </a>
</div>

<br>

## :books: Introduction

Waterlink is a [Lavalink](https://github.com/freyacodes/Lavalink) client written in Go. **The library is based on the [Lavalink 3.x.x protocol](https://github.com/freyacodes/Lavalink/blob/master/IMPLEMENTATION.md).**

### :mag_right: Tested versions

- [x] [v3.3.2.5](https://github.com/freyacodes/Lavalink/releases/tag/3.3.2.5)

---

## :ballot_box: Installation

To use a waterlink as a Go package, you must of course have Go installed on your system.

It is assumed that you have already worked with the Go environment. If this is not the case, see [this page first](https://golang.org/doc/install).

```shell
go get -u github.com/lukasl-dev/waterlink
```

---

## :art: Structural design

### :house: Architecture

I have tried to implement my interpretation of [**Clean Architecture by Robert C. Martin**](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html). If you have any corrections or suggestions, please create an issue.

### :mosquito: Mocking

To simplify testing for the handling of the library, waterlink offers the possibility of mock implementations. The mocking library used for this is [stretchr/testify](https://github.com/stretchr/testify).
