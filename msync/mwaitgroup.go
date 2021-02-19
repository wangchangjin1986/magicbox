package msync

import (
	"sync"
	"sync/atomic"
)

type MwaitGroup struct {
	wg    *sync.WaitGroup
	count int64
	mux   sync.Mutex
}

func (m *MwaitGroup) Add(d int) {
	m.mux.Lock()
	atomic.AddInt64(&m.count, int64(d))
	m.wg.Add(d)
	m.mux.Unlock()
}
func (m *MwaitGroup) Done() {
	m.mux.Lock()
	atomic.AddInt64(&m.count, int64(-1))
	m.wg.Done()
	m.mux.Unlock()
}
func (m *MwaitGroup) Wait() {
	m.wg.Wait()
}
func (m *MwaitGroup) Size() int64 {
	return m.count
}

func New() *MwaitGroup {
	return &MwaitGroup{
		wg:    &sync.WaitGroup{},
		count: 0,
		mux:   sync.Mutex{},
	}
}
