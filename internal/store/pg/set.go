package pg

import (
	"errors"
	"strings"
)

func (box *Box) Set(url string, short string, user string) error {

	var id int64
	err := box.db.QueryRow("INSERT INTO short (original_url, short_url, user_id) VALUES ($1, $2, $3) RETURNING id",
		url,
		short,
		user,
	).Scan(&id)

	//log.Println("err: ", err)
	//log.Println(id)

	if strings.Contains(err.Error(), "duplicate") {
		return errors.New("duplicate")
	}

	return nil
}
