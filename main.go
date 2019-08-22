package main

import (
	db "github.com/ghjan/hello_gind/database"
)

func main() {
	defer db.SqlDB.Close()
	router := initRouter()
	router.Run(":8000")
}
