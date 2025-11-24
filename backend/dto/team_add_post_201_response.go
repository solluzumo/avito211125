package dto

type TeamAddPost201Response struct {
	Team Team `json:"team,omitempty"`
}

// AssertTeamAddPost201ResponseRequired checks if the required fields are not zero-ed
func AssertTeamAddPost201ResponseRequired(obj TeamAddPost201Response) error {
	if err := AssertTeamRequired(obj.Team); err != nil {
		return err
	}
	return nil
}

// AssertTeamAddPost201ResponseConstraints checks if the values respects the defined constraints
func AssertTeamAddPost201ResponseConstraints(obj TeamAddPost201Response) error {
	if err := AssertTeamConstraints(obj.Team); err != nil {
		return err
	}
	return nil
}
