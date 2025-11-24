package api

import (
	"avito/dto"
	"avito/pkg"
	"encoding/json"
	"net/http"
	"strings"
)

// PullRequestsAPIController binds http requests to an api service and writes the service results to the http response
type PullRequestsAPIController struct {
	service      PullRequestsAPIServicer
	errorHandler pkg.ErrorHandler
}

// PullRequestsAPIOption for how the controller is set up.
type PullRequestsAPIOption func(*PullRequestsAPIController)

// WithPullRequestsAPIErrorHandler inject ErrorHandler into controller
func WithPullRequestsAPIErrorHandler(h pkg.ErrorHandler) PullRequestsAPIOption {
	return func(c *PullRequestsAPIController) {
		c.errorHandler = h
	}
}

// NewPullRequestsAPIController creates a default api controller
func NewPullRequestsAPIController(s PullRequestsAPIServicer, opts ...PullRequestsAPIOption) *PullRequestsAPIController {
	controller := &PullRequestsAPIController{
		service:      s,
		errorHandler: pkg.DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the PullRequestsAPIController
func (c *PullRequestsAPIController) Routes() Routes {
	return Routes{
		"PullRequestCreatePost": Route{
			"PullRequestCreatePost",
			strings.ToUpper("Post"),
			"/pullRequest/create",
			c.PullRequestCreatePost,
		},
		"PullRequestMergePost": Route{
			"PullRequestMergePost",
			strings.ToUpper("Post"),
			"/pullRequest/merge",
			c.PullRequestMergePost,
		},
		"PullRequestReassignPost": Route{
			"PullRequestReassignPost",
			strings.ToUpper("Post"),
			"/pullRequest/reassign",
			c.PullRequestReassignPost,
		},
	}
}

// OrderedRoutes returns all the api routes in a deterministic order for the PullRequestsAPIController
func (c *PullRequestsAPIController) OrderedRoutes() []Route {
	return []Route{
		Route{
			"PullRequestCreatePost",
			strings.ToUpper("Post"),
			"/pullRequest/create",
			c.PullRequestCreatePost,
		},
		Route{
			"PullRequestMergePost",
			strings.ToUpper("Post"),
			"/pullRequest/merge",
			c.PullRequestMergePost,
		},
		Route{
			"PullRequestReassignPost",
			strings.ToUpper("Post"),
			"/pullRequest/reassign",
			c.PullRequestReassignPost,
		},
	}
}

// PullRequestCreatePost - Создать PR и автоматически назначить до 2 ревьюверов из команды автора
func (c *PullRequestsAPIController) PullRequestCreatePost(w http.ResponseWriter, r *http.Request) {
	var pullRequestCreatePostRequestParam dto.PullRequestCreatePostRequest
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	if err := d.Decode(&pullRequestCreatePostRequestParam); err != nil {
		c.errorHandler(w, r, &pkg.ParsingError{Err: err}, nil)
		return
	}
	if err := dto.AssertPullRequestCreatePostRequestRequired(pullRequestCreatePostRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := dto.AssertPullRequestCreatePostRequestConstraints(pullRequestCreatePostRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.PullRequestCreatePost(r.Context(), pullRequestCreatePostRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = pkg.EncodeJSONResponse(result.Body, &result.Code, w)
}

// PullRequestMergePost - Пометить PR как MERGED (идемпотентная операция)
func (c *PullRequestsAPIController) PullRequestMergePost(w http.ResponseWriter, r *http.Request) {
	var pullRequestMergePostRequestParam dto.PullRequestMergePostRequest
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&pullRequestMergePostRequestParam); err != nil {
		c.errorHandler(w, r, &pkg.ParsingError{Err: err}, nil)
		return
	}
	if err := dto.AssertPullRequestMergePostRequestRequired(pullRequestMergePostRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := dto.AssertPullRequestMergePostRequestConstraints(pullRequestMergePostRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.PullRequestMergePost(r.Context(), pullRequestMergePostRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = pkg.EncodeJSONResponse(result.Body, &result.Code, w)
}

// PullRequestReassignPost - Переназначить конкретного ревьювера на другого из его команды
func (c *PullRequestsAPIController) PullRequestReassignPost(w http.ResponseWriter, r *http.Request) {
	var pullRequestReassignPostRequestParam dto.PullRequestReassignPostRequest
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&pullRequestReassignPostRequestParam); err != nil {
		c.errorHandler(w, r, &pkg.ParsingError{Err: err}, nil)
		return
	}
	if err := dto.AssertPullRequestReassignPostRequestRequired(pullRequestReassignPostRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := dto.AssertPullRequestReassignPostRequestConstraints(pullRequestReassignPostRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.PullRequestReassignPost(r.Context(), pullRequestReassignPostRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = pkg.EncodeJSONResponse(result.Body, &result.Code, w)
}
