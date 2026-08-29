package main

import (
	"local.learn.go/db"
	"local.learn.go/router"
)
// init function runs before loading main
func init() {
}

func main() {
	db.InitPostgresDB()
	router.InitRouter().Run()
}