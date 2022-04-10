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
- [⛵ Connection related](#-connection-related)
  - [Opening a connection](#opening-a-connection)
  - [🦷 Configuring session resuming](#-configuring-session-resuming)
  - [❌ Disabling session resuming](#-disabling-session-resuming)
  - [📜 Getting a guild](#-getting-a-guild)
- [🏠 Guild related](#-guild-related)
  - [Destroying its player](#destroying-its-player)
  - [Updating its voice server](#updating-its-voice-server)
  - [Playing a track](#playing-a-track)
  - [Stopping the playback](#stopping-the-playback)
  - [Pausing/Resuming the playback](#pausingresuming-the-playback)
  - [Seeking the playback](#seeking-the-playback)
  - [Updating the volume](#updating-the-volume)
- [📂 Examples](#-examples)

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

## ⛵ Connection related

### Opening a connection

The `opts` parameter is optional. In this example, it is used to register an EventHandler. **If this is not needed, omit it.**

```go
creds := waterlink.Credentials{
  Authorization: "youshallnotpass", // password of the Lavalink instance
  UserID:        0,                 // id of the bot user
}

opts := waterlink.ConnectionOptions{
  EventHandler: waterlink.EventHandlerFunc(func(evt interface{}) {
    fmt.Printf("%s received\n", reflect.TypeOf(evt))
  }),
}

conn, err := waterlink.Open("ws://localhost:2333", creds, opts)
```

To restore a past session, its resume key can be defined in the credentials.

> See  [Configuring session resuming](#-configuring-session-resuming)

```go
creds := waterlink.Credentials{
  Authorization: "youshallnotpass", // password of the Lavalink instance
  UserID:        0,                 // id of the bot user
  ResumeKey:     "myResumeKey",     // the resume key of the previous session
}
```

### 🦷 Configuring session resuming

Configures a resume key with a timeout of 5 minutes.

```go
err := conn.ConfigureResuming("myResumeKey", 5*time.Minute)
```

### ❌ Disabling session resuming

```go
err := conn.DisableResuming()
```

### 📜 Getting a guild

A guild is necessary to access its audio player. **The function does not check whether the bot user is on this guild.**

```go
g := conn.Guild(0) // id of the guild to access
```

---

## 🏠 Guild related

A guild can be obtained via one's own ID with the use of a connection.

> See [Getting a guild](#-getting-a-guild)

### Destroying its player

```go
err := g.Destroy()
```

### Updating its voice server

This function is primarily performed by 3rd party libraries event listeners.

> See [Examples](#-examples)

```go
err := g.UpdateVoice("session", "token", "endpoint")
```

### Playing a track

The `params` parameter is optional. It can be used to specify more information. **If this is not needed, omit it.**

```go
params := waterlink.PlayParams{
  StartTime: 0,
  EndTime:   0,
  Volume:    0,
  NoReplace: false,
  Pause:     false,
}
err := g.PlayTrack(t, params)
```

### Stopping the playback

```go
err := g.Stop()
```

### Pausing/Resuming the playback

```go
err := g.SetPaused(true)
```

### Seeking the playback

```go
err := g.Seek(25 * time.Second) // seek to 25 seconds
```

### Updating the volume

```go
err := g.UpdateVolume(25) // 25%
```

---

## 📂 Examples
