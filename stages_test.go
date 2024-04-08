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

		got := stages.Get(1)

		if containersLength != sequenceLength {
			t.Errorf("stage is not removed from both slices: containers & sequence")
		}
		if got != nil {
			t.Errorf("stages is not removed. Got %v, want nil", got)
		}
	})

	t.Run("remove an order ticket ID from a stage", func(t *testing.T) {
		stage := newStageContainer().New("new")
		stage.Add(1)
		stage.Add(2)
		stage.Add(3)

		stage.Remove(1)

		if len(stage.orders) != 2 {
			t.Errorf("order ticket ID is not removed from stage: %v", stage)
		}

	})
}
