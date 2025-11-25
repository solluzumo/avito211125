package postgres

import (
	"avito/domain"
	"avito/models"
	"avito/service/common"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type UserPostgresRepository struct {
	base *BaseRepository[models.UserModel]
}

func NewUserPostgresRepository(db *sqlx.DB) *UserPostgresRepository {
	NewBaseRepository[models.UserModel](db)
	return &UserPostgresRepository{
		base: NewBaseRepository[models.UserModel](db),
	}
}

func (r *UserPostgresRepository) GetById(ctx context.Context, id string) (*domain.UserDomain, error) {
	var model *models.UserModel
	model, err := r.base.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	res := &domain.UserDomain{
		UserId:   model.UserId,
		Username: model.Username,
		TeamName: model.TeamName,
		IsActive: model.IsActive,
	}
	return res, nil
}

func (r *UserPostgresRepository) GetList(ctx context.Context, req *common.ListRequest) (*common.ListResponse[domain.UserDomain], error) {
	prList, err := r.base.GetList(ctx, req)
	if err != nil {
		return nil, err
	}

	return &common.ListResponse[domain.UserDomain]{
		Data: *r.mapListModeltoListDomain(ctx, prList.Data),
	}, nil
}

func (r *UserPostgresRepository) UpdateUser(ctx context.Context, user *domain.UserDomain) error {

	modelObj := &models.UserModel{
		UserId:   user.UserId,
		Username: user.Username,
		TeamName: user.TeamName,
		IsActive: user.IsActive,
	}

	query := `
		UPDATE users
		SET	username = :username,
		team_name = :team_name,
		is_active = :is_active
	`
	result, err := r.base.DB.NamedExec(query, *modelObj)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("user with id %s is found or not updated", user.UserId)
	}

	return nil
}

// Перевести список моделей в список доменов
func (r *UserPostgresRepository) mapListModeltoListDomain(ctx context.Context, data []models.UserModel) *[]domain.UserDomain {
	domainData := make([]domain.UserDomain, len(data))

	for i := 0; i < len(data); i++ {
		domainObj := &domain.UserDomain{
			UserId:   data[i].UserId,
			Username: data[i].Username,
			TeamName: data[i].TeamName,
			IsActive: data[i].IsActive,
		}
		domainData[i] = *domainObj
	}
	return &domainData
}
