package slide_window

import (
	"sync"
	"time"
)

type SlideWindow struct {
	mutex sync.Mutex

	Interval       int // 滑动窗口时长
	BucketInterval int // 每个桶时长
	BucketSize     int // 桶的个数
	BucketIndex    int // 桶下标
	BucketCounter  []int
	PrvBucketTime  time.Time
}

func (s *SlideWindow) Allow() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	now := time.Now()

	m := s.Interval / s.BucketInterval
	// 如果间隔时间超过一个窗口的时间 当前窗口置0 指向下一个窗口
	if now.Unix()-s.PrvBucketTime.Unix() > int64(m) {
		s.BucketCounter[s.BucketIndex] = 0
		s.BucketIndex = (s.BucketIndex + 1) % s.BucketSize
		s.PrvBucketTime = now
	}

	// 当前窗口未满则正常计数
	if s.BucketCounter[s.BucketIndex] < s.BucketSize {
		s.BucketCounter[s.BucketIndex]++
		return true
	}
	return false
}
