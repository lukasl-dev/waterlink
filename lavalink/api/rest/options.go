package rest

// Options contains a set of configurable values that are sent to the Lavalink
// REST API in every request.
type Options struct {
	// Authorization is the password set in the application.yml file of the
	// Lavalink node.
	Authorization string `json:"authorization,omitempty"`
}
