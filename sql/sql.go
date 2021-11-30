package sql

import (
	"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func InitSql() (err error) {
	db, err := gorm.Open("sqlite3", "./sql/data.db")
	defer db.Close()
	return err
}
