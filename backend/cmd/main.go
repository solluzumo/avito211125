package main

import (
	"avito/interfaces/api"
	"avito/pkg"
	"avito/repository/postgres"
	"avito/service"
	"log"
	"net/http"
)

func main() {
	log.Printf("Server started")
	db, err := pkg.NewPostgres()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database is connected.")

	PullRequestRepository := postgres.NewPRPostgresRepository(db)
	PullRequestsAPIService := service.NewPullRequestsAPIService(*PullRequestRepository)
	PullRequestsAPIController := api.NewPullRequestsAPIController(PullRequestsAPIService)

	TeamsAPIService := service.NewTeamsAPIService()
	TeamsAPIController := api.NewTeamsAPIController(TeamsAPIService)

	UsersAPIService := service.NewUsersAPIService()
	UsersAPIController := api.NewUsersAPIController(UsersAPIService)

	router := api.NewRouter(PullRequestsAPIController, TeamsAPIController, UsersAPIController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
