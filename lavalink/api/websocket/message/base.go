package message

// Base is the base of a websocket message. It is only used to distinguish
// between different messages
type Base struct {
	Op Op `json:"op,omitempty"`
}
