package postgres

import (
	"avito/domain"
	"avito/models"
	"avito/service/common"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type PRPostgresRepository struct {
	base *BaseRepository[models.PullRequestModel]
}

func NewPRPostgresRepository(db *sqlx.DB) *PRPostgresRepository {
	NewBaseRepository[models.PullRequestModel](db)
	return &PRPostgresRepository{
		base: NewBaseRepository[models.PullRequestModel](db),
	}
}

// GetListForUser - получить список pull request, где пользователь является ревьером
func (r *PRPostgresRepository) GetPullRequestListForUser(ctx context.Context, userId string) (*[]domain.PullRequestDomain, error) {
	var list []models.PullRequestModel

	query := `
		SELECT 
			pr.*,
			ARRAY_AGG(prr_all.user_id ORDER BY prr_all.user_id) AS reviewers
		FROM pull_requests pr
		JOIN pr_reviewers prr_filter ON prr_filter.pull_request_id = pr.id AND prr_filter.user_id = $1
		JOIN pr_reviewers prr_all ON prr_all.pull_request_id = pr.id
		GROUP BY pr.id;
	`
	if err := r.base.DB.SelectContext(ctx, &list, query, userId); err != nil {
		return nil, err
	}

	return r.mapListModeltoListDomain(ctx, list), nil
}

func (r *PRPostgresRepository) GetList(ctx context.Context, req *common.ListRequest) (*common.ListResponse[domain.PullRequestDomain], error) {
	var list []models.PullRequestModel

	query := `
		SELECT 
			pr.*,
			ARRAY_AGG(prr_all.user_id ORDER BY prr_all.user_id) AS reviewers
		FROM pull_requests pr
		JOIN pr_reviewers prr_filter ON prr_filter.pull_request_id = pr.id 
		JOIN pr_reviewers prr_all ON prr_all.pull_request_id = pr.id
		GROUP BY pr.id;
	`
	if err := r.base.DB.SelectContext(ctx, &list, query); err != nil {
		return nil, err
	}

	return &common.ListResponse[domain.PullRequestDomain]{
		Data: *r.mapListModeltoListDomain(ctx, list),
	}, nil
}

func (r *PRPostgresRepository) GetById(ctx context.Context, id string) (*domain.PullRequestDomain, error) {
	var model models.PullRequestModel

	query := `
		SELECT
			pr.*,
			ARRAY_AGG(prr.user_id ORDER BY prr.user_id) AS reviewers
		FROM pull_requests pr
		LEFT JOIN pr_reviewers prr ON prr.pull_request_id = pr.id
		WHERE pr.id = $1
		GROUP BY pr.id;

	`
	if err := r.base.DB.GetContext(ctx, &model, query, id); err != nil {
		return nil, err
	}

	return &domain.PullRequestDomain{
		PullRequestId:     model.PullRequestId,
		PullRequestName:   model.PullRequestName,
		AuthorId:          model.AuthorId,
		Status:            model.Status,
		AssignedReviewers: []string(model.AssignedReviewers),
		CreatedAt:         model.CreatedAt,
		MergedAt:          model.MergedAt,
	}, nil
}

func (r *PRPostgresRepository) UpdatePullRequest(ctx context.Context, pr *domain.PullRequestDomain) error {
	prModel := &models.PullRequestModel{
		PullRequestId:   pr.PullRequestId,
		PullRequestName: pr.PullRequestName,
		AuthorId:        pr.AuthorId,
		Status:          pr.Status,
		CreatedAt:       pr.CreatedAt,
		MergedAt:        pr.MergedAt,
	}

	query := `
		UPDATE pull_requests
		SET pull_request_name = :pull_request_name,
			author_id = :author_id,
			status = :status,
			created_at = :created_at,
			merged_at = :merged_at
		WHERE id = :id
	`
	result, err := r.base.DB.NamedExec(query, *prModel)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("pull request with id %s is found or not updated", pr.PullRequestId)
	}

	return nil
}

func (r *PRPostgresRepository) DeleteReviewers(ctx context.Context, pullRequestID string) error {
	query := `
			DELETE FROM pr_reviewers
			WHERE pull_request_id = $1
		`

	_, err := r.base.DB.ExecContext(ctx, query, pullRequestID)
	if err != nil {
		return err
	}

	return nil
}

// UpdatePullRequestReviewers - Обновить ревьюеров
func (r *PRPostgresRepository) UpdatePullRequestReviewers(
	ctx context.Context,
	pullRequestID string,
	userIDs []string,
) error {

	// Удаляем ВСЕ старые связи
	if err := r.DeleteReviewers(ctx, pullRequestID); err != nil {
		return err
	}

	// Если новых нет, то выходим
	if len(userIDs) == 0 {
		return nil
	}

	// Создаём новые связи
	links := make([]models.PullRequstReviewersModel, len(userIDs))
	for i := range userIDs {
		links[i] = models.PullRequstReviewersModel{
			PullRequestId: pullRequestID,
			UserId:        userIDs[i],
		}
	}

	return r.LinkReviewers(ctx, links)
}

func (r *PRPostgresRepository) CreatePullRequest(ctx context.Context, pr *domain.PullRequestDomain) error {

	prModel := &models.PullRequestModel{
		PullRequestId:   pr.PullRequestId,
		PullRequestName: pr.PullRequestName,
		AuthorId:        pr.AuthorId,
		Status:          pr.Status,
		CreatedAt:       pr.CreatedAt,
		MergedAt:        pr.MergedAt,
	}

	query := `
		INSERT INTO pull_requests(
			id,
			pull_request_name,
			author_id,
			status,			
			created_at,
			merged_at
		)VALUES(
		 :id,
		 :pull_request_name,
		 :author_id,
		 :status,
		 :created_at,
		 :merged_at
		 )
	`
	_, err := r.base.DB.NamedExec(query, prModel)
	if err != nil {
		return err
	}

	links := make([]models.PullRequstReviewersModel, len(pr.AssignedReviewers))
	for i := 0; i < len(pr.AssignedReviewers); i++ {
		links[i] = models.PullRequstReviewersModel{PullRequestId: pr.PullRequestId, UserId: pr.AssignedReviewers[i]}
	}

	//Привязываем ревьюеров
	if err := r.LinkReviewers(ctx, links); err != nil {
		return err
	}

	return nil
}

// Привязать ревьеров к PR
func (r *PRPostgresRepository) LinkReviewers(ctx context.Context, links []models.PullRequstReviewersModel) error {
	query := `
		INSERT INTO pr_reviewers(
			pull_request_id,
			user_id
		)VALUES(
			:pull_request_id,
			:user_id
		)
	`
	_, err := r.base.DB.NamedExec(query, links)
	if err != nil {
		return err
	}
	return nil
}

// Перевести список моделей в список доменов
func (r *PRPostgresRepository) mapListModeltoListDomain(ctx context.Context, data []models.PullRequestModel) *[]domain.PullRequestDomain {
	domainData := make([]domain.PullRequestDomain, len(data))

	for i := 0; i < len(data); i++ {
		domainObj := &domain.PullRequestDomain{
			PullRequestId:     data[i].PullRequestId,
			PullRequestName:   data[i].PullRequestName,
			AuthorId:          data[i].AuthorId,
			Status:            data[i].Status,
			AssignedReviewers: []string(data[i].AssignedReviewers),
			CreatedAt:         data[i].CreatedAt,
			MergedAt:          data[i].MergedAt,
		}
		domainData[i] = *domainObj
	}
	return &domainData
}
