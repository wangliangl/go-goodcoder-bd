package engine


type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
}


type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}


func (e *ConcurrentEngine) Run(seeds ...Request) {

	in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.ConfigureMasterWorkerChan(in)

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

}

func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()

}
