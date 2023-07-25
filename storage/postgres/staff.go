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

type StaffRepo struct {
	db *pgxpool.Pool
}

func NewStaffRepo(db *pgxpool.Pool) *StaffRepo {
	return &StaffRepo{
		db: db,
	}
}

func (r *StaffRepo) Create(ctx context.Context, req *models.StaffCreate) (string, error) {
	var (
		id    = uuid.New().String()
		query string
	)
	query = `
		INSERT INTO staff(id, name, type, branch_id, tarif_id, updated_at)
		VALUES ($1, $2, $3, $4, $5, NOW())
	`
	_, err := r.db.Exec(ctx, query,
		id,
		req.Name,
		req.Type,
		req.BranchId,
		req.TarifId,
	)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	return id, nil
}

func (r *StaffRepo) GetByID(ctx context.Context, req *models.StaffPrimaryKey) (*models.Staff, error) {
	var (
		query string

		id        sql.NullString
		name      sql.NullString
		Type      sql.NullString
		BranchId  sql.NullString
		TarifId   sql.NullString
		Balance   sql.NullInt64
		createdAt sql.NullString
		updatedAt sql.NullString
		deleted   sql.NullBool
		deletedAt sql.NullString
	)

	query = `
		SELECT
			id,
			branch_id,
			tarif_id,
			type,
			name,
			balace,
			created_at,
			updated_at,
			deleted,
			deleted_at
		FROM staff
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&BranchId,
		&TarifId,
		&Type,
		&name,
		&Balance,
		&createdAt,
		&updatedAt,
		&deleted,
		&deletedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.Staff{
		Id:        id.String,
		Name:      name.String,
		Type:      Type.String,
		BranchId:  BranchId.String,
		TarifId:   TarifId.String,
		Balance:   Balance.Int64,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
		Deleted:   deleted.Bool,
		DeletedAt: deletedAt.String,
	}, nil
}

func (r *StaffRepo) GetList(ctx context.Context, req *models.StaffGetListRequest) (*models.StaffGetListResponse, error) {

	var (
		resp   = &models.StaffGetListResponse{}
		query  string
		where  = " WHERE deleted = false"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			branch_id,
			tarif_id,
			type,
			name,
			balace,
			created_at,
			updated_at,
			deleted,
			deleted_at
		FROM staff
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

	query += where + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id        sql.NullString
			name      sql.NullString
			Type      sql.NullString
			BranchId  sql.NullString
			TarifId   sql.NullString
			Balance   sql.NullInt64
			createdAt sql.NullString
			updatedAt sql.NullString
			deleted   sql.NullBool
			deletedAt sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&name,
			&Type,
			&BranchId,
			&TarifId,
			&Balance,
			&createdAt,
			&updatedAt,
			&deleted,
			&deletedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Staffes = append(resp.Staffes, &models.Staff{
			Id:        id.String,
			Name:      name.String,
			Type:      Type.String,
			BranchId:  BranchId.String,
			TarifId:   TarifId.String,
			Balance:   Balance.Int64,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
			Deleted:   deleted.Bool,
			DeletedAt: deletedAt.String,
		})
	}

	return resp, nil
}

func (r *StaffRepo) Update(ctx context.Context, req *models.StaffUpdate) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			staff
		SET
			id = :id,
			name = :name,
			type = :type,
			branch_id = :branch_id,
			tarif_id = :tarif_id,
			balance = :balance,
			updated_at = NOW()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":        req.Id,
		"name":      req.Name,
		"type":      req.Type,
		"branch_id": req.BranchId,
		"tarif_id":  req.TarifId,
		"balance":   req.Balance,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *StaffRepo) Delete(ctx context.Context, req *models.StaffPrimaryKey) (int64, error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			staff
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
