package main

import "testing"

func TestEventHandler(t *testing.T) {
	eventHandler := newEventHandler()

	t.Run("create new order", func(t *testing.T) {
		eventHandler.NewOrder()

		got := eventHandler.FindOrder(1)
		want := Order{1, ""}

		if *got != want {
			t.Errorf("order can't created, got %+v, want %+v", got, want)
		}
	})

	t.Run("find an order", func(t *testing.T) {
		numberOfOrders := 10
		createOrders(eventHandler, numberOfOrders)

		got := eventHandler.FindOrder(8)
		want := &eventHandler.orders[7]

		if got != want {
			t.Errorf("got %+v, want %+v", got, want)
		}
	})

	t.Run("update an order", func(t *testing.T) {
		numberOfOrders := 10
		createOrders(eventHandler, numberOfOrders)

		got := *eventHandler.UpdateOrder(3, "updated")
		want := Order{3, "updated"}

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func createOrders(eventHandler *EventHandler, numberOfOrders int) {
	for i := 0; i < numberOfOrders; i++ {
		eventHandler.NewOrder()
	}
}
