package lavalink

import "time"

// Stats contains statistics about a Lavalink node.
type Stats struct {
	// Players is the number of players connected to the node.
	Players uint `json:"players,omitempty"`

	// PlayingPlayers is the number of players playing a track.
	PlayingPlayers uint `json:"playingPlayers,omitempty"`

	// Uptime is the uptime of the node.
	Uptime time.Duration `json:"uptime,omitempty"`

	// Memory represents the memory stats of the node.
	Memory MemoryStats `json:"memory,omitempty"`

	// CPU represents the CPU stats of the node.
	CPU CPUStats `json:"cpu,omitempty"`

	// FrameStats represents the frame stats of the node.
	FrameStats *FrameStats `json:"frameStats,omitempty"`
}

type MemoryStats struct {
	// Free is the number of free memory in bytes.
	Free uint `json:"free,omitempty"`

	// Used is the number of used memory in bytes.
	Used uint `json:"used,omitempty"`

	// Allocated is the number of allocated memory in bytes.
	Allocated uint `json:"allocated,omitempty"`

	// Reservable is the number of reservable memory in bytes.
	Reservable uint `json:"reservable,omitempty"`
}

type CPUStats struct {
	// Cores is the number of CPU cores the node has.
	Cores uint `json:"cores,omitempty"`

	// SystemLoad is the system load of the node.
	SystemLoad float64 `json:"systemLoad,omitempty"`

	// LavalinkLoad is the Lavalink load of the node.
	LavalinkLoad float64 `json:"lavalinkLoad,omitempty"`
}

type FrameStats struct {
	// Sent is the number of frames sent to Discord.
	Sent uint `json:"sent,omitempty"`

	// Nulled is the number of frames that were nulled.
	Nulled uint `json:"nulled,omitempty"`

	// Deficit is the number of frames that were deficit.
	Deficit uint `json:"deficit,omitempty"`
}
