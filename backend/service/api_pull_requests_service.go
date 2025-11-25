package service

import (
	"avito/domain"
	"avito/dto"
	"avito/pkg"
	"avito/repository"
	"avito/service/common"
	"context"
	"fmt"
	"net/http"
	"time"
)

// PullRequestsAPIService is a service that implements the logic for the PullRequestsAPIServicer
// This service should implement the business logic for every endpoint for the PullRequestsAPI API.
// Include any external packages or services that will be required by this service.
type PullRequestsAPIService struct {
	pullRequestRepo repository.PullRequestRepository
	userRepo        repository.UserRepository
}

// NewPullRequestsAPIService creates a default api service
func NewPullRequestsAPIService(pullRequestRepo repository.PullRequestRepository, userRepo repository.UserRepository) *PullRequestsAPIService {
	return &PullRequestsAPIService{
		pullRequestRepo: pullRequestRepo,
		userRepo:        userRepo,
	}
}

// GetUsersWithFilter - Получить список пользователей по специальному фильтру
func (s *PullRequestsAPIService) GetUsersWithFilter(ctx context.Context, filter *common.ListRequest, reviewersCount int) ([]string, error) {
	var reviewers []string

	//Получаем список пользователей по фильтру и выделяем id для рандомного выбора
	activeUsers, err := s.userRepo.GetList(ctx, filter)
	if err != nil {
		return nil, err
	}

	activeUsersData := (*activeUsers).Data
	activeUsersIds := make([]string, len(activeUsersData))
	if len(activeUsersData) > 0 {
		for i := 0; i < len(activeUsersData); i++ {
			activeUsersIds[i] = activeUsersData[i].UserId
		}
		//Получаем одного или двух случайных пользователей
		reviewers = pkg.GetRandomElements(activeUsersIds, reviewersCount)
	}

	return reviewers, nil
}

// PullRequestCreatePost - Создать PR и автоматически назначить до 2 ревьюверов из команды автора
func (s *PullRequestsAPIService) PullRequestCreatePost(ctx context.Context, pullRequestCreatePostRequest dto.PullRequestCreatePostRequest) (pkg.ImplResponse, error) {
	timeNow := time.Now().UTC()

	//Проверяем наличие дупликата
	alreadyExsits, _ := s.pullRequestRepo.GetById(ctx, pullRequestCreatePostRequest.PullRequestId)
	if alreadyExsits != nil {
		return pkg.Response(http.StatusConflict, nil), &pkg.AlreadyExistsError{Field: "pull_request"}
	}

	//Проверяем существование user
	user, err := s.userRepo.GetById(ctx, pullRequestCreatePostRequest.AuthorId)
	if err != nil {
		return pkg.Response(http.StatusNotFound, nil), err
	}

	//Добавляем автора в исключения
	except := make(map[string]interface{}, 0)
	except["userId"] = user.UserId

	//Фильтруем по команде автора и флагу активности = true
	userFilter := make(map[string]interface{}, 0)
	userFilter["teamName"] = user.TeamName
	userFilter["isActive"] = true

	filter := &common.ListRequest{
		Filters:   userFilter,
		Exception: except,
	}
	//Получаем ДО двух ревьюеров
	reviewers, err := s.GetUsersWithFilter(ctx, filter, 2)
	if err != nil {
		return pkg.Response(http.StatusInternalServerError, nil), err
	}

	prObj := &domain.PullRequestDomain{
		PullRequestId:     pullRequestCreatePostRequest.PullRequestId,
		PullRequestName:   pullRequestCreatePostRequest.PullRequestName,
		AuthorId:          pullRequestCreatePostRequest.AuthorId,
		Status:            string(pkg.PullRequestStatusOpen),
		CreatedAt:         &timeNow,
		AssignedReviewers: reviewers,
		MergedAt:          nil,
	}

	//Создаём запись в бд - новый pull request
	if err := s.pullRequestRepo.CreatePullRequest(ctx, prObj); err != nil {
		return pkg.Response(http.StatusInternalServerError, nil), err
	}

	pullRequestResponse := &dto.PullRequest{
		PullRequestId:     prObj.PullRequestId,
		PullRequestName:   prObj.PullRequestName,
		AuthorId:          prObj.AuthorId,
		Status:            string(pkg.PullRequestStatusOpen),
		AssignedReviewers: prObj.AssignedReviewers,
		CreatedAt:         prObj.CreatedAt,
		MergedAt:          prObj.MergedAt,
	}

	return pkg.Response(201, dto.PullRequestCreatePost201Response{Pr: *pullRequestResponse}), nil
}

// PullRequestMergePost - Пометить PR как MERGED (идемпотентная операция)
func (s *PullRequestsAPIService) PullRequestMergePost(ctx context.Context, pullRequestMergePostRequest dto.PullRequestMergePostRequest) (pkg.ImplResponse, error) {

	//Проверяем существование PR и получаем его из БД
	prObj, err := s.pullRequestRepo.GetById(ctx, pullRequestMergePostRequest.PullRequestId)
	if err != nil {
		return pkg.Response(http.StatusNotFound, nil), fmt.Errorf("pull request с id %s не существует: %v", pullRequestMergePostRequest.PullRequestId, err)
	}

	//Идемпотентность
	if prObj.Status == string(pkg.PullRequestStatusMerged) {
		return pkg.Response(http.StatusOK, prObj), nil
	}

	//Обновляем статус и время мержа PR
	mergeTime := time.Now().UTC()
	prObj.MergedAt = &mergeTime
	prObj.Status = string(pkg.PullRequestStatusMerged)

	//Обновляем значение в бд
	if err := s.pullRequestRepo.UpdatePullRequest(ctx, prObj); err != nil {
		return pkg.Response(http.StatusInternalServerError, nil), fmt.Errorf("pull request с id %s не получилось обновить", pullRequestMergePostRequest.PullRequestId)
	}

	return pkg.Response(http.StatusOK, prObj), nil
}

// PullRequestReassignPost - Переназначить конкретного ревьювера на другого из его команды
func (s *PullRequestsAPIService) PullRequestReassignPost(ctx context.Context, pullRequestReassignPostRequest dto.PullRequestReassignPostRequest) (pkg.ImplResponse, error) {
	var replacedBy string

	//Проверяем существование PR и получаем его из БД
	prObj, err := s.pullRequestRepo.GetById(ctx, pullRequestReassignPostRequest.PullRequestId)
	if err != nil {
		return pkg.Response(http.StatusNotFound, nil), fmt.Errorf("pull request с id%s не существует", prObj.PullRequestId)
	}
	//Если pull request в статусе MERGED - его запрещено изменять
	if prObj.Status == string(pkg.PullRequestStatusMerged) {
		return pkg.Response(http.StatusConflict, nil), fmt.Errorf("pull request с id%s в статусе MERGED", prObj.PullRequestId)
	}

	//Проверяем что userOlderId действительно был ревьюером данного pull request
	if c := pkg.SliceContains(prObj.AssignedReviewers, pullRequestReassignPostRequest.OldUserId); !c {
		return pkg.Response(http.StatusNotFound, nil), fmt.Errorf("пользователь с id %s не является ревьюером pull request с id %s", pullRequestReassignPostRequest.OldUserId, prObj.PullRequestId)
	}

	//Проверяем существование user и получаем его из бД
	user, err := s.userRepo.GetById(ctx, prObj.AuthorId)
	if err != nil {
		return pkg.Response(http.StatusNotFound, nil), err
	}

	//Формируем исключения для user объектов(автор и старый ревьюер)
	except := make(map[string]interface{}, 0)
	exceptIds := []string{user.UserId}
	exceptIds = append(exceptIds, prObj.AssignedReviewers...)
	except["userId"] = exceptIds

	//Берём пользователей из команды пользователя со статусом "Активный"
	userFilter := make(map[string]interface{}, 0)
	userFilter["teamName"] = user.TeamName
	userFilter["isActive"] = true

	//Формируем фильтр для запроса в бд на получение списка
	filter := &common.ListRequest{
		Filters:   userFilter,
		Exception: except,
	}

	//Получаем ноль или одного ревьюера
	reviewers, err := s.GetUsersWithFilter(ctx, filter, 1)
	if err != nil {
		return pkg.Response(http.StatusInternalServerError, nil), err
	}
	if len(reviewers) == 1 {
		replacedBy = reviewers[0]
	}
	//Добавляем ревьюера, которого не заменяли
	for _, v := range prObj.AssignedReviewers {
		if v != pullRequestReassignPostRequest.OldUserId {
			reviewers = append(reviewers, v)
		}
	}
	prObj.AssignedReviewers = reviewers
	//Обновляем ревьюеров в бд
	if err := s.pullRequestRepo.UpdatePullRequestReviewers(ctx, prObj.PullRequestId, prObj.AssignedReviewers); err != nil {
		return pkg.Response(http.StatusInternalServerError, nil), err
	}

	prDto := &dto.PullRequest{
		PullRequestId:     prObj.PullRequestId,
		PullRequestName:   prObj.PullRequestName,
		AuthorId:          prObj.AuthorId,
		Status:            prObj.Status,
		AssignedReviewers: prObj.AssignedReviewers,
		CreatedAt:         prObj.CreatedAt,
		MergedAt:          prObj.MergedAt,
	}

	response := &dto.PullRequestReassignPost200Response{
		Pr:         *prDto,
		ReplacedBy: replacedBy,
	}

	return pkg.Response(http.StatusOK, response), nil
}
