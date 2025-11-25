package api

import (
	"avito/dto"
	"avito/pkg"
	"encoding/json"
	"net/http"
	"strings"
)

// TeamsAPIController binds http requests to an api service and writes the service results to the http response
type TeamsAPIController struct {
	service      TeamsAPIServicer
	errorHandler pkg.ErrorHandler
}

// TeamsAPIOption for how the controller is set up.
type TeamsAPIOption func(*TeamsAPIController)

// WithTeamsAPIErrorHandler inject ErrorHandler into controller
func WithTeamsAPIErrorHandler(h pkg.ErrorHandler) TeamsAPIOption {
	return func(c *TeamsAPIController) {
		c.errorHandler = h
	}
}

// NewTeamsAPIController creates a default api controller
func NewTeamsAPIController(s TeamsAPIServicer, opts ...TeamsAPIOption) *TeamsAPIController {
	controller := &TeamsAPIController{
		service:      s,
		errorHandler: pkg.DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the TeamsAPIController
func (c *TeamsAPIController) Routes() Routes {
	return Routes{
		"TeamAddPost": Route{
			"TeamAddPost",
			strings.ToUpper("Post"),
			"/team/add",
			c.TeamAddPost,
		},
		"TeamGetGet": Route{
			"TeamGetGet",
			strings.ToUpper("Get"),
			"/team/get",
			c.TeamGetGet,
		},
	}
}

// OrderedRoutes returns all the api routes in a deterministic order for the TeamsAPIController
func (c *TeamsAPIController) OrderedRoutes() []Route {
	return []Route{
		Route{
			"TeamAddPost",
			strings.ToUpper("Post"),
			"/team/add",
			c.TeamAddPost,
		},
		Route{
			"TeamGetGet",
			strings.ToUpper("Get"),
			"/team/get",
			c.TeamGetGet,
		},
		Route{
			"TeamDeactivate",
			strings.ToUpper("Post"),
			"/team/deactivate",
			c.TeamDeactivate,
		},
	}
}

// TeamDeactivate - массовая деактивация пользователей в команде
func (c *TeamsAPIController) TeamDeactivate(w http.ResponseWriter, r *http.Request) {
	var params dto.TeamDeactivatePostRequest
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	if err := d.Decode(&params); err != nil {
		c.errorHandler(w, r, &pkg.ParsingError{Err: err}, nil)
		return
	}

	// Проверка обязательных полей
	if params.TeamName == "" {
		c.errorHandler(w, r, &pkg.RequiredError{Field: "team_name"}, nil)
		return
	}

	// Вызов сервисного слоя
	result, err := c.service.TeamDeactivate(r.Context(), params.TeamName)
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}

	// Возврат результата клиенту
	_ = pkg.EncodeJSONResponse(result.Body, &result.Code, w)
}

// TeamAddPost - Создать команду с участниками (создаёт/обновляет пользователей)
func (c *TeamsAPIController) TeamAddPost(w http.ResponseWriter, r *http.Request) {
	var teamParam dto.Team
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&teamParam); err != nil {
		c.errorHandler(w, r, &pkg.ParsingError{Err: err}, nil)
		return
	}
	if err := dto.AssertTeamRequired(teamParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := dto.AssertTeamConstraints(teamParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.TeamAddPost(r.Context(), teamParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = pkg.EncodeJSONResponse(result.Body, &result.Code, w)
}

// TeamGetGet - Получить команду с участниками
func (c *TeamsAPIController) TeamGetGet(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &pkg.ParsingError{Err: err}, nil)
		return
	}
	var teamNameParam string
	if query.Has("team_name") {
		param := query.Get("team_name")

		teamNameParam = param
	} else {
		c.errorHandler(w, r, &pkg.RequiredError{Field: "team_name"}, nil)
		return
	}
	result, err := c.service.TeamGetGet(r.Context(), teamNameParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = pkg.EncodeJSONResponse(result.Body, &result.Code, w)
}
