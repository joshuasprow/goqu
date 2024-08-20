package main

import (
	"github.com/doug-martin/goqu/v9"
)

func main() {
	ds := goqu.
		Update("table").
		Set(goqu.Record{"name": "steve"}).
		Where(goqu.Ex{"id": 2})
}
