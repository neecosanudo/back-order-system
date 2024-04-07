package main

import "testing"

func TestOrders(t *testing.T) {
	ordersContainer := newOrderContainer()

	t.Run("create a new order ticket", func(t *testing.T) {
		ordersContainer.New()

		got := *ordersContainer.Get(1)
		want := Ticket{1, false, false}

		if len(ordersContainer.orders) != 1 {
			t.Errorf("order ticket is not added at ticket container. Got %v", ordersContainer.orders)
		}
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("send a not existing ID to get an order ticket", func(t *testing.T) {
		ordersContainer.New()

		got := ordersContainer.Get(47)

		if got != nil {
			t.Errorf("want nil, got %v", got)
		}

	})

	t.Run("update status order ticket", func(t *testing.T) {
		ordersContainer.New()

		ordersContainer.UpdateStatus(1)

		got := *ordersContainer.Get(1)
		want := Ticket{1, true, false}

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}

	})

	t.Run("cancel an order ticket", func(t *testing.T) {
		ordersContainer.Cancel(1)

		got := *ordersContainer.Get(1)
		want := Ticket{1, true, true}

		if got != want {
			t.Errorf("order ticket is not cancelled. Got %+v, want %+v", got, want)
		}
	})
}
