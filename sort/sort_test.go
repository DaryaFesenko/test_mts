package sort

import (
	"testing"
)

func TestSort(t *testing.T) {
	testCases := []struct {
		Name string
		In1  []int
		In2  []int
		Out  []int
	}{
		{
			Name: "consistent slices",
			In1:  []int{1, 2, 3},
			In2:  []int{4, 5, 6},
			Out:  []int{1, 2, 3, 4, 5, 6},
		},
		{
			Name: "def",
			In1:  []int{1, 2, 3},
			In2:  []int{1, 2, 3},
			Out:  []int{1, 1, 2, 2, 3, 3},
		},
		{
			Name: "first empty",
			In1:  []int{},
			In2:  []int{1, 2, 3},
			Out:  []int{1, 2, 3},
		},
		{
			Name: "second empty",
			In1:  []int{1, 2, 3},
			In2:  []int{},
			Out:  []int{1, 2, 3},
		},
		{
			Name: "two empty",
			In1:  []int{},
			In2:  []int{},
			Out:  []int{},
		},
		{
			Name: "zero",
			In1:  []int{},
			In2:  []int{0},
			Out:  []int{0},
		},
		{
			Name: "first",
			In1:  []int{},
			In2:  []int{1},
			Out:  []int{1},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.Name, func(t *testing.T) {
			chIn1 := FillChannel(tt.In1)
			chIn2 := FillChannel(tt.In2)

			outCh := SequencesFromChannels(chIn1, chIn2)

			out := []int{}

			for val := range outCh {
				out = append(out, val)
			}

			if !compareOutSlice(out, tt.Out) {
				t.Fatalf("got %v, but want %v", out, tt.Out)
			}
		})
	}
}

//сравнивает два слайса
func compareOutSlice(got []int, trueOut []int) bool {
	if len(got) != len(trueOut) {
		return false
	}

	for i := 0; i < len(trueOut); i++ {
		if got[i] != trueOut[i] {
			return false
		}
	}

	return true
}

func BenchmarkSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n1 := []int{1, 4, 6, 8, 9, 12, 23}
		n2 := []int{3, 5, 8, 9, 10, 15, 18, 24}
		chIn1 := FillChannel(n1)
		chIn2 := FillChannel(n2)

		outCh := SequencesFromChannels(chIn1, chIn2)

		out := []int{}

		for val := range outCh {
			_ = append(out, val)
		}
	}
}
