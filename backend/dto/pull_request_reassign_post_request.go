package dto

import "avito/pkg"

type PullRequestReassignPostRequest struct {
	PullRequestId string `json:"pull_request_id"`

	OldUserId string `json:"old_user_id"`
}

// AssertPullRequestReassignPostRequestRequired checks if the required fields are not zero-ed
func AssertPullRequestReassignPostRequestRequired(obj PullRequestReassignPostRequest) error {
	elements := map[string]interface{}{
		"pull_request_id": obj.PullRequestId,
		"old_user_id":     obj.OldUserId,
	}
	for name, el := range elements {
		if isZero := pkg.IsZeroValue(el); isZero {
			return &pkg.RequiredError{Field: name}
		}
	}

	return nil
}

// AssertPullRequestReassignPostRequestConstraints checks if the values respects the defined constraints
func AssertPullRequestReassignPostRequestConstraints(obj PullRequestReassignPostRequest) error {
	return nil
}
