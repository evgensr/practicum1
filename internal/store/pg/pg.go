package pg

import (
	"database/sql"
	"log"
	"sync"
)

type Box struct {
	sync.RWMutex
	Items           []Line
	fileStoragePath string
	db              *sql.DB
}
type Line struct {
	User   string `json:"user,omitempty"`
	Url    string `json:"original_url"`
	Short  string `json:"short_url"`
	Status int    `json:"status"`
}

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
