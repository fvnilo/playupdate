package worker

func newWorker(id int, inputs <-chan string, results chan<- int) {
	for j := range inputs {
		// call the api and await for response

		// send back the data to the results channel
	}
}
