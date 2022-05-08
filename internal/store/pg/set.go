package pg

import "log"

func (box *Box) Set(url string, short string, user string) {
	//id := struct {
	//	id int64
	//}{}
	var id int64
	err := box.db.QueryRow("INSERT INTO short (original_url, short_url, user_id) VALUES ($1, $2, $3) RETURNING id",
		url,
		short,
		user,
	).Scan(&id)

	log.Println("err: ", err)
	log.Println(id)

}
