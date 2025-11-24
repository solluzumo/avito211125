package dto

import "avito/pkg"

type PullRequestReassignPost200Response struct {
	Pr PullRequest `json:"pr"`

	// user_id нового ревьювера
	ReplacedBy string `json:"replaced_by"`
}

// AssertPullRequestReassignPost200ResponseRequired checks if the required fields are not zero-ed
func AssertPullRequestReassignPost200ResponseRequired(obj PullRequestReassignPost200Response) error {
	elements := map[string]interface{}{
		"pr":          obj.Pr,
		"replaced_by": obj.ReplacedBy,
	}
	for name, el := range elements {
		if isZero := pkg.IsZeroValue(el); isZero {
			return &pkg.RequiredError{Field: name}
		}
	}

	if err := AssertPullRequestRequired(obj.Pr); err != nil {
		return err
	}
	return nil
}

// AssertPullRequestReassignPost200ResponseConstraints checks if the values respects the defined constraints
func AssertPullRequestReassignPost200ResponseConstraints(obj PullRequestReassignPost200Response) error {
	if err := AssertPullRequestConstraints(obj.Pr); err != nil {
		return err
	}
	return nil
}
