package models

type TeamMemberModel struct {
	UserId string `db:"user_id"`

	TeamName string `db:"team_name"`
}

func (TeamMemberModel) TableName() string { return "team_members" }

func (TeamMemberModel) FilterFieldMap() map[string]string {
	return map[string]string{
		"teamName": "team_name",
		"userId":   "user_id",
	}
}
