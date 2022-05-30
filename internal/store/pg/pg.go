package pg

import (
	"database/sql"
	"fmt"
	"github.com/evgensr/practicum1/internal/store"
	"log"
	"sync"
)

type Box struct {
	sync.RWMutex
	Items           []Line
	fileStoragePath string
	db              *sql.DB
	chTaskDeleteUrl chan []Line
}

type Line = store.Line

func New(param string) *Box {

	box := &Box{
		fileStoragePath: param,
	}

	db, err := newDB(param)
	if err != nil {
		log.Println(err)

		return box
	}
	// defer db.Close()

	box = &Box{
		db:              db,
		chTaskDeleteUrl: make(chan []Line),
	}

	go box.taskDelUrl(box.chTaskDeleteUrl)

	return box

}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, err

}

func (box *Box) taskDelUrl(ch chan []Line) {
	for x := range ch {
		fmt.Println("reader ", x)

		for _, row := range x {
			sqlStatement := `UPDATE short SET status = 1 WHERE short_url = $1 and user_id = $2;`
			_, err := box.db.Exec(sqlStatement, row.Short, row.User)
			if err != nil {
				log.Println(err)
			}
		}

	}
}
