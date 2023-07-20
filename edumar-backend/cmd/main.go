package main

import (
	"log"
	"os"

	edumarbackend "github.com/dwirobbin/edumar-backend"
	. "github.com/dwirobbin/edumar-backend/cors"
	. "github.com/dwirobbin/edumar-backend/db/conn"
	. "github.com/dwirobbin/edumar-backend/db/drop"
	. "github.com/dwirobbin/edumar-backend/db/migrate"
	. "github.com/dwirobbin/edumar-backend/db/seed"
	. "github.com/dwirobbin/edumar-backend/helper"
	. "github.com/dwirobbin/edumar-backend/pkg/controller"
	. "github.com/dwirobbin/edumar-backend/pkg/repository"
	. "github.com/dwirobbin/edumar-backend/pkg/service"
	. "github.com/dwirobbin/edumar-backend/router"
)

func main() {
	db, err := DBConnection()
	PanicIfError(err)
	defer db.Close()

	Drop(db)
	Migrate(db)
	Seed(db)

	repositories := NewRepository(db)
	services := NewService(repositories)
	controllers := NewController(services)

	router := NewRouter(controllers)

	srv := new(edumarbackend.Server)

	port := os.Getenv("PORT")
	if port == "" {
		port = os.Getenv("API_PORT")
	}

	if err := srv.Run(port, AllowOrigin(router)); err != nil {
		log.Fatal("Error occured while running server: ", err.Error())
	}
}
