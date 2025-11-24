package pkg

type PullRequestStauts string

const (
	PullRequestStatusMerged PullRequestStauts = "MERGED"
	PullRequestStatusOpen   PullRequestStauts = "OPEN"
)
