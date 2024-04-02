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

func newEventHandler() *EventHandler {
	return &EventHandler{}
}
