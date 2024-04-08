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
	for i, v := range s.containers {
		if v.id == id {
			return &s.containers[i]
		}
	}

	return nil
}

func (s *Stages) Remove(id uint) error {

	stage := s.Get(id)

	if stage == nil {
		return errors.New("stage does not exist")
	}

	removedStage := removeStage(id, s.containers)
	s.containers = append([]Container{}, removedStage...)

	newSequence := removeStageFromSequence(stage.id, s.sequence)
	s.sequence = append([]uint{}, newSequence...)

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

func removeStage(id uint, stages []Container) []Container {
	index := int(id) - 1

	switch index {
	case 0:
		stages = append([]Container{}, stages[1:]...)
		return stages
	case len(stages) - 1:
		stages = append([]Container{}, stages[:len(stages)-1]...)
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
