package dto

import (
	"avito/pkg"
	"time"
)

type PullRequest struct {
	PullRequestId string `json:"pull_request_id"`

	PullRequestName string `json:"pull_request_name"`

	AuthorId string `json:"author_id"`

	Status string `json:"status"`

	// user_id назначенных ревьюверов (0..2)
	AssignedReviewers []string `json:"assigned_reviewers"`

	CreatedAt *time.Time `json:"createdAt,omitempty"`

	MergedAt *time.Time `json:"mergedAt,omitempty"`
}

// AssertPullRequestRequired checks if the required fields are not zero-ed
func AssertPullRequestRequired(obj PullRequest) error {
	elements := map[string]interface{}{
		"pull_request_id":    obj.PullRequestId,
		"pull_request_name":  obj.PullRequestName,
		"author_id":          obj.AuthorId,
		"status":             obj.Status,
		"assigned_reviewers": obj.AssignedReviewers,
	}
	for name, el := range elements {
		if isZero := pkg.IsZeroValue(el); isZero {
			return &pkg.RequiredError{Field: name}
		}
	}

	return nil
}

// AssertPullRequestConstraints checks if the values respects the defined constraints
func AssertPullRequestConstraints(obj PullRequest) error {
	return nil
}
