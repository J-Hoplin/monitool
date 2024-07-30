package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path/filepath"
)

var databaseIntialQuery = `
CREATE TABLE IF NOT EXISTS config_persistant (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  config_key VARCHAR(50) NOT NULL,
  config_value TEXT NOT NULL
);
`

type appPersistantManager struct {
	conn *sql.DB
}

var AppPersistant *appPersistantManager

func init() {
	var err error
	rootdir, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	sqliteDir := filepath.Join(rootdir, "db.sqlite")
	AppPersistant = new(appPersistantManager)
	err = AppPersistant.initConnection(sqliteDir)

	if err != nil {
		panic(err)
	}
}

// Init database connection
func (c *appPersistantManager) initConnection(database string) (err error) {
	db, err := sql.Open("sqlite3", database)

	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(databaseIntialQuery)
	if err != nil {
		return err
	}
	return nil
}
