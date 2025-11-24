package repository

import (
	"avito/domain"
	"avito/service/common"
	"context"
)

type UserRepository interface {
	GetById(ctx context.Context, idString string) (*domain.UserDomain, error)
	UpdateUser(ctx context.Context, data *domain.UserDomain) error
	GetList(ctx context.Context, req *common.ListRequest) (*[]domain.UserDomain, error)
}
