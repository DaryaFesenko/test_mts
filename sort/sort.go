package sort

// Function for sorting two sorted sequences into one
func SequencesFromChannels(in1 <-chan int, in2 <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		val1, ok1 := <-in1
		val2, ok2 := <-in2

		for {
			if !ok1 && !ok2 {
				close(out)
				return
			}

			if !ok1 && ok2 {
				out <- val2
				readAllValues(in2, out)

				close(out)

				return
			}

			if ok1 && !ok2 {
				out <- val1
				readAllValues(in1, out)

				close(out)

				return
			}

			if val1 < val2 {
				out <- val1
				val1, ok1 = <-in1
			} else {
				out <- val2
				val2, ok2 = <-in2
			}
		}
	}()

	return out
}

// read all values ​​from the channel
func readAllValues(in <-chan int, out chan int) {
	for val := range in {
		out <- val
	}
}

//Method for filling incoming channels
func FillChannel(n []int) <-chan int {
	chIn := make(chan int)

	go func() {
		for _, val := range n {
			chIn <- val
		}

		close(chIn)
	}()

	return chIn
}
