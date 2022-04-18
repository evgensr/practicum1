package memory

import (
	"github.com/davecgh/go-spew/spew"
	"sync"
)

// RWMap структура Mutex
type RWMap struct {
	mutex sync.RWMutex
	row   map[string]string
	// config *app.Config
}

// New создание структуры
func New(param string) *RWMap {

	return &RWMap{
		row: make(map[string]string),
		// config: config,
	}
}

// Get is a wrapper for getting the value from the underlying map
func (c *RWMap) Get(key string) string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.row[key]
}

// Set is a wrapper for setting the value of a key in the underlying map
func (c *RWMap) Set(key string, val string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, ok := c.row[key]; !ok {
		c.row[key] = val
	}
}

func (c *RWMap) Delete(key string) error {
	return nil
}

func (c *RWMap) Debug() {
	spew.Dump(c.row)
}
