package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	uuid "github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	"exam/api/models"
	"exam/pkg/helper"
)

type StaffTarifRepo struct {
	db *pgxpool.Pool
}

func NewStaffTarifRepo(db *pgxpool.Pool) *StaffTarifRepo {
	return &StaffTarifRepo{
		db: db,
	}
}

func (r *StaffTarifRepo) Create(ctx context.Context, req *models.StaffTarifCreate) (string, error) {
	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO staff_tarif(id, name, type, amountforcash, amountforcard, updated_at)
		VALUES ($1, $2, $3, $4, $5, NOW())
	`
	_, err := r.db.Exec(ctx, query,
		id,
		req.Name,
		req.Type,
		req.AmountForCash,
		req.AmountForCard,
	)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	fmt.Println(id)
	return id, nil
}

func (r *StaffTarifRepo) GetByID(ctx context.Context, req *models.StaffTarifPrimaryKey) (*models.StaffTarif, error) {

	var (
		query string

		id            sql.NullString
		name          sql.NullString
		Type          sql.NullString
		amountForCash sql.NullInt64
		amountForCard sql.NullInt64
		createdAt     sql.NullString
		updatedAt     sql.NullString
		deleted       sql.NullBool
		deletedAt     sql.NullString
	)

	query = `
		SELECT
			id,
			name,
			type,
			amountForCash,
			amountForCard,
			created_at,
			updated_at,
			deleted,
			deleted_at
		FROM staff_tarif
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&Type,
		&amountForCash,
		&amountForCard,
		&createdAt,
		&updatedAt,
		&deleted,
		&deletedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.StaffTarif{
		Id:            id.String,
		Name:          name.String,
		Type:          Type.String,
		AmountForCash: amountForCash.Int64,
		AmountForCard: amountForCard.Int64,
		CreatedAt:     createdAt.String,
		UpdatedAt:     updatedAt.String,
		Deleted:       deleted.Bool,
		DeletedAt:     deletedAt.String,
	}, nil
}

func (r *StaffTarifRepo) GetList(ctx context.Context, req *models.StaffTarifGetListRequest) (*models.StaffTarifGetListResponse, error) {

	var (
		resp    = &models.StaffTarifGetListResponse{}
		query   string
		where   = " WHERE deleted = false"
		offset  = " OFFSET 0"
		limit   = " LIMIT 10"
		ordered = " ORDER BY created_at desc"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			name,
			type,
			amountForCash,
			amountForCard,
			created_at,
			updated_at,
			deleted,
			deleted_at
		FROM staff_tarif
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Search != "" {
		where += ` AND name ILIKE '%' || '` + req.Search + `' || '%'`
	}

	query += where + ordered + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id            sql.NullString
			name          sql.NullString
			Type          sql.NullString
			amountForCash sql.NullInt64
			amountForCard sql.NullInt64
			createdAt     sql.NullString
			updatedAt     sql.NullString
			deleted       sql.NullBool
			deletedAt     sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&name,
			&Type,
			&amountForCash,
			&amountForCard,
			&createdAt,
			&updatedAt,
			&deleted,
			&deletedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.StaffTarifes = append(resp.StaffTarifes, &models.StaffTarif{
			Id:            id.String,
			Name:          name.String,
			Type:          Type.String,
			AmountForCash: amountForCash.Int64,
			AmountForCard: amountForCard.Int64,
			CreatedAt:     createdAt.String,
			UpdatedAt:     updatedAt.String,
			Deleted:       deleted.Bool,
			DeletedAt:     deletedAt.String,
		})
	}

	return resp, nil
}

func (r *StaffTarifRepo) Update(ctx context.Context, req *models.StaffTarifUpdate) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			branch
		SET
			id = :id,
			name = :name,
			type = :type,
			amountForCash = :amountForCash,
			amountForCard = :amountForCard,
			updated_at = NOW()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":            req.Id,
		"name":          req.Name,
		"type":          req.Type,
		"amountForCash": req.AmountForCash,
		"amountForCard": req.AmountForCard,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *StaffTarifRepo) Delete(ctx context.Context, req *models.StaffTarifPrimaryKey) (int64, error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			staff_tarif
		SET
			deleted = :deleted,
			deleted_at = NOW(),
			updated_at = NOW()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":      req.Id,
		"deleted": true,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}
