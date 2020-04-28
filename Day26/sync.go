package sync

import (
	"sync"
)

type Counter struct{
	m sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.m.Lock()
	defer c.m.Unlock()
	c.value++	
}

func (c *Counter) Value() int {
	return c.value
}