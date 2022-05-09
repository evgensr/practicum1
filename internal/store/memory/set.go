package memory

import (
	"errors"
)

func (box *Box) Set(url string, short string, user string) error {

	box.RLock()
	defer box.RUnlock()

	if isDuplicate := fineDuplicate(box, short); isDuplicate {
		return errors.New("duplicate")
	}

	line := Line{
		URL:   url,
		Short: short,
		User:  user,
	}

	box.addItem(line)
	return nil

}
