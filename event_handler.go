package main

type EventHandler struct {
	orders  []Order
	counter uint
	stages  []Stage
}

type Stage struct {
	name      string
	container []Order
}

type Order struct {
	id    uint
	state string
}

func (e *EventHandler) NewStage(name string) {
	stage := Stage{
		name,
		[]Order{},
	}

	e.stages = append(e.stages, stage)
}

func (e *EventHandler) FindStage(name string) *Stage {
	var stage *Stage

	for i, v := range e.stages {
		if v.name == name {
			stage = &e.stages[i]
			return stage
		}
	}

	return nil
}

func (e *EventHandler) NewOrder() {
	defaultState := "default"
	e.counter++

	newOrder := Order{e.counter, defaultState}
	defaultStage := e.FindStage(defaultState)

	e.orders = append(e.orders, newOrder)
	defaultStage.container = append(defaultStage.container, newOrder)

}

func (e *EventHandler) FindOrder(id uint) (*Order, *Stage) {
	var order *Order

	for i, v := range e.orders {
		if v.id == id {
			order = &e.orders[i]
			return order, e.FindStage(order.state)
		}
	}

	return nil, nil
}

func (e *EventHandler) UpdateOrder(id uint, newState string) *Order {
	currentOrder, _ := e.FindOrder(id)

	currentOrder.state = newState

	return currentOrder
}

func newEventHandler() *EventHandler {
	var eventHandler EventHandler

	eventHandler.NewStage("default")

	return &eventHandler
}
