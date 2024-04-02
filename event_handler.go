package main

type EventHandler struct {
	orders  []Order
	counter uint
}

type Order struct {
	id    uint
	state string
}

func (e *EventHandler) NewOrder() {
	defaultState := ""
	e.counter++

	newOrder := Order{e.counter, defaultState}

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

func (e *EventHandler) UpdateOrder(id uint, newState string) *Order {
	currentOrder := e.FindOrder(id)

	currentOrder.state = newState

	return currentOrder
}

func newEventHandler() *EventHandler {
	return &EventHandler{}
}
