package event

type event struct {
	stages stages
	orders orders
}

// /
func NewEvent() *event {
	return &event{
		*newStageContainer(),
		*newOrderContainer(),
	}
}
