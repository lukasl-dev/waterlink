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
  <a href="https://goreportcard.com/report/github.com/lukasl-dev/waterlink">
    <img
      src="https://goreportcard.com/badge/github.com/lukasl-dev/waterlink?style=for-the-badge"
      height="30"
    />
  </a>
</div>

<br />

## 📖 Table of Contents

- [📖 Table of Contents](#-table-of-contents)
- [📦 Installation](#-installation)
- [🍀 Getting started](#-getting-started)
- [✨ Client related](#-client-related)
  - [Creating a client](#creating-a-client)
  - [Loading tracks](#loading-tracks)
  - [Decoding a single track](#decoding-a-single-track)
  - [Decoding multiple tracks](#decoding-multiple-tracks)

---

## 📦 Installation

```shell
go get -u github.com/lukasl-dev/waterlink
```

---

## 🍀 Getting started

For the further guides, a Lavalink instance is used, which uses the following `application.yml` configuration:

```yml
server:
  port: 2333
  address: 0.0.0.0
lavalink:
  server:
    password: "youshallnotpass"
```

---

## ✨ Client related

### Creating a client

```go
creds := waterlink.Credentials{
  Authorization: "youshallnotpass",
}
client, err := waterlink.NewClient("http://localhost:2333", creds)
```

### Loading tracks

```go
res, err := client.LoadTracks(query.Of("https://www.youtube.com/watch?v=dQw4w9WgXcQ"))
```

```go
res, err := client.LoadTracks(query.YouTube("Never Gonna Give You Up"))
```

```go
res, err := client.LoadTracks(query.SoundCloud("Never Gonna Give You Up"))
```

### Decoding a single track

```go
info, err := client.DecodeTrack(
  "QAAAoQIAPFJpY2sgQXN0bGV5IC0gTmV2ZXIgR29ubmEgR2l2ZSBZb3UgVXAgKE9mZmljaWFsIE11c2ljIFZpZGVvKQALUmljayBBc3RsZXkAAAAAAANACAALZFF3NHc5V2dYY1EAAQAraHR0cHM6Ly93d3cueW91dHViZS5jb20vd2F0Y2g/dj1kUXc0dzlXZ1hjUQAHeW91dHViZQAAAAAAAAAA",
)
```

### Decoding multiple tracks

```go
tracks, err := client.DecodeTracks([]string{
  "QAAAoQIAPFJpY2sgQXN0bGV5IC0gTmV2ZXIgR29ubmEgR2l2ZSBZb3UgVXAgKE9mZmljaWFsIE11c2ljIFZpZGVvKQALUmljayBBc3RsZXkAAAAAAANACAALZFF3NHc5V2dYY1EAAQAraHR0cHM6Ly93d3cueW91dHViZS5jb20vd2F0Y2g/dj1kUXc0dzlXZ1hjUQAHeW91dHViZQAAAAAAAAAA",
})
```

---
