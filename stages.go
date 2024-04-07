package main

type Stages struct {
	containers []Container
	counter    uint
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

	return &container
}

func (s *Stages) Get(id uint) *Container {
	index := int(id) - 1

	if len(s.containers) > index && index >= 0 {
		return &s.containers[index]
	}

	return nil
}

// /
func newStageContainer() *Stages {
	return &Stages{
		[]Container{},
		0,
	}
}
