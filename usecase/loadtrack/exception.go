package loadtrack

type Exception struct {
	Message  string `json:"message,omitempty"`
	Severity string `json:"severity,omitempty"`
}
