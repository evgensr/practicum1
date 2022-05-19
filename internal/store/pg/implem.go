package pg

import (
	"errors"
	"log"
	"strings"
)

func (box *Box) Get(key string) string {

	var originalURL string
	err := box.db.QueryRow("SELECT original_url FROM  short  WHERE  short_url = $1",
		key,
	).Scan(&originalURL)

	log.Println("err: ", err)
	log.Println(originalURL)
	return originalURL

}

// GetByUser получить url по id юзера
func (box *Box) GetByUser(idUser string) (lines []Line) {
	var line []Line
	var originalURL string

	log.Println(idUser)
	rows, err := box.db.Query("SELECT original_url FROM  short  WHERE  user_id = $1",
		idUser,
	)

	if err != nil {
		log.Println("err ** ", err)
	}

	for rows.Next() {
		err = rows.Scan(&originalURL)
		if err != nil {
			log.Println("Scan ", err)
		}
		log.Println("original_url ", originalURL)
		line = append(line, Line{
			URL: originalURL,
		})

	}

	log.Println("err: ", err)
	log.Println("lin: ", originalURL)
	return line

}

func (box *Box) Set(url string, short string, user string) error {

	var id int64
	err := box.db.QueryRow("INSERT INTO short (original_url, short_url, user_id) VALUES ($1, $2, $3) RETURNING id",
		url,
		short,
		user,
	).Scan(&id)

	//log.Println("err: ", err)
	//log.Println(id)

	// log.Println(strings.Contains(err.Error(), "duplicate"))
	duplicate := false
	if err != nil {
		duplicate = strings.Contains(err.Error(), "duplicate")
	}

	if duplicate {
		return errors.New("duplicate")
	}

	return nil
}

func (box *Box) Delete(key string) error {
	return nil
}
