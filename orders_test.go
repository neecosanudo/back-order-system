package event

import "testing"

func TestOrders(t *testing.T) {

	t.Run("create a new order ticket", func(t *testing.T) {
		ordersContainer := newOrderContainer()
		ordersContainer.new()

		got := *ordersContainer.get(1)
		want := ticket{1, status{false, 0}, false}

		if len(ordersContainer.orders) != 1 {
			t.Errorf("order ticket is not added at ticket container. Got %v", ordersContainer.orders)
		}
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("send a not existing ID to get an order ticket", func(t *testing.T) {
		ordersContainer := newOrderContainer()
		ordersContainer.new()

		got := ordersContainer.get(47)

		if got != nil {
			t.Errorf("want nil, got %v", got)
		}

	})

	t.Run("update status order ticket", func(t *testing.T) {
		ordersContainer := newOrderContainer()
		ordersContainer.new()

		ordersContainer.updateStatus(1)

		got := *ordersContainer.get(1)
		want := ticket{1, status{true, 0}, false}

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}

	})

	t.Run("cancel an order ticket", func(t *testing.T) {
		ordersContainer := newOrderContainer()
		ordersContainer.new()

		ordersContainer.cancel(1)

		got := *ordersContainer.get(1)
		want := ticket{1, status{false, 0}, true}

		if got != want {
			t.Errorf("order ticket is not cancelled. Got %+v, want %+v", got, want)
		}
	})

	t.Run("a canceled order ticket shouldn't update its status", func(t *testing.T) {
		ordersContainer := newOrderContainer()
		ordersContainer.new()
		ordersContainer.cancel(1)

		got := ordersContainer.updateStatus(1)

		if got != nil {
			t.Errorf("canceled order ticket was updated. got %+v", got)
		}
	})

	t.Run("should not be cancel an order ticket if it is status completed", func(t *testing.T) {
		ordersContainer := newOrderContainer()
		ordersContainer.new()
		ordersContainer.updateStatus(1)

		got := ordersContainer.cancel(1)

		if got != nil {
			t.Errorf("completed order ticket was canceled. got %+v", got)
		}

	})
}
