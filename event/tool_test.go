package event_test

import (
	"fmt"
	"github.com/go-chassis/go-archaius/event"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPopulateEvents(t *testing.T) {
	events, err := event.PopulateEvents(
		"test",
		map[string]interface{}{
			"k1": "v1",
			"k3": "v2",
			"k4": "v4",
		},
		map[string]interface{}{
			"k1": "v1",
			"k2": "v2",
			"k3": "v3",
		})
	assert.NoError(t, err)
	for _, eve := range events {
		fmt.Printf("actual: %+v\n", eve)
	}
	expected := []*event.Event{
		{
			EventSource: "test",
			EventType:   event.Create,
			Key:         "k2",
			Value:       "v2",
		},
		{
			EventSource: "test",
			EventType:   event.Update,
			Key:         "k3",
			Value:       "v3",
		},
		{
			EventSource: "test",
			EventType:   event.Delete,
			Key:         "k4",
			Value:       "v4",
		},
	}

	for _, eve := range expected {
		fmt.Printf("expected: %+v\n", eve)
	}

	assert.Equal(t, len(expected), len(events))

	for i := 0; i < len(expected); i++ {
		assert.Equal(t, expected[i].EventSource, events[i].EventSource)
		assert.Equal(t, expected[i].Key, events[i].Key)
	}

}
