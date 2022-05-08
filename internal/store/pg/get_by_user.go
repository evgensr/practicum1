package pg

import (
	"log"
)

// GetByUser получить url по id юзера
func (box *Box) GetByUser(idUser string) (lines []Line) {
	var line []Line
	var originalUrl string

	log.Println(idUser)
	rows, err := box.db.Query("SELECT original_url FROM  short  WHERE  user_id = $1",
		idUser,
	)

	if err != nil {
		log.Println("err ** ", err)
	}

	for rows.Next() {
		err = rows.Scan(&originalUrl)
		if err != nil {
			log.Println("Scan ", err)
		}
		log.Println("original_url ", originalUrl)
		line = append(line, Line{
			Url: originalUrl,
		})

	}

	log.Println("err: ", err)
	log.Println("lin: ", originalUrl)
	return line

}
