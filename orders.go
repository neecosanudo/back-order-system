package main

type Orders struct {
	counter uint
	orders  []Ticket
}

type Ticket struct {
	id        uint
	completed bool
}

// /
func (o *Orders) New() *Ticket {
	o.counter++

	defaultStatus := false

	ticket := Ticket{
		o.counter,
		defaultStatus,
	}

	o.orders = append(o.orders, ticket)

	return &ticket
}

func (o *Orders) Get(id uint) *Ticket {
	index := int(id) - 1

	if len(o.orders) > index && index >= 0 {
		return &o.orders[index]
	}

	return nil
}

func (o *Orders) UpdateStatus(id uint) *Ticket {
	index := int(id) - 1
	o.orders[index].completed = !o.orders[index].completed

	return &o.orders[index]
}

// /
func newOrderContainer() *Orders {
	return &Orders{
		0,
		[]Ticket{},
	}
}
