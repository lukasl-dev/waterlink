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

## :books: Introduction

Waterlink is a [Lavalink](https://github.com/freyacodes/Lavalink) client written in Go. **The library is based on
the [Lavalink 3.x.x protocol](https://github.com/freyacodes/Lavalink/blob/master/IMPLEMENTATION.md).**

---

## :mag_right: Compatibility

- [x] [v3.3.2.5](https://github.com/freyacodes/Lavalink/releases/tag/3.3.2.5)

---

## :ballot_box: Installation

It is assumed that you have already worked with the Go environment. If this is not the case,
see [this page first](https://golang.org/doc/install).

```shell
go get -u github.com/lukasl-dev/waterlink
```

---

## :art: Structural design

### :house: Architecture

I have tried to implement my interpretation of [**Clean Architecture by Robert C. Martin (Uncle
Bob)**](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html). If you have any corrections or
suggestions, please create an issue.

### :mosquito: Mocking

To simplify testing for the handling of the library, waterlink offers the possibility of mock implementations. The
mocking library used for this is [stretchr/testify](https://github.com/stretchr/testify).

---

## :bamboo: Getting started

Firstly, we need to differentiate between **connectionless** and **connection-oriented** use cases. **Connection-oriented** use cases require an **active web socket connection** to the Lavalink server and **connectionless** use cases are **only based on simple HTTP requests**.

### :boat: Opening a connection

<details>
  <summary>Usage</summary>
  <p>
  
  ```go
  package main
  
  import (
    "context"
    "net/url"
  
    "github.com/lukasl-dev/waterlink"
  )
  
  var (
    host = url.URL{                // TODO: adjust
      Scheme: "ws",
      Host:   "localhost:2333",
    }
    passphrase = "youshallnotpass" // TODO: adjust
  )
  
  func main() {
    opts := waterlink.NewConnectOptions().WithPassphrase(passphrase)
    conn, err := waterlink.Connect(context.TODO(), host, opts)
    if err != nil {
      // TODO: handle error
      return
    }
    // TODO: use conn
  }
  ```

  </p>
</details>

### :phone: Create a requester

<details>
  <summary>Usage</summary>
  <p>
  
  ```go
  package main
  
  import (
    "net/url"
  
    "github.com/lukasl-dev/waterlink"
  )
  
  var (
    host = url.URL{                // TODO: adjust
      Scheme: "http",
      Host:   "localhost:2333",
    }
    passphrase = "youshallnotpass" // TODO: adjust
  )
  
  func main() {
    opts := waterlink.NewRequesterOptions().WithPassphrase(passphrase)
    req := waterlink.NewRequester(host, opts)
    // TODO: use req
  }
  ```

  </p>
</details>

### :musical_keyboard: Interacting with tracks

#### :kite: Loading multiple tracks

<details>
  <summary>Usage</summary>
  <p>
  
  ```go
  package main

  import (
    "github.com/lukasl-dev/waterlink"
  )

  var (
    req        waterlink.Requester                             // TODO: create req
    identifier = "https://www.youtube.com/watch?v=dQw4w9WgXcQ" // TODO: adjust
  )

  func main() {
    resp, err := req.LoadTracks(identifier)
    if err != nil {
      // TODO: handle error
      return
    }
    // TODO: use resp
  }
  ```

  </p>
</details>

#### :ticket: Decoding multiple tracks

<details>
  <summary>Usage</summary>
  <p>
  
  ```go
  package main
  
  import (
    "github.com/lukasl-dev/waterlink"
  )
  
  var (
    req      waterlink.Requester // TODO: create req
    trackIDs []string            // TODO: define trackIDs
  )
  
  func main() {
    tracks, err := req.DecodeTracks(trackIDs...)
    if err != nil {
      // handle error
      return
    }
    // TODO: use tracks
  }
  ```

  </p>
</details>

### :notes: Interacting with an audio player

The interaction with an audio player **requires an active web socket connection**.

Additionally, a [voice update event **must be intercepted**](#briefcase-intercepting-a-voice-update-event) to play a track.

#### :x: Destroying an audio player

<details>
  <summary>Usage</summary>
  <p>

  ```go
  package main
  
  import "github.com/lukasl-dev/waterlink"
  
  var (
    conn    waterlink.Connection // TODO: open conn
    guildID uint                 // TODO: define guildID
  )
  
  func main() {
    if err := conn.Destroy(guildID); err != nil {
      // TODO: handle error
    }
  }
  ```

  </p>
</details>

#### :play_or_pause_button: Pausing/Resuming the current playing track

<details>
  <summary>Usage</summary>
  <p>

  ```go
  package main
  
  import "github.com/lukasl-dev/waterlink"
  
  var (
    conn    waterlink.Connection // TODO: open conn
    guildID uint                 // TODO: define guildID
    paused  bool                 // TODO: define paused
  )
  
  func main() {
    if err := conn.SetPaused(guildID, paused); err != nil {
      // TODO: handle error
    }
  }
  ```

  </p>
</details>

#### :headphones: Playing a track

<details>
  <summary>Usage without options</summary>
  <p>

  ```go
  package main
  
  import (
    "github.com/lukasl-dev/waterlink"
  )
  
  var (
    conn    waterlink.Connection // TODO: open conn
    guildID uint                 // TODO: define guildID
    trackID string               // TODO: load trackID
  )
  
  func main() {
    if err := conn.Play(guildID, trackID); err != nil {
      // TODO: handle error
    }
  }
  ```

  </p>
</details>

<details>
  <summary>Usage with options</summary>
  <p>

  ```go
  package main
  
  import (
    "github.com/lukasl-dev/waterlink"
    "github.com/lukasl-dev/waterlink/usecase/play"
  )
  
  var (
    conn    waterlink.Connection // TODO: open conn
    guildID uint                 // TODO: define guildID
    trackID string               // TODO: load trackID
    volume  uint                 // TODO: define volume
  )
  
  func main() {
    opts := play.NewOptions().WithVolume(volume) // more options available
    if err := conn.Play(guildID, trackID, opts); err != nil {
      // TODO: handle error
    }
  }
  ```

  </p>
</details>

#### :next_track_button: Seeking the current playing track

<details>
  <summary>Usage</summary>
  <p>

  ```go
  package main
  
  import (
    "github.com/lukasl-dev/waterlink"
  )
  
  var (
    conn     waterlink.Connection // TODO: open conn
    guildID  uint                 // TODO: define guildID
    position uint                 // TODO: define position
  )
  
  func main() {
    if err := conn.Seek(guildID, position); err != nil {
      // TODO: handle error
    }
  }

  ```

  </p>
</details>

#### :pause_button: Stopping the current playing track

<details>
  <summary>Usage</summary>
  <p>

  ```go
  package main
  
  import (
    "github.com/lukasl-dev/waterlink"
  )
  
  var (
    conn    waterlink.Connection // TODO: open conn
    guildID uint                 // TODO: define guildID
  )
  
  func main() {
    if err := conn.Stop(guildID); err != nil {
      // TODO: handle error
    }
  }
  ```

  </p>
</details>

#### :briefcase: Intercepting a voice update event

<details>
  <summary>Usage</summary>
  <p>

  ```go
  package main
  
  import (
    "github.com/lukasl-dev/waterlink"
  )
  
  var (
    conn      waterlink.Connection // TODO: open conn
    guildID   uint                 // TODO: define guildID
    sessionID string               // TODO: define sessionID
    token     string               // TODO: define token
    endpoint  string               // TODO: define endpoint
  )
  
  func main() {
    if err := conn.UpdateVoice(guildID, sessionID, token, endpoint); err != nil {
      // TODO: handle error
    }
  }
  ```

  </p>
</details>

#### :loud_sound: Updating the volume of an audio player

<details>
  <summary>Usage</summary>
  <p>

  ```go
  package main
  
  import (
    "github.com/lukasl-dev/waterlink"
  )
  
  var (
    conn    waterlink.Connection // TODO: open conn
    guildID uint                 // TODO: define guildID
    volume  uint                 // TODO: define volume
  )
  
  func main() {
    if err := conn.UpdateVolume(guildID, volume); err != nil {
      // TODO: handle error
    }
  }
  ```

  </p>
</details>
