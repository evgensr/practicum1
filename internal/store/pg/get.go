package pg

import "log"

func (box *Box) Get(key string) string {

	var originalUrl string
	err := box.db.QueryRow("SELECT original_url FROM  short  WHERE  short_url = $1",
		key,
	).Scan(&originalUrl)

	log.Println("err: ", err)
	log.Println(originalUrl)
	return originalUrl

}
