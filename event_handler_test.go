package main

import "testing"

func TestEventHandler(t *testing.T) {
	eventHandler := newEventHandler()

	t.Run("create new order", func(t *testing.T) {
		eventHandler.NewOrder()

		got := eventHandler.orders[0]
		want := Order{1}

		if got != want {
			t.Errorf("order can't created, got %+v, want %+v", got, want)
		}
	})
}
