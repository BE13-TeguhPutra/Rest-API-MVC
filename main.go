package main

import (
	"BE13/MVC/config"
	"BE13/MVC/routes"
)

func main() {
	config.InitDB()
	config.InitialMigration()

	e := routes.New()

	e.Logger.Fatal(e.Start(":8080"))

}
