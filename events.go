package event

type event struct {
	stages stages
	orders orders
}

// /
func NewEvent() *event {
	event := &event{
		*newStageContainer(),
		*newOrderContainer(),
	}

	event.stages.new("tracked")

	return event
}
