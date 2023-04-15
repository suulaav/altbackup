package config

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/suulaav/altbackup/pkg/altutils"
)

type Connection interface {
	Connect() *sql.DB
}

func (db Db) Connect() *sql.DB {
	if db.DbUrl == "" {
		panic("Url Property Not Found In Config File")
	}
	if db.DbType == "" {
		panic("Type Property Not Found In Config File")
	}
	dbConnection, err := sql.Open(db.DbType, db.DbUrl)
	altutils.CheckError(err)
	return dbConnection
}
