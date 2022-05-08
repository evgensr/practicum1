package file

import "github.com/davecgh/go-spew/spew"

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
