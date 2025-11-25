package models

import (
	"time"

	"github.com/lib/pq"
)

type PullRequestModel struct {
	PullRequestId string `db:"id"`

	PullRequestName string `db:"pull_request_name"`

	AuthorId string `db:"author_id"`

	Status string `db:"status"`

	// user_id назначенных ревьюверов (0..2)
	AssignedReviewers pq.StringArray `db:"reviewers"`

	CreatedAt *time.Time `db:"created_at,omitempty"`

	MergedAt *time.Time `db:"merged_at,omitempty"`
}

func (PullRequestModel) TableName() string { return "pull_requests" }

func (PullRequestModel) FilterFieldMap() map[string]string {
	return map[string]string{
		"pullRequestId":     "id",
		"pullRequestName":   "pull_request_name",
		"authorId":          "author_id",
		"status":            "status",
		"assignedReviewers": "assigned_reviewers",
		"createdAt":         "createdAt",
		"mergedAt":          "mergedAt",
	}
}
