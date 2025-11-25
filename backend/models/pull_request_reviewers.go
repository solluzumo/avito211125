package models

type PullRequstReviewersModel struct {
	PullRequestId string `db:"pull_request_id"`

	UserId string `db:"user_id"`
}

func (PullRequstReviewersModel) TableName() string { return "pr_reviewers" }

func (PullRequstReviewersModel) FilterFieldMap() map[string]string {
	return map[string]string{
		"pullRequestId": "pull_request_id",
		"userId":        "user_id",
	}
}
