package file

import (
	"errors"
	"log"
)

func (box *Box) Set(url string, short string, user string) error {

	box.RLock()
	defer box.RUnlock()

	if isNew := fineDuplicate(box, short); !isNew {
		return errors.New("duplicate")
	}

	line := Line{
		URL:   url,
		Short: short,
		User:  user,
	}

	box.addItem(line)
	err := save(box.fileStoragePath, line)
	if err != nil {
		log.Println(err)
	}
	return nil

}
