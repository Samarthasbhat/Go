package structs

import (
	"github.com/jmoiron/sqlx"
)

type Item struct{
	Name string `db:"name"`
	When string `db:"created"`
}

func PutStats(db *sqlx.DB, item *Item) error{
	stmt := `INSERT INTO items(name, created)
			VALUES (:name, :created);`
	_, err := db.NamedExec(stmt, item)

	return err
}