package event

type Stats struct {
	Players        uint        `json:"players,omitempty"`
	PlayingPlayers uint        `json:"playingPlayers,omitempty"`
	Memory         MemoryStats `json:"memory,omitempty"`
	CPU            CPUStats    `json:"cpu,omitempty"`
	Frame          FrameStats  `json:"frameStats,omitempty"`
}

type MemoryStats struct {
	Free       uint `json:"free,omitempty"`
	Used       uint `json:"used,omitempty"`
	Allocated  uint `json:"allocated,omitempty"`
	Reservable uint `json:"reservable,omitempty"`
}

type CPUStats struct {
	Cores        uint    `json:"cores,omitempty"`
	SystemLoad   float64 `json:"systemLoad,omitempty"`
	LavalinkLoad float64 `json:"lavalinkLoad,omitempty"`
}

type FrameStats struct {
	Sent    uint `json:"sent,omitempty"`
	Nulled  uint `json:"nulled,omitempty"`
	Deficit int  `json:"deficit,omitempty"`
}
