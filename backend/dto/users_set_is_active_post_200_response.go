package dto

type UsersSetIsActivePost200Response struct {
	User User `json:"user,omitempty"`
}

// AssertUsersSetIsActivePost200ResponseRequired checks if the required fields are not zero-ed
func AssertUsersSetIsActivePost200ResponseRequired(obj UsersSetIsActivePost200Response) error {
	if err := AssertUserRequired(obj.User); err != nil {
		return err
	}
	return nil
}

// AssertUsersSetIsActivePost200ResponseConstraints checks if the values respects the defined constraints
func AssertUsersSetIsActivePost200ResponseConstraints(obj UsersSetIsActivePost200Response) error {
	if err := AssertUserConstraints(obj.User); err != nil {
		return err
	}
	return nil
}
