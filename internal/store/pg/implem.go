package pg

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/evgensr/practicum1/internal/store"
)

//Get take an entry by short name
func (box *Box) Get(key string) (Line, error) {

	var line Line
	err := box.db.QueryRow("SELECT original_url, short_url, user_id, correlation_id, status FROM  short  WHERE  short_url = $1",
		key,
	).Scan(&line.URL, &line.Short, &line.User, &line.CorrelationID, &line.Status)

	// log.Println("err: ", err)
	// log.Println(line)
	return line, err

}

// GetByUser get url by user id
func (box *Box) GetByUser(idUser string) (lines []Line) {
	var line []Line
	var bLine Line

	// log.Println(idUser)
	rows, err := box.db.Query("SELECT original_url, short_url, user_id, correlation_id, status FROM  short  WHERE  user_id = $1",
		idUser,
	)
	// обязательно закрываем перед возвратом функции
	defer func() {
		errClose := rows.Close()
		if errClose != nil {
			log.Println(errClose)
		}
		errRows := rows.Err() // or modify return value
		if errClose != nil {
			log.Println(errRows)
		}
	}()

	if err != nil {
		log.Println("err ** ", err)
	}

	for rows.Next() {
		err = rows.Scan(&bLine.URL, &bLine.Short, &bLine.User, &bLine.CorrelationID, &bLine.Status)
		if err != nil {
			log.Println("Scan ", err)
		}
		// log.Println("original_url ", bLine)
		line = append(line, bLine)

	}

	return line

}

//Set write a string with data
func (box *Box) Set(line Line) error {

	var id int64
	err := box.db.QueryRow("INSERT INTO short (original_url, short_url, user_id, correlation_id) VALUES ($1, $2, $3, $4) RETURNING id",
		line.URL,
		line.Short,
		line.User,
		line.CorrelationID,
	).Scan(&id)

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

//Delete change the status of an entry to a deleted one
func (box *Box) Delete(line []Line) error {
	box.chTaskDeleteURL <- line
	return nil
}

func (box *Box) GetStats() (store.Urls, store.Users, error) {

	var urls store.Urls
	var users store.Users

	row := box.db.QueryRow("SELECT count(*) FROM short;")
	err := row.Scan(&urls)
	if err != nil {
		return 0, 0, fmt.Errorf("DB Error: %v", err)
	}

	row = box.db.QueryRow("SELECT count(DISTINCT user_id) FROM short;")
	err = row.Scan(&users)
	if err != nil {
		return 0, 0, fmt.Errorf("DB Error: %v", err)
	}

	return urls, users, nil
}

func (box *Box) Shutdown(ctx context.Context) error {

	if err := box.db.Close(); err != nil {
		return err
	}
	return nil
}
