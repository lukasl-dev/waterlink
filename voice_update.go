package waterlink

type VoiceServerUpdate struct {
	GuildID  string `json:"guild_id"`
	Token    string `json:"token"`
	Endpoint string `json:"endpoint"`
}
