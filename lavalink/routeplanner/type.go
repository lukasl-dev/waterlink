package routeplanner

type Type string

const (
	// TypeRotatingIPRoutePlanner is the type of the rotating IP route planner,
	// where the IP address used is switched on ban. Recommended for IPv4
	// blocks or IPv6 blocks smaller than a /64.
	TypeRotatingIPRoutePlanner Type = "IPRoutePlanner"

	// TypeNanoIPRoutePlanner is the type of the nano IP route planner, where
	// the IP address used is switched on clock update. Use with at least 1 /64
	// IPv6 block.
	TypeNanoIPRoutePlanner Type = "NanoIpRoutePlanner"

	// TypeRotatingNanoIPRoutePlanner is the type of the rotating nano IP route
	// planner, where the IP address used is switched on clock update, rotates
	// to a different /64 block on ban. Use with at least 2x /64 IPv6 blocks.
	TypeRotatingNanoIPRoutePlanner Type = "RotatingNanoIpRoutePlanner"

	// TypeBalancingIPRoutePlanner is the type of the balancing IP route
	// planner, where the IP address used is selected at random per request.
	// Recommended for larger IP blocks.
	TypeBalancingIPRoutePlanner Type = "BalancingIpRoutePlanner"
)
