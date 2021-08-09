package routinePool

import (
	"fmt"
	"testing"
)

func TestNewCachedRoutinePool(t *testing.T) {
	pool := NewCachedRoutinePool()
	//test nil process
	pool.Add(nil)
	//test adding one job
	f := Job(aa)
	pool.Add(f)

}
func aa(...interface{}) interface{} {
	fmt.Println("this is testcase!")
	return nil
}
