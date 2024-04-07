package main

type Orders struct {
	counter uint
	orders  []Ticket
}

type Ticket struct {
	id        uint
	completed bool
	canceled  bool
}

// /
func (o *Orders) New() *Ticket {
	o.counter++

	defaultStatus := false
	defaultCanceledStatus := false

	ticket := Ticket{
		o.counter,
		defaultStatus,
		defaultCanceledStatus,
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

	if o.orders[index].canceled {
		return nil
	}

	o.orders[index].completed = !o.orders[index].completed

	return &o.orders[index]
}

func (o *Orders) Cancel(id uint) *Ticket {
	o.Get(id).canceled = true
	return o.Get(id)
}

// /
func newOrderContainer() *Orders {
	return &Orders{
		0,
		[]Ticket{},
	}
}
