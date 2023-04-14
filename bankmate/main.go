package main

import (
	"go-paybro/delivery"

	_ "github.com/lib/pq"
)

func main() {
	delivery.Server().Run()
}
