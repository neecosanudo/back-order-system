package event

import (
	"errors"
)

type stages struct {
	containers []container
	counter    uint
	sequence   []uint
}

type container struct {
	id     uint
	name   string
	orders []uint
}

// /
func (s *stages) new(name string) *container {

	container := container{
		s.counter,
		name,
		[]uint{},
	}
	s.counter++

	s.containers = append(s.containers, container)
	s.sequence = append(s.sequence, container.id)

	return &container
}

func (s *stages) get(id uint) *container {
	for i, v := range s.containers {
		if v.id == id {
			return &s.containers[i]
		}
	}

	return nil
}

func (s *stages) remove(id uint) error {

	stage := s.get(id)

	if stage == nil {
		return errors.New("stage does not exist")
	}

	removedStage := removeStage(id, s.containers)
	s.containers = append([]container{}, removedStage...)

	newSequence := removeStageFromSequence(stage.id, s.sequence)
	s.sequence = append([]uint{}, newSequence...)

	return nil
}

func (s *stages) updateSequence(newSequence []uint) error {
	if len(newSequence) != len(s.sequence) {
		return errors.New("sequence slice should have equal length")
	}

	s.sequence = newSequence

	return nil
}

func (s *stages) getSequence() []uint {
	return s.sequence
}

// /
func (c *container) add(id uint) {
	c.orders = append(c.orders, id)
}

func (c *container) remove(id uint) {
	updatedOrders := removeOrderFromContainer(id, c.orders)
	c.orders = append([]uint{}, updatedOrders...)
}

// /
func newStageContainer() *stages {
	stages := &stages{
		[]container{},
		0,
		[]uint{},
	}

	return stages
}

func removeStage(id uint, stages []container) []container {
	index := int(id)

	switch index {
	case 0:
		stages = append([]container{}, stages[1:]...)
		return stages
	case len(stages):
		stages = append([]container{}, stages[:len(stages)]...)
		return stages
	default:
		stages = append(stages[0:index], stages[index+1:]...)
		return stages
	}
}

func removeStageFromSequence(id uint, sequence []uint) []uint {
	var index int

	for i, v := range sequence {
		if v == id {
			index = i
		}
	}

	switch index {
	case 0:
		sequence = append([]uint{}, sequence[1:]...)
		return sequence
	case len(sequence) - 1:
		sequence = append([]uint{}, sequence[:len(sequence)-1]...)
		return sequence
	default:
		sequence = append(sequence[0:index], sequence[index+1:]...)
		return sequence
	}
}

func removeOrderFromContainer(id uint, container []uint) []uint {
	var index int

	for i, v := range container {
		if v == id {
			index = i
		}
	}

	switch index {
	case 0:
		container = append([]uint{}, container[1:]...)
		return container
	case len(container) - 1:
		container = append([]uint{}, container[:len(container)-1]...)
		return container
	default:
		container = append(container[0:index], container[index+1:]...)
		return container
	}
}
