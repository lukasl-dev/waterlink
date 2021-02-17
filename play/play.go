package play

type play struct {
	OP        string `json:"op"`
	GuildID   string `json:"guildId"`
	Track     string `json:"track"`
	StartTime string `json:"startTime,omitempty"`
	EndTime   string `json:"endTime,omitempty"`
	Volume    string `json:"volume,omitempty"`
	NoReplace bool   `json:"noReplace,omitempty"`
	Pause     bool   `json:"pause,omitempty"`
}

func New(guildID, track string) *play {
	return &play{
		OP:      "play",
		GuildID: guildID,
		Track:   track,
	}
}
