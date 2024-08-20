package main

import (
	"fmt"

	"github.com/doug-martin/goqu/v9"
)

func main() {
	ds := goqu.
		Update("table").
		Set(goqu.Record{"name": "steve"}).
		Where(goqu.Ex{"id": 2})

	sql, _, _ := ds.ToSQL()

	fmt.Println(sql)
}
