package memory

import (
	"github.com/evgensr/practicum1/internal/store"
	"sync"
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

func New(param string) *Box {

	fileStoragePath := param

	box := &Box{
		fileStoragePath: fileStoragePath,
	}

	return box

}

func (box *Box) addItem(item Line) []Line {
	box.Items = append(box.Items, item)
	return box.Items
}

func removeDuplicateElement(languages []string) []string {
	result := make([]string, 0, len(languages))
	temp := map[string]struct{}{}
	for _, item := range languages {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}

	return result
}

// fineDuplicate true если находим дубликат
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
