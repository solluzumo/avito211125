package api

import (
	"avito/dto"
	"avito/pkg"
	"encoding/json"
	"net/http"
	"strings"
)

// UsersAPIController binds http requests to an api service and writes the service results to the http response
type UsersAPIController struct {
	service      UsersAPIServicer
	errorHandler pkg.ErrorHandler
}

// UsersAPIOption for how the controller is set up.
type UsersAPIOption func(*UsersAPIController)

// WithUsersAPIErrorHandler inject ErrorHandler into controller
func WithUsersAPIErrorHandler(h pkg.ErrorHandler) UsersAPIOption {
	return func(c *UsersAPIController) {
		c.errorHandler = h
	}
}

// NewUsersAPIController creates a default api controller
func NewUsersAPIController(s UsersAPIServicer, opts ...UsersAPIOption) *UsersAPIController {
	controller := &UsersAPIController{
		service:      s,
		errorHandler: pkg.DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the UsersAPIController
func (c *UsersAPIController) Routes() Routes {
	return Routes{
		"UsersSetIsActivePost": Route{
			"UsersSetIsActivePost",
			strings.ToUpper("Post"),
			"/users/setIsActive",
			c.UsersSetIsActivePost,
		},
		"UsersGetReviewGet": Route{
			"UsersGetReviewGet",
			strings.ToUpper("Get"),
			"/users/getReview",
			c.UsersGetReviewGet,
		},
	}
}

// OrderedRoutes returns all the api routes in a deterministic order for the UsersAPIController
func (c *UsersAPIController) OrderedRoutes() []Route {
	return []Route{
		Route{
			"UsersSetIsActivePost",
			strings.ToUpper("Post"),
			"/users/setIsActive",
			c.UsersSetIsActivePost,
		},
		Route{
			"UsersGetReviewGet",
			strings.ToUpper("Get"),
			"/users/getReview",
			c.UsersGetReviewGet,
		},
	}
}

// UsersSetIsActivePost - Установить флаг активности пользователя
func (c *UsersAPIController) UsersSetIsActivePost(w http.ResponseWriter, r *http.Request) {
	var usersSetIsActivePostRequestParam dto.UsersSetIsActivePostRequest
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&usersSetIsActivePostRequestParam); err != nil {
		c.errorHandler(w, r, &pkg.ParsingError{Err: err}, nil)
		return
	}

	if err := dto.AssertUsersSetIsActivePostRequestRequired(usersSetIsActivePostRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := dto.AssertUsersSetIsActivePostRequestConstraints(usersSetIsActivePostRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}

	result, err := c.service.UsersSetIsActivePost(r.Context(), usersSetIsActivePostRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = pkg.EncodeJSONResponse(result.Body, &result.Code, w)
}

// UsersGetReviewGet - Получить PR'ы, где пользователь назначен ревьювером
func (c *UsersAPIController) UsersGetReviewGet(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &pkg.ParsingError{Err: err}, nil)
		return
	}
	var userIdParam string
	if query.Has("user_id") {
		param := query.Get("user_id")

		userIdParam = param
	} else {
		c.errorHandler(w, r, &pkg.RequiredError{Field: "user_id"}, nil)
		return
	}
	result, err := c.service.UsersGetReviewGet(r.Context(), userIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = pkg.EncodeJSONResponse(result.Body, &result.Code, w)
}
