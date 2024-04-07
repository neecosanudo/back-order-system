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
}
