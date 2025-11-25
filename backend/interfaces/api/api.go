package api

import (
	"avito/dto"
	"avito/pkg"
	"context"
	"net/http"
)

// PullRequestsAPIRouter defines the required methods for binding the api requests to a responses for the PullRequestsAPI
// The PullRequestsAPIRouter implementation should parse necessary information from the http request,
// pass the data to a PullRequestsAPIServicer to perform the required actions, then write the service results to the http response.
type PullRequestsAPIRouter interface {
	PullRequestCreatePost(http.ResponseWriter, *http.Request)
	PullRequestMergePost(http.ResponseWriter, *http.Request)
	PullRequestReassignPost(http.ResponseWriter, *http.Request)
}

// TeamsAPIRouter defines the required methods for binding the api requests to a responses for the TeamsAPI
// The TeamsAPIRouter implementation should parse necessary information from the http request,
// pass the data to a TeamsAPIServicer to perform the required actions, then write the service results to the http response.
type TeamsAPIRouter interface {
	TeamAddPost(http.ResponseWriter, *http.Request)
	TeamGetGet(http.ResponseWriter, *http.Request)
}

// UsersAPIRouter defines the required methods for binding the api requests to a responses for the UsersAPI
// The UsersAPIRouter implementation should parse necessary information from the http request,
// pass the data to a UsersAPIServicer to perform the required actions, then write the service results to the http response.
type UsersAPIRouter interface {
	UsersSetIsActivePost(http.ResponseWriter, *http.Request)
	UsersGetReviewGet(http.ResponseWriter, *http.Request)
}

// PullRequestsAPIServicer defines the api actions for the PullRequestsAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type PullRequestsAPIServicer interface {
	PullRequestCreatePost(context.Context, dto.PullRequestCreatePostRequest) (pkg.ImplResponse, error)
	PullRequestMergePost(context.Context, dto.PullRequestMergePostRequest) (pkg.ImplResponse, error)
	PullRequestReassignPost(context.Context, dto.PullRequestReassignPostRequest) (pkg.ImplResponse, error)
}

// TeamsAPIServicer defines the api actions for the TeamsAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type TeamsAPIServicer interface {
	TeamAddPost(context.Context, dto.Team) (pkg.ImplResponse, error)
	TeamGetGet(context.Context, string) (pkg.ImplResponse, error)
	TeamDeactivate(context.Context, string) (pkg.ImplResponse, error)
}

// UsersAPIServicer defines the api actions for the UsersAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type UsersAPIServicer interface {
	UsersSetIsActivePost(context.Context, dto.UsersSetIsActivePostRequest) (pkg.ImplResponse, error)
	UsersGetReviewGet(context.Context, string) (pkg.ImplResponse, error)
}
