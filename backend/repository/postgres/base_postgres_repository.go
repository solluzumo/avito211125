package postgres

import (
	"avito/models"
	"avito/service/common"
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/jmoiron/sqlx"
)

type BaseRepository[T models.Filterable] struct {
	DB *sqlx.DB
}

func NewBaseRepository[T models.Filterable](db *sqlx.DB) *BaseRepository[T] {
	return &BaseRepository[T]{DB: db}
}

func (br *BaseRepository[T]) GetById(ctx context.Context, idString string) (*T, error) {

	var model T

	tableName := model.TableName()

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", tableName)

	err := br.DB.Get(&model, query, idString)

	if err != nil {
		return nil, err
	}

	return &model, nil
}

// Получаем маппу правильных ключей
func getMappedFields[T models.Filterable](ctx context.Context, model T, filter map[string]interface{}) (*map[string]interface{}, error) {
	fieldMap := model.FilterFieldMap()
	repoFilters := make(map[string]interface{})
	for k, v := range filter {
		col, ok := fieldMap[k]
		if !ok {
			return nil, fmt.Errorf("invalid filter field:%s", k)
		}
		repoFilters[col] = v
	}
	return &repoFilters, nil
}

func (br *BaseRepository[T]) GetList(ctx context.Context, req *common.ListRequest) (*common.ListResponse[T], error) {
	var model T
	var models []T

	tableName := model.TableName()
	baseQuery := fmt.Sprintf("FROM %s WHERE 1=1", tableName)

	args := make(map[string]interface{})

	repoFilters, err := getMappedFields(ctx, model, req.Filters)
	if err != nil {
		return nil, err
	}
	exceptionFilter, err := getMappedFields(ctx, model, req.Exception)
	if err != nil {
		return nil, err
	}

	filterQuery := "SELECT * " + baseQuery

	filterQuery = br.applyFilters(ctx, filterQuery, args, *repoFilters, false)

	filterQuery = br.applyFilters(ctx, filterQuery, args, *exceptionFilter, true)

	finalQuery, finalArgs, err := sqlx.Named(filterQuery, args)
	if err != nil {
		return nil, err
	}

	finalQuery = sqlx.Rebind(sqlx.DOLLAR, finalQuery)

	// Выборка
	if err := br.DB.SelectContext(ctx, &models, finalQuery, finalArgs...); err != nil {
		return nil, err
	}

	// Ответ
	return &common.ListResponse[T]{
		Data: models,
	}, nil
}

func (br *BaseRepository[T]) applyFilters(
	ctx context.Context,
	baseQuery string,
	params map[string]interface{},
	filters map[string]interface{},
	isException bool,
) string {

	if len(filters) == 0 {
		return baseQuery
	}

	query := baseQuery
	filterParts := []string{}

	for field, value := range filters {
		if value == nil {
			continue
		}

		v := reflect.ValueOf(value)

		if v.Kind() == reflect.Slice {
			parts := []string{}
			for i := 0; i < v.Len(); i++ {
				paramName := fmt.Sprintf("%s_%d", field, len(params))
				parts = append(parts, fmt.Sprintf("%s = :%s", field, paramName))
				params[paramName] = v.Index(i).Interface()
			}
			filterParts = append(filterParts, "("+strings.Join(parts, " OR ")+")")
		} else {
			paramName := fmt.Sprintf("%s_%d", field, len(params))
			filterParts = append(filterParts, fmt.Sprintf("%s = :%s", field, paramName))
			params[paramName] = value
		}
	}

	if len(filterParts) == 0 {
		return baseQuery
	}

	if isException {
		query += " AND NOT " + strings.Join(filterParts, " AND ")
	} else {
		query += " AND " + strings.Join(filterParts, " AND ")
	}

	return query
}
