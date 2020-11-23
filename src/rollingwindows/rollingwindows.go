package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	windows := NewRollingWindows(4, 250*time.Millisecond)
	ticker := time.NewTicker(50 * time.Millisecond)
	i := 0
	for range ticker.C {
		i++
		if i == 30 {
			break
		}
		v := rand.Intn(10)
		fmt.Print(v)
		windows.Add(v)
	}
	time.Sleep(3 * time.Second)
	v := rand.Intn(10) + 100
	fmt.Print(v)
	windows.Add(v)
}

type bucket struct {
	sum   int
	count int
}
type RollingWindows struct {
	sync.Mutex
	size     int
	duration time.Duration
	win      []*bucket
	lastTime int64
	offset   int
}
type Option func(re *RollingWindows)

func NewBucket(n int) []*bucket {
	bk := make([]*bucket, 0)
	for i := 0; i < n; i++ {
		bk = append(bk, new(bucket))
	}
	return bk
}
func NewRollingWindows(size int, interval time.Duration, opts ...Option) *RollingWindows {
	if size < 1 {
		panic("size error")
	}
	rw := &RollingWindows{
		size:     size,
		duration: interval,
		lastTime: time.Now().UnixNano(),
		win:      NewBucket(size),
	}
	for _, opt := range opts {
		opt(rw)
	}
	return rw
}

func (rw *RollingWindows) Add(v int) {
	rw.Lock()
	defer rw.Unlock()
	rw.UpdateOffset()
	rw.win[rw.offset].add(v)
	fmt.Println(rw.win[0], rw.win[1], rw.win[2], rw.win[3])
}

func (rw *RollingWindows) UpdateOffset() {
	span := rw.span()
	if span > 0 {
		offset := (span + rw.offset) % rw.size
		if span >= rw.size {
			rw.win = NewBucket(rw.size)
			rw.offset = 0
		} else {
			for i := rw.offset + 1; i <= rw.offset+span; i++ {
				if i >= rw.size {
					rw.resetBucket(i - rw.size)
				} else {
					rw.resetBucket(i)
				}
			}
			rw.offset = offset
		}
		rw.lastTime = time.Now().UnixNano()
	}
}
func (rw *RollingWindows) resetBucket(n int) {
	rw.win[n] = new(bucket)
}

func (rw *RollingWindows) span() int {
	num := int((time.Now().UnixNano() - rw.lastTime) / rw.duration.Nanoseconds())
	if num > rw.size {
		return rw.size
	}
	return num
}

func (b *bucket) add(v int) {
	b.count++
	b.sum += v
}
