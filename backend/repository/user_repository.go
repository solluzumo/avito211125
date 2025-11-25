package repository

import (
	"avito/domain"
	"avito/service/common"
	"context"
)

type UserRepository interface {
	GetById(ctx context.Context, id string) (*domain.UserDomain, error)
	UpdateUser(ctx context.Context, data *domain.UserDomain) error
	GetList(ctx context.Context, req *common.ListRequest) (*common.ListResponse[domain.UserDomain], error)
}
