package dto

import "avito/pkg"

type PullRequestCreatePostRequest struct {
	PullRequestId string `json:"pull_request_id"`

	PullRequestName string `json:"pull_request_name"`

	AuthorId string `json:"author_id"`
}

// AssertPullRequestCreatePostRequestRequired checks if the required fields are not zero-ed
func AssertPullRequestCreatePostRequestRequired(obj PullRequestCreatePostRequest) error {
	elements := map[string]interface{}{
		"pull_request_id":   obj.PullRequestId,
		"pull_request_name": obj.PullRequestName,
		"author_id":         obj.AuthorId,
	}
	for name, el := range elements {
		if isZero := pkg.IsZeroValue(el); isZero {
			return &pkg.RequiredError{Field: name}
		}
	}

	return nil
}

// AssertPullRequestCreatePostRequestConstraints checks if the values respects the defined constraints
func AssertPullRequestCreatePostRequestConstraints(obj PullRequestCreatePostRequest) error {
	return nil
}
