package message

import "github.com/lukasl-dev/waterlink/v2/filter"

type Filters struct {
	Outgoing
	Guild
	filter.Filters
}
