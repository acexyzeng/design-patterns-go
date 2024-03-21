package slidingwindowlimiter

import (
	"sync"
	"testing"
	"time"
)

type SlidingWindowLimiter struct {
	windowSize  time.Duration // 滑动窗口大小
	maxRequest  int           // 每个窗口内最大请求数
	requests    []time.Time   // 窗口内的请求时间
	requestLock sync.Mutex    // 请求锁
}

func NewSlidingWindowLimiter(windowSize time.Duration, maxRequest int) *SlidingWindowLimiter {
	return &SlidingWindowLimiter{
		windowSize:  windowSize,
		maxRequest:  maxRequest,
		requests:    make([]time.Time, 0),
		requestLock: sync.Mutex{},
	}
}

func (s *SlidingWindowLimiter) Allow() bool {
	s.requestLock.Lock()
	defer s.requestLock.Unlock()

	// 移除过期请求
	now := time.Now()
	for len(s.requests) > 0 && now.Sub(s.requests[0]) > s.windowSize {
		s.requests = s.requests[1:]
	}

	// 检查请求数量是否超过阈值
	if len(s.requests) >= s.maxRequest {
		return false
	}

	s.requests = append(s.requests, now)
	return true
}

func Test_SlidingWindowLimiter(t *testing.T) {
	// 每秒2个请求
	limiter := NewSlidingWindowLimiter(500*time.Millisecond, 1)
	for i := 0; i < 10; i++ {
		now := time.Now().Format("15:04:05")
		if limiter.Allow() {
			println(now + "request success")
		} else {
			println(now + "request fail")
		}
		time.Sleep(time.Millisecond * 100)
	}
}
