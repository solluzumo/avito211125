package service

import (
	"avito/domain"
	"avito/dto"
	"avito/pkg"
	"avito/repository"
	"context"
	"fmt"
	"net/http"
)

// TeamsAPIService is a service that implements the logic for the TeamsAPIServicer
// This service should implement the business logic for every endpoint for the TeamsAPI API.
// Include any external packages or services that will be required by this service.
type TeamsAPIService struct {
	teamRepo repository.TeamRepository
}

// NewTeamsAPIService creates a default api service
func NewTeamsAPIService(teamRepo repository.TeamRepository) *TeamsAPIService {
	return &TeamsAPIService{
		teamRepo: teamRepo,
	}
}

// TeamAddPost - Создать команду с участниками (создаёт/обновляет пользователей)
func (s *TeamsAPIService) TeamAddPost(ctx context.Context, team dto.Team) (pkg.ImplResponse, error) {
	membersDomain := make([]domain.TeamMemberDomain, len(team.Members))
	result := &dto.TeamAddPost201Response{
		Team: team,
	}

	for i := 0; i < len(team.Members); i++ {
		membersDomain[i] = domain.TeamMemberDomain{
			UserId:   team.Members[i].UserId,
			Username: team.Members[i].Username,
			IsActive: team.Members[i].IsActive,
		}
	}
	teamDomain := &domain.TeamDomain{
		TeamName: team.TeamName,
		Members:  membersDomain,
	}
	//Проверяем существует ли команда и обновляем команду, если она существует
	teamExists, _ := s.teamRepo.GetById(ctx, team.TeamName)
	if teamExists != nil {
		//Обновляем участников
		if err := s.teamRepo.UpdateMembers(ctx, teamDomain); err != nil {
			return pkg.Response(http.StatusInternalServerError, nil), fmt.Errorf("не удалось обновить участников команды с name %s :%v", teamDomain.TeamName, err)
		}

		//Собираем пользователей, которые были удалены из команды
		var usersToUnlink []string
		for _, v := range teamExists.Members {
			if contains := pkg.SliceContains(teamDomain.Members, v); !contains {
				usersToUnlink = append(usersToUnlink, v.UserId)
			}
		}

		//Отвязываем удалённых пользователей от открытых PR
		if len(usersToUnlink) > 0 {
			if err := s.teamRepo.UnlinkUserToReviewers(ctx, usersToUnlink); err != nil {
				return pkg.Response(http.StatusInternalServerError, nil), err
			}
		}

		return pkg.Response(http.StatusOK, result), nil
	}

	//Создаем команду
	if err := s.teamRepo.Create(ctx, teamDomain); err != nil {
		return pkg.Response(http.StatusInternalServerError, nil), err
	}

	return pkg.Response(http.StatusOK, result), nil
}

func (s *TeamsAPIService) TeamDeactivate(ctx context.Context, teamName string) (pkg.ImplResponse, error) {
	team, err := s.teamRepo.GetById(ctx, teamName)
	if err != nil {
		return pkg.Response(http.StatusNotFound, nil), err
	}
	//Переводим в DTO
	dtoMembers := make([]dto.TeamMember, len(team.Members))
	for i := 0; i < len(team.Members); i++ {
		dtoMembers[i] = dto.TeamMember{
			UserId:   team.Members[i].UserId,
			Username: team.Members[i].Username,
			IsActive: false,
		}
	}

	userIDS := make([]string, len(team.Members))

	for i, v := range team.Members {
		userIDS[i] = v.UserId
	}
	//Меняем статус пользователей
	if err := s.teamRepo.DeactivateUsersInTeam(ctx, teamName); err != nil {
		return pkg.Response(http.StatusInternalServerError, nil), err
	}
	//Отвязываем от OPEN PR
	if err := s.teamRepo.UnlinkUserToReviewers(ctx, userIDS); err != nil {
		return pkg.Response(http.StatusInternalServerError, nil), err
	}

	return pkg.Response(http.StatusOK, dtoMembers), nil
}

// TeamGetGet - Получить команду с участниками
func (s *TeamsAPIService) TeamGetGet(ctx context.Context, teamName string) (pkg.ImplResponse, error) {

	//Проверяем существует ли команда
	team, err := s.teamRepo.GetById(ctx, teamName)
	if err != nil {
		return pkg.Response(http.StatusNotFound, nil), err
	}
	//Переводим в DTO
	dtoMembers := make([]dto.TeamMember, len(team.Members))
	for i := 0; i < len(team.Members); i++ {
		dtoMembers[i] = dto.TeamMember{
			UserId:   team.Members[i].UserId,
			Username: team.Members[i].Username,
			IsActive: team.Members[i].IsActive,
		}
	}
	result := &dto.Team{
		TeamName: teamName,
		Members:  dtoMembers,
	}

	return pkg.Response(http.StatusOK, result), nil
}
