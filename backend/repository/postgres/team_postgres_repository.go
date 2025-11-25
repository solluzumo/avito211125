package postgres

import (
	"avito/domain"
	"avito/models"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type TeamPostgresRepository struct {
	base *BaseRepository[models.TeamModel]
}

func NewTeamPostgresRepository(db *sqlx.DB) *TeamPostgresRepository {
	NewBaseRepository[models.TeamModel](db)
	return &TeamPostgresRepository{
		base: NewBaseRepository[models.TeamModel](db),
	}
}

func (r *TeamPostgresRepository) LinkMembers(ctx context.Context, links []models.TeamMemberModel) error {
	query := `
		INSERT INTO team_members(
			user_id,
			team_name
		)VALUES(
			:user_id,
			:team_name
		)
	`
	_, err := r.base.DB.NamedExec(query, links)
	if err != nil {
		return err
	}
	return nil
}

func (r *TeamPostgresRepository) Create(ctx context.Context, data *domain.TeamDomain) error {
	//создаём команду
	teamModel := &models.TeamModel{
		TeamName: data.TeamName,
	}

	_, err := r.base.DB.NamedExec(`
		INSERT INTO teams (team_name)
		VALUES (:team_name)
	`, teamModel)
	if err != nil {
		return err
	}

	userIDs := make([]string, len(data.Members))
	for i, m := range data.Members {
		userIDs[i] = m.UserId
	}
	//Обновляем команду для пользователей
	_, err = r.base.DB.ExecContext(ctx, `
		UPDATE users
		SET team_name = $1
		WHERE id = ANY($2)
	`, data.TeamName, pq.Array(userIDs))
	if err != nil {
		return err
	}

	//После изменения команды отвязываем от открытых PR
	if err = r.UnlinkUserToReviewers(ctx, userIDs); err != nil {
		return err
	}

	return nil
}

func (r *TeamPostgresRepository) DeactivateUsersInTeam(ctx context.Context, teamName string) error {
	//Обновляем команду для пользователей
	_, err := r.base.DB.ExecContext(ctx, `
		UPDATE users
		SET is_active = false
		WHERE team_name = $1
	`, teamName)
	if err != nil {
		return err
	}
	return nil
}

// UnlinkUserToReviewers - отвязать пользователя от всех его открытых PR
func (r *TeamPostgresRepository) UnlinkUserToReviewers(ctx context.Context, userIDs []string) error {
	_, err := r.base.DB.ExecContext(ctx, `
		DELETE FROM pr_reviewers prr
		USING pull_requests pr
		WHERE prr.pull_request_id = pr.id
		AND pr.status = 'OPEN'
		AND prr.user_id = ANY($1)
	`, pq.Array(userIDs))
	if err != nil {
		return err
	}

	return nil
}

func (r *TeamPostgresRepository) UpdateMembers(ctx context.Context, data *domain.TeamDomain) error {
	// Отвязываем всех пользователей этой команды
	_, err := r.base.DB.ExecContext(ctx, `
		UPDATE users
		SET team_name = NULL
		WHERE team_name = $1
	`, data.TeamName)
	if err != nil {
		return err
	}

	userIDs := make([]string, len(data.Members))
	for i, m := range data.Members {
		userIDs[i] = m.UserId
	}

	//Обновляем команду для пользователей
	_, err = r.base.DB.ExecContext(ctx, `
		UPDATE users
		SET team_name = $1
		WHERE id = ANY($2)
	`, data.TeamName, pq.Array(userIDs))
	if err != nil {
		return err
	}

	//После изменения команды отвязываем от открытых PR
	if err = r.UnlinkUserToReviewers(ctx, userIDs); err != nil {
		return err
	}

	return nil
}

func (r *TeamPostgresRepository) GetById(ctx context.Context, id string) (*domain.TeamDomain, error) {
	query := `
        SELECT 
            t.team_name,
            u.id AS user_id,
            u.username,
            u.is_active
        FROM teams t
        LEFT JOIN users u ON u.team_name = t.team_name
        WHERE t.team_name = $1;
    `

	var rows []models.TeamRows
	if err := r.base.DB.SelectContext(ctx, &rows, query, id); err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return nil, fmt.Errorf("team not found")
	}

	teamDomain := &domain.TeamDomain{
		TeamName: id,
		Members:  []domain.TeamMemberDomain{},
	}

	for _, row := range rows {
		if row.UserId != nil { // участник существует
			teamDomain.Members = append(teamDomain.Members, domain.TeamMemberDomain{
				UserId:   *row.UserId,
				Username: *row.Username,
				IsActive: *row.IsActive,
			})
		}
	}

	return teamDomain, nil

}
