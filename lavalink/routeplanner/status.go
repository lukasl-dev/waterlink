package routeplanner

type Status struct {
	// Class is the name of the route planner implementation being used by the
	// Lavalink node.
	Class Type `json:"class,omitempty"`

	// Details contains status details of the route planner.
	Details Details `json:"details,omitempty"`
}
