package main

import (
	"github.com/qndaa/backend/src/controllers"
	"github.com/qndaa/backend/src/db"
)

func run() {
	db.Connect()
	db.AutoMigration()
	a := controllers.App{}
	a.Init()
	a.ApplyMiddlewares()
	a.StartServer()
}

func main() {
	run()
}
