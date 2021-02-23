package structure

import (
	"container/list"
	"sync"
	"sync/atomic"
	"unsafe"
)

//just used for timeWheel

// Timer represents a single event. When the Timer expires, the given
// task will be executed.
type timerTimeWheel struct {
	expiration int64 // in milliseconds
	task       func()

	// The bucket that holds the list to which this timer's element belongs.
	//
	// NOTE: This field may be updated and read concurrently,
	// through Timer.Stop() and Bucket.Flush().
	b unsafe.Pointer // type: *bucket

	// The timer's element.
	element *list.Element
}

func (t *timerTimeWheel) getBucket() *bucket {
	return (*bucket)(atomic.LoadPointer(&t.b))
}

func (t *timerTimeWheel) setBucket(b *bucket) {
	atomic.StorePointer(&t.b, unsafe.Pointer(b))
}

// Stop prevents the Timer from firing. It returns true if the call
// stops the timer, false if the timer has already expired or been stopped.
//
// If the timer t has already expired and the t.task has been started in its own
// goroutine; Stop does not wait for t.task to complete before returning. If the caller
// needs to know whether t.task is completed, it must coordinate with t.task explicitly.
func (t *timerTimeWheel) Stop() bool {
	stopped := false
	for b := t.getBucket(); b != nil; b = t.getBucket() {
		// If b.Remove is called just after the timing wheel's goroutine has:
		//     1. removed t from b (through b.Flush -> b.remove)
		//     2. moved t from b to another bucket ab (through b.Flush -> b.remove and ab.Add)
		// this may fail to remove t due to the change of t's bucket.
		stopped = b.Remove(t)

		// Thus, here we re-get t's possibly new bucket (nil for case 1, or ab (non-nil) for case 2),
		// and retry until the bucket becomes nil, which indicates that t has finally been removed.
	}
	return stopped
}

type bucketTimeWheel struct {
	// 64-bit atomic operations require 64-bit alignment, but 32-bit
	// compilers do not ensure it. So we must keep the 64-bit field
	// as the first field of the struct.
	//
	// For more explanations, see https://golang.org/pkg/sync/atomic/#pkg-note-BUG
	// and https://go101.org/article/memory-layout.html.
	// 任务的过期时间
	expiration int64

	mu sync.Mutex
	// 相同过期时间的任务队列
	timers *list.List
}

func newTimeWheelBucket() *bucket {
	return &bucketTimeWheel{
		timers:     list.New(),
		expiration: -1,
	}
}

func (b *bucketTimeWheel) Expiration() int64 {
	return atomic.LoadInt64(&b.expiration)
}

func (b *bucketTimeWheel) SetExpiration(expiration int64) bool {
	return atomic.SwapInt64(&b.expiration, expiration) != expiration
}

func (b *bucketTimeWheel) Add(t *timerTimeWheel) {
	b.mu.Lock()

	e := b.timers.PushBack(t)
	t.setBucket(b)
	t.element = e

	b.mu.Unlock()
}

func (b *bucketTimeWheel) remove(t *Timer) bool {
	if t.getBucket() != b {
		// If remove is called from t.Stop, and this happens just after the timing wheel's goroutine has:
		//     1. removed t from b (through b.Flush -> b.remove)
		//     2. moved t from b to another bucket ab (through b.Flush -> b.remove and ab.Add)
		// then t.getBucket will return nil for case 1, or ab (non-nil) for case 2.
		// In either case, the returned value does not equal to b.
		return false
	}
	b.timerTimeWheel.Remove(t.element)
	t.setBucket(nil)
	t.element = nil
	return true
}

func (b *bucketTimeWheel) Remove(t *timerTimeWheel) bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.remove(t)
}

func (b *bucketTimeWheel) Flush(reinsert func(*timerTimeWheel)) {
	var ts []*Timer

	b.mu.Lock()
	// 循环获取bucket队列节点
	for e := b.timerTimeWheel.Front(); e != nil; {
		next := e.Next()

		t := e.Value.(*timerTimeWheel)
		// 将头节点移除bucket队列
		b.remove(t)
		ts = append(ts, t)

		e = next
	}
	b.mu.Unlock()

	b.SetExpiration(-1) // TODO: Improve the coordination with b.Add()

	for _, t := range ts {
		reinsert(t)
	}
}
