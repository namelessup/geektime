package fixed_window

import (
	"sync"
	"time"
)

// FixedWindow 固定窗口周期需要的属性
type FixedWindow struct {
	mutex sync.Mutex // 保证计算临界区的值准确性

	Begin  time.Time     // 开始时间点
	Period time.Duration // 窗口周期
	Count  int64         // 请求累计次数
	Limit  int64         // 限制的次数
}

// AllowPass 固定时间窗口实现
// 每次调用次数+1
// 如果Count次数到达Limit的设置，则返回false
// 否则返回true
func (f *FixedWindow) AllowPass() bool {
	// 加锁
	f.mutex.Lock()
	defer f.mutex.Unlock()

	// 如果还未到达限制，则+1
	if f.Count < f.Limit {
		f.Count++
		return true
	}

	// 到达了限制
	// 如果当前时间-开始时间已经超过窗口期，则计数器加1
	now := time.Now()
	if now.Sub(f.Begin) > f.Period {
		// 开始重置为现在
		f.Begin = now
		f.Count = 1 // 新窗口期的第1次请求
		return true
	}

	// 如果在窗口期内则返回false
	return false
}
