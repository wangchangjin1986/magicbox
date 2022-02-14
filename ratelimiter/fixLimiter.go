package ratelimiter

import (
	"sync"
	"sync/atomic"
	"time"
)

func countCompute(f *FixMetricLimiter) int64 {
	if f.Distribution == false { //如果不是分布式，则用本地统计数值
		//如果在当前窗口
		if f.LocalWindow.IsInCurrentWindow(time.Now()) {
			return atomic.LoadInt64(&f.LocalWindow.CurrentValue)
		} else {
			f.mux.Lock()
			defer f.mux.Unlock()
			atomic.StoreInt64(&f.LocalWindow.CurrentValue, 0)
			f.LocalWindow.CurrentWindowTime()
			return f.LocalWindow.CurrentValue
		}
	} else { //否则为分布式，统一限流值
		return 0
	}
}

/*
*固定限流值的限流器
 */
type FixMetricLimiter struct {
	Interval       int64                           //统计值窗口，单位s
	Max            int64                           //阈值
	MetrcisCompute func(f *FixMetricLimiter) int64 //统计值计算的函数
	Distribution   bool                            //是否为分布式全局统一的限流器
	LocalWindow    *LocalMetrciWindow              //本地的窗口metric统计值
	mux            sync.Mutex
}

func (f *FixMetricLimiter) DefaultLimiter(interval int64, threshold int64) *FixMetricLimiter {
	return &FixMetricLimiter{
		Interval:       interval,
		Max:            threshold,
		MetrcisCompute: countCompute,
		Distribution:   false,
		LocalWindow:    NewWindow(interval),
		mux:            sync.Mutex{},
	}
}
func (f *FixMetricLimiter) Allow() bool {
	return f.MetrcisCompute(f) <= f.Max
}
