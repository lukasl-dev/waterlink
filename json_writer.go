package waterlink

// jsonWriter wraps its WriteJSON() method.
type jsonWriter interface {
	// WriteJSON writes v as JSON.
	WriteJSON(v interface{}) error
}
