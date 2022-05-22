package memory

import (
	"errors"
	"github.com/davecgh/go-spew/spew"
	"log"
)

func (box *Box) Get(key string) (Line, error) {

	box.RLock()
	defer box.RUnlock()

	for _, u := range box.Items {
		if u.Short == key {
			return u, nil
		}
	}
	return Line{}, errors.New("not found")
}

// GetByUser получить url по id юзера
func (box *Box) GetByUser(idUser string) (lines []Line) {

	box.RLock()
	defer box.RUnlock()

	log.Println("idUser ", idUser)
	spew.Dump(box.Items)

	var line []Line
	for _, u := range box.Items {
		if u.User == idUser {
			line = append(line, u)
		}
	}
	spew.Dump(line)
	return line
}

func (box *Box) Set(line Line) error {

	box.RLock()
	defer box.RUnlock()

	if isDuplicate := fineDuplicate(box, line.Short); isDuplicate {
		return errors.New("duplicate")
	}

	newLine := Line{
		URL:           line.URL,
		Short:         line.Short,
		User:          line.User,
		CorrelationID: line.CorrelationID,
		Status:        line.Status,
	}

	box.addItem(newLine)
	return nil

}

func (box *Box) Delete(line []Line) error {
	box.RLock()
	defer box.RUnlock()

	for ui, u := range box.Items {
		for _, l := range line {
			if u.User == l.User && u.Short == l.Short {
				box.Items[ui].Status = 1
			}
		}
	}
	return nil
}
