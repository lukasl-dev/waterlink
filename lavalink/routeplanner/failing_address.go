package routeplanner

import "time"

type FailingAddress struct {
	// FailingAddress is the failing address.
	FailingAddress string `json:"failingAddress,omitempty"`

	// FailingTimestamp is the timestamp at which the address started failing.
	FailingTimestamp time.Time `json:"failingTimestamp,omitempty"`

	// FailingTime is the time at which the address started failing (pretty
	// printed).
	FailingTime string `json:"failingTime,omitempty"`
}
