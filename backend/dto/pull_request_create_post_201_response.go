package dto

type PullRequestCreatePost201Response struct {
	Pr PullRequest `json:"pr,omitempty"`
}

// AssertPullRequestCreatePost201ResponseRequired checks if the required fields are not zero-ed
func AssertPullRequestCreatePost201ResponseRequired(obj PullRequestCreatePost201Response) error {
	if err := AssertPullRequestRequired(obj.Pr); err != nil {
		return err
	}
	return nil
}

// AssertPullRequestCreatePost201ResponseConstraints checks if the values respects the defined constraints
func AssertPullRequestCreatePost201ResponseConstraints(obj PullRequestCreatePost201Response) error {
	if err := AssertPullRequestConstraints(obj.Pr); err != nil {
		return err
	}
	return nil
}
