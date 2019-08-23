package main

import (
	"github.com/ghjan/hello_gind/config"
	db "github.com/ghjan/hello_gind/database"
)

func main() {
	defer db.SqlDB.Close()
	router := config.InitRouter()
	router.Run(":8002")
}
