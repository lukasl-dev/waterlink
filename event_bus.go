package waterlink

import (
	"encoding/json"
	"fmt"
	"github.com/lukasl-dev/waterlink/event"
	"github.com/lukasl-dev/waterlink/internal/message"
	"github.com/lukasl-dev/waterlink/internal/message/opcode"
	"reflect"
)

// eventBus is responsible for handling incoming message payloads and dispatching
// them to the defined handler.
type eventBus struct {
	// handler is the handler to dispatch incoming messages to.
	handler EventHandler
}

// newEventBus returns a new eventBus that dispatches events to the given
// handler.
func newEventBus(handler EventHandler) *eventBus {
	return &eventBus{handler: handler}
}

// receive handles payloads of incoming messages from the server and handles it
// using the bus' handler.
func (bus *eventBus) receive(data []byte) error {
	msg, err := bus.unmarshalMessage(data)
	if err != nil {
		return err
	}
	bus.handler.HandleEvent(reflect.ValueOf(msg).Elem().Interface())
	return nil
}

// unmarshalMessage unmarshals the given data into an event specialization and
// returns it.
func (bus *eventBus) unmarshalMessage(data []byte) (interface{}, error) {
	var msg message.Incoming
	if err := json.Unmarshal(data, &msg); err != nil {
		return nil, bus.wrapErr("unmarshal incoming message", err)
	}

	switch msg.Op {
	case opcode.PlayerUpdate:
		return bus.unmarshalPlayerUpdate(data)
	case opcode.Stats:
		return bus.unmarshalStats(data)
	case opcode.Event:
		return bus.unmarshalEvent(data)
	default:
		return nil, bus.newErr(fmt.Sprintf("unsupported opcode %q received", msg.Op))
	}
}

// unmarshalPlayerUpdate unmarshals the given data into an event.PlayerUpdate
// and returns it.
func (bus *eventBus) unmarshalPlayerUpdate(data []byte) (*event.PlayerUpdate, error) {
	var evt event.PlayerUpdate
	if err := json.Unmarshal(data, &evt); err != nil {
		return nil, bus.wrapErr("unmarshal player update", err)
	}
	return &evt, nil
}

// unmarshalStats unmarshals the given data into an event.Stats and returns it.
func (bus *eventBus) unmarshalStats(data []byte) (*event.Stats, error) {
	var evt event.Stats
	if err := json.Unmarshal(data, &evt); err != nil {
		return nil, bus.wrapErr("unmarshal stats", err)
	}
	return &evt, nil
}

// unmarshalEvent unmarshals the given data into an event specialization and
// returns it.
func (bus *eventBus) unmarshalEvent(data []byte) (interface{}, error) {
	var evt message.Event
	if err := json.Unmarshal(data, &evt); err != nil {
		return nil, bus.wrapErr("unmarshal event", err)
	}

	switch evt.Type {
	case message.EventTypeTrackStart:
		return bus.unmarshalTrackStart(data)
	case message.EventTypeTrackEnd:
		return bus.unmarshalTrackEnd(data)
	case message.EventTypeTrackException:
		return bus.unmarshalTrackException(data)
	case message.EventTypeTrackStuck:
		return bus.unmarshalTrackStuck(data)
	case message.EventTypeWebSocketClosed:
		return bus.unmarshalWebSocketClosed(data)
	default:
		return nil, bus.newErr(fmt.Sprintf("unsupported event type %q received", evt.Type))
	}
}

// unmarshalTrackStart unmarshals the given data into an event.TrackStart and
// returns it.
func (bus *eventBus) unmarshalTrackStart(data []byte) (e *event.TrackStart, err error) {
	var evt event.TrackStart
	if err := json.Unmarshal(data, &evt); err != nil {
		return nil, bus.wrapErr("unmarshal track start", err)
	}
	return &evt, nil
}

// unmarshalTrackEnd unmarshals the given data into an event.TrackEnd and
// returns it.
func (bus *eventBus) unmarshalTrackEnd(data []byte) (e *event.TrackEnd, err error) {
	var evt event.TrackEnd
	if err := json.Unmarshal(data, &evt); err != nil {
		return nil, bus.wrapErr("unmarshal track end", err)
	}
	return &evt, nil
}

// unmarshalTrackException unmarshals the given data into an event.TrackException
// and returns it.
func (bus *eventBus) unmarshalTrackException(data []byte) (e *event.TrackException, err error) {
	var evt event.TrackException
	if err := json.Unmarshal(data, &evt); err != nil {
		return nil, bus.wrapErr("unmarshal track exception", err)
	}
	return &evt, nil
}

// unmarshalTrackStuck unmarshals the given data into an event.TrackStuck and
// returns it.
func (bus *eventBus) unmarshalTrackStuck(data []byte) (e *event.TrackStuck, err error) {
	var evt event.TrackStuck
	if err := json.Unmarshal(data, &evt); err != nil {
		return nil, bus.wrapErr("unmarshal track stuck", err)
	}
	return &evt, nil
}

// unmarshalWebSocketClosed unmarshals the given data into an event.WebSocketClosed
// and returns it.
func (bus *eventBus) unmarshalWebSocketClosed(data []byte) (e *event.WebSocketClosed, err error) {
	var evt event.WebSocketClosed
	if err := json.Unmarshal(data, &evt); err != nil {
		return nil, bus.wrapErr("unmarshal web socket closed", err)
	}
	return &evt, nil
}

// newError returns a new error with 'event bus' as prefix and the given msg.
func (bus *eventBus) newErr(msg string) error {
	return fmt.Errorf("event bus: %s", msg)
}

// wrapErr wraps an error with an 'event bus' prefix to make it easier to
// identify the source of the error.
func (bus *eventBus) wrapErr(action string, err error) error {
	return fmt.Errorf("event bus: %s: %w", action, err)
}
