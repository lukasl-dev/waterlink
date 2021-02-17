# waterlink

<div align="center">
  <a href="https://golang.org/">
    <img
      src="https://img.shields.io/badge/MADE%20WITH-GO-%23EF4041?style=for-the-badge"
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

<br>

- [waterlink](#waterlink)
  - [What is `waterlink`?](#what-is-waterlink)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Getting started](#getting-started)
    - [Create a Client](#create-a-client)
    - [Load one or more Tracks](#load-one-or-more-tracks)
    - [Play a Track](#play-a-track)
  - [Integration](#integration)
    - [Integrate `discordgo`](#integrate-discordgo)

---

## What is `waterlink`?

`Waterlink` is a [Lavalink](https://github.com/Frederikam/Lavalink) client written in Go. The client is based on [Lavalink](https://github.com/Frederikam/Lavalink) version **3.0** or higher.

---

## Prerequisites

To use a Go package such as `waterlink`, you must of course have Go installed on your system.

It is assumed that you have already worked with the Go environment. If this is not the case,
see [this page first](https://golang.org/doc/install).

---

## Installation

To use `waterlink` as a Go package, you must have it installed on your current system. If this is not the case, run the
command below.

```console
go get -u github.com/lukasl-dev/waterlink
```

---

## Getting started

### Create a [Client](https://pkg.go.dev/github.com/lukasl-dev/waterlink#Client)

The [Client](https://pkg.go.dev/github.com/lukasl-dev/waterlink#Client) is used to connect to the HTTP and Websocket server.

```go
client, err := waterlink.New(
  waterlink.HTTP(":8080"), // HTTP(S) server host address
  waterlink.WS(":8080"), // WS(S) server host address
  waterlink.Password("youshallnotpass"), // server authorisation password
  waterlink.UserID("botID"), // bot's user id
)

// handle error
```

### Load one or more [Tracks](https://pkg.go.dev/github.com/lukasl-dev/waterlink#Track)

See the [Track Loading API](https://github.com/Frederikam/Lavalink/blob/master/IMPLEMENTATION.md#track-loading-api).

```go
typ, playlist, tracks, err := client.LoadTracks("source")

// handle error
```

### Play a [Track](https://pkg.go.dev/github.com/lukasl-dev/waterlink#Track)

**To be able to play music, `waterlink` must be [integrated](#integration).**

```go
err := client.PlayTrack("guildID", track)

// handle error
```

---

## Integration

### Integrate [`discordgo`](https://github.com/bwmarrin/discordgo)

Here it is assumed that you already have experience with the [discordgo](https://github.com/bwmarrin/discordgo) package. If this is not the case, see [this page first](https://github.com/bwmarrin/discordgo#getting-started).

To be able to play audio, a `sessionID` is required.

```go
var sessionID string

discord.AddHandler(func(session *discordgo.Session, event *discordgo.Ready) {
  sessionID = event.SessionID
})
```

To be able to stream audio, the [VoiceServerUpdate](#https://github.com/bwmarrin/discordgo#VoiceServerUpdate) must be provided. See [this](https://github.com/Frederikam/Lavalink/blob/master/IMPLEMENTATION.md#provide-a-voice-server-update).

```go
discord.AddHandler(session *discordgo.Session, event *discordgo.VoiceServerUpdate) {
  err := client.VoiceUpdate(event.GuildID, sessionID, waterlink.VoiceServerUpdate{
    GuildID:  event.GuildID,
    Token:    event.Token,
    Endpoint: event.Endpoint,
  })

  // handle error
}
```

**After [`discordgo`](https://github.com/bwmarrin/discordgo) has been integrated, music can be [played](#play-a-track).**

Use the code below to join a guild's voice channel. After that music can be [played](#play-a-track).

```go
err := dg.ChannelVoiceJoinManual("guildID", "voiceChannelID", false, false)

// handle error
```
