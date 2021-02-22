// +build sqlite3

package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var open = map[string]func(string) gorm.Dialector{
	"mysql":    mysql.Open,
	"postgres": postgres.Open,
	"sqlite3":  sqlite.Open,
}