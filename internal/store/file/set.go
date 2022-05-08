package file

import "log"

func (box *Box) Set(url string, short string, user string) {

	box.RLock()
	defer box.RUnlock()

	if isNew := fineDuplicate(box, short); isNew == false {
		return
	}

	line := Line{
		Url:   url,
		Short: short,
		User:  user,
	}

	box.addItem(line)
	err := save(box.fileStoragePath, line)
	if err != nil {
		log.Println(err)
	}

}
