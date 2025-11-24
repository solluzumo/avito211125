package dto

import "avito/pkg"

type PullRequestMergePostRequest struct {
	PullRequestId string `json:"pull_request_id"`
}

// AssertPullRequestMergePostRequestRequired checks if the required fields are not zero-ed
func AssertPullRequestMergePostRequestRequired(obj PullRequestMergePostRequest) error {
	elements := map[string]interface{}{
		"pull_request_id": obj.PullRequestId,
	}
	for name, el := range elements {
		if isZero := pkg.IsZeroValue(el); isZero {
			return &pkg.RequiredError{Field: name}
		}
	}

	return nil
}

// AssertPullRequestMergePostRequestConstraints checks if the values respects the defined constraints
func AssertPullRequestMergePostRequestConstraints(obj PullRequestMergePostRequest) error {
	return nil
}
