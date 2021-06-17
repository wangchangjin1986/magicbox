package structure

import (
	"fmt"
	"testing"
	"time"
)

func Test_newTimingWheel(t *testing.T) {
	type args struct {
		tickMs    int64
		wheelSize int64
		startMs   int64
		queue     *BlockingDelayQueue
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test1",
			args: args{
				tickMs:    1,
				wheelSize: 100,
				startMs:   time.Now().Unix() * 1000,
				queue:     NewBlockingDelayQueue(100),
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tw := newTimingWheel(tt.args.tickMs, tt.args.wheelSize, tt.args.startMs, tt.args.queue)
			tw.Start()
			tw.Stop()
		})
	}
}

func TestTimeWheel_AfterFunc(t *testing.T) {
	tw := newTimingWheel(1000, 2, time.Now().Unix()*1000, NewBlockingDelayQueue(1))
	tw.Start()
	fmt.Println(time.Now())
	tw.AfterFunc(time.Duration(5*time.Second), task)

	time.Sleep(10 * time.Second)
	tw.Stop()
}
func task() {
	fmt.Printf("aaaa\n")
	fmt.Println(time.Now())

}
