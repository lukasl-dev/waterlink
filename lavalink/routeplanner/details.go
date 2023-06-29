package routeplanner

type Details struct {
	// IPBlock is the IPBlock being used.
	IPBlock IPBlock `json:"ipBlock,omitempty"`

	// FailingAddresses is an array of failing addresses.
	FailingAddresses []FailingAddress `json:"failingAddresses,omitempty"`

	// RotateIndex is the number of rotations. It is only present in a route
	// planner of type TypeRotatingIPRoutePlanner.
	RotateIndex string `json:"rotateIndex,omitempty"`

	// IPIndex is the current offset in the block. It is only present in a route
	// planner of type TypeRotatingIPRoutePlanner.
	IPIndex string `json:"ipIndex,omitempty"`

	// CurrentAddress is the current address being used. It is only present in a
	// route planner of type TypeRotatingIPRoutePlanner.
	CurrentAddress string `json:"currentAddress,omitempty"`

	// CurrentAddressIndex is the current offset in the IP block. It is only in
	// a route planner of type TypeNanoIPRoutePlanner or TypeRotatingNanoIPRoutePlanner.
	CurrentAddressIndex string `json:"currentAddressIndex,omitempty"`

	// BlockIndex holds information which /64-block IPs are chosen. This number
	// increases on each ban. It is only present in a route planner of type
	// TypeRotatingNanoIPRoutePlanner.
	BlockIndex string `json:"blockIndex,omitempty"`
}
