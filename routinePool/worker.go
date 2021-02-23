package routinePool

import (
	"time"
)

type Job func(...interface{}) interface{}

//执行任务的goruntine
type Worker struct {
	job2run Job
}

func (w *Worker) run(pool *Pool) {
	go func() {
		defer pool.wg.Done()
		w.job2run()
	}()
}

type FixWorker struct {
	job2runChan *chan Job
}

func (w *FixWorker) run(pool *FixPool) {
	go func() {
		defer pool.wg.Done()
		for {
			if len(*w.job2runChan) == 0 {
				time.Sleep(1000)
				continue
			}
			select {
			case job2run := <-pool.job2runChan:
				job2run()
			case command := <-pool.commandChan:
				switch command {
				case CLOSE:
					goto exit
				}
			}
		}
	exit:
		return
	}()

}
