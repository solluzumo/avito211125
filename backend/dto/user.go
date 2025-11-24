package dto

import "avito/pkg"

type User struct {
	UserId string `json:"user_id"`

	Username string `json:"username"`

	TeamName string `json:"team_name"`

	IsActive bool `json:"is_active"`
}

// AssertUserRequired checks if the required fields are not zero-ed
func AssertUserRequired(obj User) error {
	elements := map[string]interface{}{
		"user_id":   obj.UserId,
		"username":  obj.Username,
		"team_name": obj.TeamName,
		"is_active": obj.IsActive,
	}
	for name, el := range elements {
		if isZero := pkg.IsZeroValue(el); isZero {
			return &pkg.RequiredError{Field: name}
		}
	}

	return nil
}

// AssertUserConstraints checks if the values respects the defined constraints
func AssertUserConstraints(obj User) error {
	return nil
}
