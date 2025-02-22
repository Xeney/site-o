package database

import (
	"log"

	"github.com/dgraph-io/badger/v3"
)

var db *badger.DB

// InitDB инициализирует BadgerDB.
func InitDB(dataDir string) error {
	var err error
	// Открываем базу данных
	db, err = badger.Open(badger.DefaultOptions(dataDir))
	if err != nil {
		return err
	}
	log.Println("База данных успешно открыта!")
	return nil
}

// CloseDB закрывает соединение с базой данных.
func CloseDB() {
	if db != nil {
		db.Close()
		log.Println("База данных закрыта.")
	}
}
