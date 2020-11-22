package workPool

type worker func(id int, jobs <-chan interface{}, results chan<- interface{})

type workerHandle func(workerId int, job interface{}) (result interface{})

type SimplePool struct {
	workers      []worker
	resultHandle func(p SimplePool)
	results      chan interface{}
	jobs         chan interface{}
}

func NewSimplePool(workers int, handle workerHandle) *SimplePool {
	ws := make([]worker, workers)
	results := make(chan interface{}, workers)
	jobs := make(chan interface{}, workers)

	for i := 0; i < len(ws); i++ {
		ws[i] = newWorker(handle)
	}

	return &SimplePool{
		workers: ws,
		results: results,
		jobs:    jobs,
	}
}

func (p SimplePool) StartWorkes(lock bool) {
	for i := 0; i < len(p.workers); i++ {
		go p.workers[i](i, p.jobs, p.results)
	}

	if p.resultHandle != nil {
		if lock {
			p.resultHandle(p)
		} else {
			go p.resultHandle(p)
		}
	} else if lock {
		<-p.results
	}

}

func (p *SimplePool) SetOnResult(handle func(p *SimplePool, result interface{})) {
	p.resultHandle = func(p SimplePool) {
		for r := range p.results {
			handle(&p, r)
		}
	}
}

func (p SimplePool) AddJob(job interface{}) {
	p.jobs <- job
}

func (p SimplePool) Close() {
	close(p.jobs)
	close(p.results)
}

func newWorker(handle workerHandle) worker {
	return func(id int, jobs <-chan interface{}, results chan<- interface{}) {
		for j := range jobs {
			results <- handle(id, j)
		}
	}
}
