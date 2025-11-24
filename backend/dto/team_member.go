package dto

import "avito/pkg"

type TeamMember struct {
	UserId string `json:"user_id"`

	Username string `json:"username"`

	IsActive bool `json:"is_active"`
}

// AssertTeamMemberRequired checks if the required fields are not zero-ed
func AssertTeamMemberRequired(obj TeamMember) error {
	elements := map[string]interface{}{
		"user_id":   obj.UserId,
		"username":  obj.Username,
		"is_active": obj.IsActive,
	}
	for name, el := range elements {
		if isZero := pkg.IsZeroValue(el); isZero {
			return &pkg.RequiredError{Field: name}
		}
	}

	return nil
}

// AssertTeamMemberConstraints checks if the values respects the defined constraints
func AssertTeamMemberConstraints(obj TeamMember) error {
	return nil
}
