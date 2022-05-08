package memory

import (
	"sync"
)

type Box struct {
	sync.RWMutex
	Items           []Line
	fileStoragePath string
}
type Line struct {
	User   string `json:"user,omitempty"`
	Url    string `json:"original_url"`
	Short  string `json:"short_url"`
	Status int    `json:"status"`
}

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

func fineDuplicate(items *Box, str string) bool {

	for _, item := range items.Items {
		if item.Short == str {
			return false
		}
	}
	return true
}
