package test

import (
	"homework-3/Test/postgresql"
)

var (
	Db *postgresql.TDB
)

func init() {
	Db = postgresql.NewTDB()
}
