// Package memory stores data in RAM
package memory

import (
	"sync"

	"github.com/evgensr/practicum1/internal/store"
)

type Box struct {
	sync.RWMutex
	Items           []Line
	fileStoragePath string
}

type Line = store.Line

// RWMap структура Mutex
type RWMap struct {
	sync.RWMutex
	row             map[string]string
	fileStoragePath string
}

type Row struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

//New init
func New(param string) *Box {

	fileStoragePath := param

	box := &Box{
		fileStoragePath: fileStoragePath,
		Items:           make([]Line, 5000),
	}

	return box

}

//addItem a row with data in the storage is added
func (box *Box) addItem(item Line) []Line {
	box.Items = append(box.Items, item)
	return box.Items
}

//fineDuplicate true if we find a duplicate
func fineDuplicate(items *Box, str string) bool {

	if len(items.Items) == 0 {
		return false
	}
	for _, item := range items.Items {
		if item.Short == str {
			return true
		}
	}
	return false
}
