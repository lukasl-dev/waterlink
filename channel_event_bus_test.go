package waterlink

import "fmt"

func ExampleNewChannelEventBus() {
	events := NewChannelEventBus()

	go func() {
		for {
			select {
			case e := <-events.TrackEnds():
				fmt.Println("TrackEnd:", e)
			}
		}
	}()
}
