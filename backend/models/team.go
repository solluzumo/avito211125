package models

type TeamModel struct {
	TeamName string `db:"team_name"`
}

func (TeamModel) TableName() string { return "team" }

func (TeamModel) FilterFieldMap() map[string]string {
	return map[string]string{
		"teamName": "team_name",
	}
}
