package routeplanner

// IPBlockType is used to distinguish between IPv4 and IPv6 IP blocks.
type IPBlockType string

const (
	// IPBlockTypeIPv4 represents an IPv4 IP block.
	IPBlockTypeIPv4 IPBlockType = "Inet4Address"

	// IPBlockTypeIPv6 represents an IPv6 IP block.
	IPBlockTypeIPv6 IPBlockType = "Inet6Address"
)

type IPBlock struct {
	// Type is the type of the IP block.
	Type IPBlockType `json:"type,omitempty"`

	// Size is the size of the IP block.
	Size string `json:"size,omitempty"`
}
