package worker

import "sync"
import "github.com/nylo-andry/playupdate"

type Pool struct {
	workerCount   int
	updateService playupdate.UpdateService
}

func NewPool(workerCount int, updateService playupdate.UpdateService) *Pool {
	return &Pool{
		workerCount,
		updateService,
	}
}

func (p *Pool) Start(macAddresses []string) {
	var wg sync.WaitGroup
	inputs := make(chan string)

	for i := 0; i < p.workerCount; i++ {
		wg.Add(1)
		go startWorker(inputs, &wg, p.updateService)
	}

	for _, mac := range macAddresses {
		inputs <- mac
	}
	close(inputs)

	wg.Wait()

}
