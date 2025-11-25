package models

type TeamRows struct {
	TeamName string  `db:"team_name"`
	UserId   *string `db:"user_id"`
	Username *string `db:"username"`
	IsActive *bool   `db:"is_active"`
}

type TeamModel struct {
	TeamName string `db:"team_name"`
}

func (TeamModel) TableName() string { return "team" }

func (TeamModel) FilterFieldMap() map[string]string {
	return map[string]string{
		"teamName": "team_name",
	}
}
