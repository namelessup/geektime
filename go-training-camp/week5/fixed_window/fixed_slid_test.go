package fixed_window

import (
	"sync"
	"testing"
	"time"
)

func TestFixedWindow_Allow(t *testing.T) {
	var wg sync.WaitGroup

	fw := &FixedWindow{
		Begin:  time.Now(),
		Period: 3 * time.Second,
		Count:  0,
		Limit:  10,
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
			allow := fw.AllowPass()
			t.Logf("end i=%d, allow=%v", i, allow)
		}(i)
	}

	wg.Wait()
}
