package main

import "testing"

func TestStages(t *testing.T) {
	t.Run("create a new stage", func(t *testing.T) {
		stages := newStageContainer()

		stages.New("new")

		got := stages.Get(1)
		want := Container{1, "new", []uint{}}

		if got.name != want.name {
			t.Errorf("got %+v, want %+v", got, want)
		}
	})

	t.Run("update sequence of stages", func(t *testing.T) {
		stages := newStageContainer()
		stages.New("first")
		stages.New("second")
		stages.New("third")

		newSequence := []uint{1, 3, 2}

		err := stages.UpdateSequence(newSequence)

		if err != nil {
			t.Errorf("sequence is not updated. Got error '%v', want nil", err)
		}
	})

	t.Run("remove a stage", func(t *testing.T) {
		stages := newStageContainer()
		stages.New("first")
		stages.New("second")
		stages.New("third")

		stages.Remove(1)

		containersLength := len(stages.containers)
		sequenceLength := len(stages.sequence)

		if containersLength != sequenceLength {
			t.Errorf("stage is not removed")
		}
	})
}
