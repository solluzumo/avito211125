package models

type UserModel struct {
	UserId string `db:"id"`

	Username string `db:"username"`

	TeamName string `db:"team_name"`

	IsActive bool `db:"is_active"`
}

func (UserModel) TableName() string { return "users" }

func (UserModel) FilterFieldMap() map[string]string {
	return map[string]string{
		"userId":   "id",
		"username": "user_name",
		"teamName": "team_name",
		"isActive": "is_active",
	}
}
