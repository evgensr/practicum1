// Package pg stores data in PG
package pg

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/evgensr/practicum1/internal/store"
)

type Box struct {
	sync.RWMutex
	Items           []Line
	fileStoragePath string
	db              *sql.DB
	chTaskDeleteURL chan []Line
}

type Line = store.Line

//New init
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
		chTaskDeleteURL: make(chan []Line),
	}

	go box.taskDelURL(box.chTaskDeleteURL)

	return box

}

//newDB create connect to DB
func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, err

}

//taskDelURL change the status of an entry to a deleted one
func (box *Box) taskDelURL(ch chan []Line) {
	for x := range ch {
		// log.Println("reader ", x)

		for _, row := range x {
			sqlStatement := `UPDATE short SET status = 1 WHERE short_url = $1 and user_id = $2;`
			_, err := box.db.Exec(sqlStatement, row.Short, row.User)
			if err != nil {
				log.Println(err)
			}
		}

	}
}

func RunMigrations(dsn string, migrationsPath string) error {
	m, err := migrate.New(migrationsPath, dsn)
	if err != nil {
		return err
	}

	err = m.Up()
	if errors.Is(err, migrate.ErrNoChange) {
		fmt.Println("Nothing to migrate")
		return nil
	}
	if err != nil {
		return err
	}

	fmt.Println("Migrated successfully")
	return nil
}
