package waterlink

type Info struct {
	Identifier string `json:"identifier"`
	Seekable   bool   `json:"isSeekable"`
	Author     string `json:"author"`
	Length     int    `json:"length"`
	Stream     bool   `json:"isStream"`
	Position   int    `json:"position"`
	Title      string `json:"title"`
	URI        string `json:"uri"`
}
