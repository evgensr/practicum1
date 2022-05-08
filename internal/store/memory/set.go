package memory

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

}
