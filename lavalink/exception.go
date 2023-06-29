package lavalink

// Exception is the error (exception) implementation of the exception
// behind TrackException.
type Exception struct {
	Message  string            `json:"message,omitempty"`
	Severity ExceptionSeverity `json:"severity,omitempty"`
	Cause    string            `json:"cause,omitempty"`
}

var _ error = (*Exception)(nil)

// Error returns the error Message of tee.
func (ex *Exception) Error() string {
	return ex.Message
}

// ExceptionSeverity represents the severity of a Exception.
type ExceptionSeverity string

const (
	// ExceptionSeverityCommon indicates that the cause is known and
	// expected, meaning that there is nothing wrong with Lavalink itself.
	ExceptionSeverityCommon ExceptionSeverity = "COMMON"

	// ExceptionSeveritySuspicious indicates that the cause might not
	// be exactly known, but it is possibly caused by outside factors.
	//
	// Example:
	// An outside service responds in a format that Lavalink does not expect.
	ExceptionSeveritySuspicious ExceptionSeverity = "SUSPICIOUS"

	// ExceptionSeverityFault indicates that the probable cause is an
	// issue with Lavalink or that there is no way to tell what the cause might
	// be. This is the default level and other levels are used in cases where
	// the thrower has more in-depth knowledge about the error.
	ExceptionSeverityFault ExceptionSeverity = "FAULT"
)
