package repository

import (
	"avito/domain"
	"avito/service/common"
	"context"
)

type PullRequestRepository interface {
	GetById(ctx context.Context, id string) (*domain.PullRequestDomain, error)
	UpdatePullRequest(ctx context.Context, pR *domain.PullRequestDomain) error
	UpdatePullRequestReviewers(ctx context.Context, pullRequestID string, userIDS []string) error
	GetList(ctx context.Context, req *common.ListRequest) (*common.ListResponse[domain.PullRequestDomain], error)
	GetPullRequestListForUser(ctx context.Context, userId string) (*[]domain.PullRequestDomain, error)
	CreatePullRequest(ctx context.Context, pR *domain.PullRequestDomain) error
}
