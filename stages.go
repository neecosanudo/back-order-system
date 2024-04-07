package main

import (
	"errors"
)

type Stages struct {
	containers []Container
	counter    uint
	sequence   []uint
}

type Container struct {
	id     uint
	name   string
	orders []uint
}

// /
func (s *Stages) New(name string) *Container {
	s.counter++

	container := Container{
		s.counter,
		name,
		[]uint{},
	}
	s.containers = append(s.containers, container)
	s.sequence = append(s.sequence, container.id)

	return &container
}

func (s *Stages) Get(id uint) *Container {
	index := int(id) - 1

	if len(s.containers) > index && index >= 0 {
		return &s.containers[index]
	}

	return nil
}

func (s *Stages) UpdateSequence(newSequence []uint) error {
	if len(newSequence) != len(s.sequence) {
		return errors.New("sequence slice should have equal length")
	}

	s.sequence = newSequence

	return nil
}

func (s *Stages) GetSequence() []uint {
	return s.sequence
}

// /
func newStageContainer() *Stages {
	return &Stages{
		[]Container{},
		0,
		[]uint{},
	}
}
