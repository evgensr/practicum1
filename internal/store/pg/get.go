package pg

import "log"

func (box *Box) Get(key string) string {

	var originalURL string
	err := box.db.QueryRow("SELECT original_url FROM  short  WHERE  short_url = $1",
		key,
	).Scan(&originalURL)

	log.Println("err: ", err)
	log.Println(originalURL)
	return originalURL

}
