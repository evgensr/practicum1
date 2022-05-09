package memory

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
