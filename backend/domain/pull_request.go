package domain

import (
	"time"
)

type PullRequestDomain struct {
	PullRequestId string

	PullRequestName string

	AuthorId string

	Status string

	// user_id назначенных ревьюверов (0..2)
	AssignedReviewers []string

	CreatedAt *time.Time

	MergedAt *time.Time
}
