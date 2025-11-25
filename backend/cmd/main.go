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
	PullUserRepository := postgres.NewUserPostgresRepository(db)
	PullRequestsAPIService := service.NewPullRequestsAPIService(PullRequestRepository, PullUserRepository)
	PullRequestsAPIController := api.NewPullRequestsAPIController(PullRequestsAPIService)

	TeamRepository := postgres.NewTeamPostgresRepository(db)
	TeamsAPIService := service.NewTeamsAPIService(TeamRepository)
	TeamsAPIController := api.NewTeamsAPIController(TeamsAPIService)

	UsersAPIService := service.NewUsersAPIService(PullUserRepository, PullRequestRepository)
	UsersAPIController := api.NewUsersAPIController(UsersAPIService)

	router := api.NewRouter(PullRequestsAPIController, TeamsAPIController, UsersAPIController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
