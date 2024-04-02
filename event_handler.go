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
