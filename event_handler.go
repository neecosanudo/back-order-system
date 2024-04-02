package main

type EventHandler struct {
	orders  []Order
	counter uint
}

type Order struct {
	id uint
}

func (e *EventHandler) NewOrder() {
	e.counter++

	newOrder := Order{e.counter}

	e.orders = append(e.orders, newOrder)
}

func (e *EventHandler) FindOrder(id uint) *Order {
	var order *Order

	for i, v := range e.orders {
		if v.id == id {
			order = &e.orders[i]
			return order
		}
	}

	return nil
}

func newEventHandler() *EventHandler {
	return &EventHandler{}
}
