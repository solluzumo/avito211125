package common

type ListRequest struct {
	Filters map[string]interface{}

	Exception map[string]interface{}
}
