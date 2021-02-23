package routinePool

import (
	"errors"
	"magicbox/msync"
)

//通用池子，新来一个作业就会直接起一个goruntine，超过池子的数量就会报错
type Pool struct {
	size int
	wg   *msync.MwaitGroup
}

func (p *Pool) start() {
	if p.wg != nil {
		go func() {
			p.wg.Wait()
		}()
	}
}
func (p *Pool) Size() int64 {
	return p.wg.Size()
}
func (p *Pool) Add(j Job) {
	w := Worker{
		job2run: j,
	}
	p.wg.Add(1)
	w.run(p)
}
func (p *Pool) Close() {}
func (p *Pool) Wait() {

	p.wg.Wait()
}

//固定大小的池子，池子启动时就会初始化对应size个数的goruntine，要执行的作业通过chan传递，如果当前chan的缓冲区满就会报错
type FixPool struct {
	size        int
	chanSize    int
	wg          *msync.MwaitGroup
	job2runChan chan Job
	commandChan chan int
}

func (f *FixPool) start() {
	for i := 0; i < f.size; i++ {
		w := FixWorker{
			job2runChan: &f.job2runChan,
		}
		f.wg.Add(1)
		w.run(f)
	}
	go func() {
		f.wait()
	}()
}
func (f *FixPool) Add(j Job) error {
	if f.chanSize == len(f.job2runChan) {
		return errors.New("pool is full, please wait moment")
	}
	f.job2runChan <- j
	return nil
}
func (f *FixPool) Close() {
	for i := 0; i < f.size; i++ {
		f.commandChan <- CLOSE
	}
	f.wg.Wait()
	close(f.commandChan)
	close(f.job2runChan)
}
func (f *FixPool) wait() {
	f.wg.Wait()
}
