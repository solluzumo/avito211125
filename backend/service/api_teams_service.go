package service

import (
	"avito/dto"
	"avito/pkg"
	"context"
	"errors"
	"net/http"
)

// TeamsAPIService is a service that implements the logic for the TeamsAPIServicer
// This service should implement the business logic for every endpoint for the TeamsAPI API.
// Include any external packages or services that will be required by this service.
type TeamsAPIService struct {
}

// NewTeamsAPIService creates a default api service
func NewTeamsAPIService() *TeamsAPIService {
	return &TeamsAPIService{}
}

// TeamAddPost - Создать команду с участниками (создаёт/обновляет пользователей)
func (s *TeamsAPIService) TeamAddPost(ctx context.Context, team dto.Team) (pkg.ImplResponse, error) {

	return pkg.Response(http.StatusNotImplemented, nil), errors.New("TeamAddPost method not implemented")
}

// TeamGetGet - Получить команду с участниками
func (s *TeamsAPIService) TeamGetGet(ctx context.Context, teamName string) (pkg.ImplResponse, error) {

	return pkg.Response(http.StatusNotImplemented, nil), errors.New("TeamGetGet method not implemented")
}
