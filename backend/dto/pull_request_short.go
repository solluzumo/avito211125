package dto

import "avito/pkg"

type PullRequestShort struct {
	PullRequestId string `json:"pull_request_id"`

	PullRequestName string `json:"pull_request_name"`

	AuthorId string `json:"author_id"`

	Status string `json:"status"`
}

// AssertPullRequestShortRequired checks if the required fields are not zero-ed
func AssertPullRequestShortRequired(obj PullRequestShort) error {
	elements := map[string]interface{}{
		"pull_request_id":   obj.PullRequestId,
		"pull_request_name": obj.PullRequestName,
		"author_id":         obj.AuthorId,
		"status":            obj.Status,
	}
	for name, el := range elements {
		if isZero := pkg.IsZeroValue(el); isZero {
			return &pkg.RequiredError{Field: name}
		}
	}

	return nil
}

// AssertPullRequestShortConstraints checks if the values respects the defined constraints
func AssertPullRequestShortConstraints(obj PullRequestShort) error {
	return nil
}
