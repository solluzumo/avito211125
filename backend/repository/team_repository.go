package repository

import (
	"avito/domain"
	"context"
)

type TeamRepository interface {
	GetById(ctx context.Context, id string) (*domain.TeamDomain, error)
	UpdateMembers(ctx context.Context, data *domain.TeamDomain) error
	Create(ctx context.Context, data *domain.TeamDomain) error
	UnlinkUserToReviewers(ctx context.Context, userIDs []string) error
	DeactivateUsersInTeam(ctx context.Context, teamName string) error
}
