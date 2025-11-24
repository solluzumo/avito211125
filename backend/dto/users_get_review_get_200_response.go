package dto

import "avito/pkg"

type UsersGetReviewGet200Response struct {
	UserId string `json:"user_id"`

	PullRequests []PullRequestShort `json:"pull_requests"`
}

// AssertUsersGetReviewGet200ResponseRequired checks if the required fields are not zero-ed
func AssertUsersGetReviewGet200ResponseRequired(obj UsersGetReviewGet200Response) error {
	elements := map[string]interface{}{
		"user_id":       obj.UserId,
		"pull_requests": obj.PullRequests,
	}
	for name, el := range elements {
		if isZero := pkg.IsZeroValue(el); isZero {
			return &pkg.RequiredError{Field: name}
		}
	}

	for _, el := range obj.PullRequests {
		if err := AssertPullRequestShortRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertUsersGetReviewGet200ResponseConstraints checks if the values respects the defined constraints
func AssertUsersGetReviewGet200ResponseConstraints(obj UsersGetReviewGet200Response) error {
	for _, el := range obj.PullRequests {
		if err := AssertPullRequestShortConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
