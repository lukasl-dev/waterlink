package waterlink

import (
	"fmt"
	"github.com/lukasl-dev/waterlink/event"
)

func ExampleNewUnaryEventBus() {
	_ = NewUnaryEventBus(func(e interface{}) {
		switch e.(type) {
		case event.TrackEnd:
			fmt.Println("TrackEnd:", e)
		}
	})
}
