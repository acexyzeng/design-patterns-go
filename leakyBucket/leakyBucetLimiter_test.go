package leakybucket

import (
	"fmt"
	"testing"
	"time"
)

type leakybucketLimiter struct {
	rate       float64 // rquests per second
	capacity   int     // max requests
	water      int     // current requests
	lastLeakMs int64   // last leak time
}

func NewLeakyBucketLimiter(rate float64, capacity int) *leakybucketLimiter {
	return &leakybucketLimiter{
		rate:       rate,
		capacity:   capacity,
		water:      0,
		lastLeakMs: time.Now().Unix(),
	}
}

func (l *leakybucketLimiter) Allow() bool {
	now := time.Now().Unix()
	elapsed := now - l.lastLeakMs

	leakAmount := int(float64(elapsed) / 1000 * l.rate)
	if leakAmount > 0 {
		if leakAmount > l.water {
			l.water = 0
		} else {
			l.water -= leakAmount
		}
	}

	if l.water >= l.capacity {
		l.water--
		return false
	}

	l.water++
	l.lastLeakMs = now
	return true
}

func Test_testLeakBucket(t *testing.T) {
	leakybucketLimiter := NewLeakyBucketLimiter(4, 10)

	for i := 1; i <= 20; i++ {
		now := time.Now().Format("15:04:05")
		if !leakybucketLimiter.Allow() {
			fmt.Printf("time: %s, allowed", now)
		} else {
			fmt.Printf("time: %s, request rejected", now)
		}
		time.Sleep(200 * time.Millisecond)
	}
}
