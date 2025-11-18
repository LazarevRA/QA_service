package main

import (
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func main() {

	goose.Up()

}
