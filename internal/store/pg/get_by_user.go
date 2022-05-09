package pg

import (
	"log"
)

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
