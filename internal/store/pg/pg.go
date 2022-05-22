package pg

import (
	"database/sql"
	"github.com/evgensr/practicum1/internal/store"
	"log"
	"sync"
)

type Box struct {
	sync.RWMutex
	Items           []Line
	fileStoragePath string
	db              *sql.DB
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
		db: db,
	}

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
