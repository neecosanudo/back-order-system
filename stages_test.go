package event

import "testing"

func TestStages(t *testing.T) {
	t.Run("create a new stage", func(t *testing.T) {
		stages := newStageContainer()

		stages.new("new")

		got := stages.get(0)
		want := container{0, "new", []uint{}}

		if got.name != want.name {
			t.Errorf("got %+v, want %+v", got, want)
		}
	})

	t.Run("update sequence of stages", func(t *testing.T) {
		stages := newStageContainer()
		stages.new("first")
		stages.new("second")
		stages.new("third")

		newSequence := []uint{1, 3, 2}

		err := stages.updateSequence(newSequence)

		if err != nil {
			t.Errorf("sequence is not updated. Got error '%v', want nil", err)
		}
	})

	t.Run("remove a stage", func(t *testing.T) {
		stages := newStageContainer()
		stages.new("first")
		stages.new("second")
		stages.new("third")

		stages.remove(3)

		containersLength := len(stages.containers)
		sequenceLength := len(stages.sequence)

		got := stages.get(3)

		if containersLength != sequenceLength {
			t.Errorf("stage is not removed from both slices: containers & sequence")
		}
		if got != nil {
			t.Errorf("stages is not removed. Got %v, want nil", got)
		}
	})

	t.Run("remove an order ticket ID from a stage", func(t *testing.T) {
		stage := newStageContainer().new("new")
		stage.add(1)
		stage.add(2)
		stage.add(3)

		stage.remove(1)

		if len(stage.orders) != 2 {
			t.Errorf("order ticket ID is not removed from stage: %v", stage)
		}

	})
}
