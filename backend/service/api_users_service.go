package service

import (
	"avito/dto"
	"avito/pkg"
	"avito/repository"
	"context"
	"net/http"
)

type UsersAPIService struct {
	userRepo        repository.UserRepository
	pullRequestRepo repository.PullRequestRepository
}

// NewUsersAPIService creates a default api service
func NewUsersAPIService(userRepo repository.UserRepository, pullRequestRepo repository.PullRequestRepository) *UsersAPIService {
	return &UsersAPIService{
		userRepo:        userRepo,
		pullRequestRepo: pullRequestRepo,
	}
}

// UsersSetIsActivePost - Установить флаг активности пользователя
func (s *UsersAPIService) UsersSetIsActivePost(ctx context.Context, usersSetIsActivePostRequest dto.UsersSetIsActivePostRequest) (pkg.ImplResponse, error) {
	//Проверяем существование user и получаем его из БД
	user, err := s.userRepo.GetById(ctx, usersSetIsActivePostRequest.UserId)
	if err != nil {
		return pkg.Response(http.StatusNotFound, nil), err
	}

	//Меняем статус пользователя и сохраняем изменения в бд
	user.IsActive = usersSetIsActivePostRequest.IsActive
	if err := s.userRepo.UpdateUser(ctx, user); err != nil {
		return pkg.Response(http.StatusNotFound, nil), err
	}
	result := &dto.UsersSetIsActivePost200Response{
		User: dto.User{
			UserId:   (*user).UserId,
			Username: (*user).Username,
			TeamName: (*user).TeamName,
			IsActive: (*user).IsActive,
		},
	}

	return pkg.Response(http.StatusOK, result), nil
}

// UsersGetReviewGet - Получить PRы, где пользователь назначен ревьювером
func (s *UsersAPIService) UsersGetReviewGet(ctx context.Context, userId string) (pkg.ImplResponse, error) {
	//Проверяем существование user и получаем его из БД
	user, err := s.userRepo.GetById(ctx, userId)
	if err != nil {
		return pkg.Response(http.StatusNotFound, nil), err
	}

	//Получаем список PR запросов, где пользователь ревьюер
	prResult, err := s.pullRequestRepo.GetPullRequestListForUser(ctx, user.UserId)
	if err != nil {
		return pkg.Response(http.StatusInternalServerError, nil), err

	}

	return pkg.Response(http.StatusOK, prResult), nil
}
