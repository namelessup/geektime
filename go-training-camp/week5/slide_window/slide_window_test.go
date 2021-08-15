package slide_window

import (
	"sync"
	"testing"
	"time"
)

func TestSlideWindow_Allow(t *testing.T) {
	var wg sync.WaitGroup

	sw := &SlideWindow{
		Interval:       3,
		BucketInterval: 1,
		BucketSize:     5,
		BucketCounter:  make([]int, 5, 5),
	}
	time.Sleep(1 * time.Second)
	for i := 0; i < 100; i++ {
		// 每10个线程睡眠1s
		if i%10 == 0 {
			time.Sleep(1 * time.Second)
		}
		wg.Add(1)
		t.Logf("start i=%d", i)
		go func(i int) {
			defer wg.Done()
			allow := sw.Allow()
			t.Logf("end i=%d, allow=%v", i, allow)
		}(i)
	}

	wg.Wait()
}
