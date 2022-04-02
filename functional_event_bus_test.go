package waterlink

import (
	"fmt"
	"github.com/lukasl-dev/waterlink/event"
)

func ExampleNewFunctionalEventBus() {
	events := NewFunctionalEventBus()

	events.HandleTrackEnd(func(e event.TrackEnd) {
		fmt.Println("TrackEnd:", e)
	})
}
