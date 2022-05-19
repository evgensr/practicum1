package file

import (
	"errors"
	"github.com/davecgh/go-spew/spew"
	"log"
)

func (box *Box) Get(key string) string {

	box.RLock()
	defer box.RUnlock()

	for _, u := range box.Items {
		if u.Short == key {
			return u.URL
		}
	}
	return ""
}

// GetByUser получить url по id юзера
func (box *Box) GetByUser(idUser string) (lines []Line) {

	box.RLock()
	defer box.RUnlock()

	var line []Line
	for _, u := range box.Items {
		if u.User == idUser {
			line = append(line, u)
		}
	}
	spew.Dump(line)
	return line
}

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
	err := save(box.fileStoragePath, line)
	if err != nil {
		log.Println(err)
	}
	return nil

}

func (box *Box) Delete(key string) error {
	return nil
}
