package main

import (
	"fmt"

	"github.com/DaryaFesenko/test_mts/sort"
)

func main() {
	chIn1 := sort.FillChannel([]int{1, 2, 3})
	chIn2 := sort.FillChannel([]int{1, 2, 3})

	outCh := sort.SequencesFromChannels(chIn1, chIn2)

	for val := range outCh {
		fmt.Println(val)
	}
}
