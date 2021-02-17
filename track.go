package waterlink

import "fmt"

type Track struct {
	ID   string `json:"track"`
	Info Info   `json:"info"`
}

func (t Track) String() string {
	return fmt.Sprintf("%s by %s", t.Info.Title, t.Info.Author)
}
