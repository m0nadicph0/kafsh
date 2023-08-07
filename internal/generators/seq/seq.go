package seq

import "sync"

type sequence struct {
	mu    sync.Mutex
	value int
}

func (c *sequence) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *sequence) Get() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func (c *sequence) Reset() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value = 0
}

var seq sequence

func Get() int {
	return seq.Get()
}

func Incr() {
	seq.Increment()
}

func Reset() {
	seq.Reset()
}
