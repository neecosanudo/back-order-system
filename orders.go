package event

type orders struct {
	counter uint
	orders  []ticket
}

type ticket struct {
	id        uint
	completed bool
	canceled  bool
}

// /
func (o *orders) new() *ticket {
	o.counter++

	defaultStatus := false
	defaultCanceledStatus := false

	ticket := ticket{
		o.counter,
		defaultStatus,
		defaultCanceledStatus,
	}

	o.orders = append(o.orders, ticket)

	return &ticket
}

func (o *orders) get(id uint) *ticket {
	index := int(id) - 1

	if len(o.orders) > index && index >= 0 {
		return &o.orders[index]
	}

	return nil
}

func (o *orders) updateStatus(id uint) *ticket {
	index := int(id) - 1

	if o.orders[index].canceled {
		return nil
	}

	o.orders[index].completed = !o.orders[index].completed

	return &o.orders[index]
}

func (o *orders) cancel(id uint) *ticket {
	if o.get(id).completed {
		return nil
	}

	o.get(id).canceled = true
	return o.get(id)
}

// /
func newOrderContainer() *orders {
	return &orders{
		0,
		[]ticket{},
	}
}
