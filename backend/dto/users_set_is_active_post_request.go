package dto

import "avito/pkg"

type UsersSetIsActivePostRequest struct {
	UserId string `json:"user_id"`

	IsActive bool `json:"is_active"`
}

// AssertUsersSetIsActivePostRequestRequired checks if the required fields are not zero-ed
func AssertUsersSetIsActivePostRequestRequired(obj UsersSetIsActivePostRequest) error {
	elements := map[string]interface{}{
		"user_id":   obj.UserId,
		"is_active": obj.IsActive,
	}
	for name, el := range elements {
		if isZero := pkg.IsZeroValue(el); isZero {
			return &pkg.RequiredError{Field: name}
		}
	}

	return nil
}

// AssertUsersSetIsActivePostRequestConstraints checks if the values respects the defined constraints
func AssertUsersSetIsActivePostRequestConstraints(obj UsersSetIsActivePostRequest) error {
	return nil
}
