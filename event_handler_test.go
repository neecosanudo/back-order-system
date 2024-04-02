package main

import "testing"

func TestEventHandler(t *testing.T) {
	eventHandler := newEventHandler()

	t.Run("create new order", func(t *testing.T) {
		eventHandler.NewOrder()

		order, stage := eventHandler.FindOrder(1)
		expectedOrder := Order{1, "default"}
		expectedStage := eventHandler.FindStage(order.state)

		if *order != expectedOrder {
			t.Errorf("did not create order, got %+v, want %+v", order, expectedOrder)
		}

		if stage != expectedStage {
			t.Errorf("did not add at default stage, got %v, want %v", stage, expectedStage)
		}
	})

	t.Run("find an order", func(t *testing.T) {
		numberOfOrders := 10
		createOrders(eventHandler, numberOfOrders)

		got, _ := eventHandler.FindOrder(8)
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

	t.Run("create a new stage", func(t *testing.T) {
		stageName := "new stage"
		eventHandler.NewStage(stageName)

		got := eventHandler.FindStage(stageName)
		want := &Stage{"new stage", []Order{}}

		if got.name != want.name {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func createOrders(eventHandler *EventHandler, numberOfOrders int) {
	for i := 0; i < numberOfOrders; i++ {
		eventHandler.NewOrder()
	}
}
