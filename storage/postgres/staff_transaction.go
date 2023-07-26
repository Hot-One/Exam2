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

type StaffTransactionRepo struct {
	db *pgxpool.Pool
}

func NewStaffTransactionRepo(db *pgxpool.Pool) *StaffTransactionRepo {
	return &StaffTransactionRepo{
		db: db,
	}
}

func (r *StaffTransactionRepo) Create(ctx context.Context, req *models.StaffTransactionCreate) (string, error) {
	trx, err := r.db.Begin(ctx)
	if err != nil {
		return "", nil
	}

	defer func() {
		if err != nil {
			trx.Rollback(ctx)
		} else {
			trx.Commit(ctx)
		}
	}()
	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO staff_transaction(id, sales_id, type, source_type, text, amount, staff_id, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, NOW())
	`
	_, err = trx.Exec(ctx, query,
		id,
		req.SaleId,
		req.Type,
		req.SourceType,
		req.Text,
		req.Amount,
		helper.NewNullString(req.StaffId),
	)

	if err != nil {
		log.Println("Here", err.Error())
		return "", err
	}
	return id, nil

}

func (r *StaffTransactionRepo) GetByID(ctx context.Context, req *models.StaffTransactionPrimaryKey) (*models.StaffTransaction, error) {
	var (
		query string

		id         sql.NullString
		saleId     sql.NullString
		Type       sql.NullString
		sourceType sql.NullString
		text       sql.NullString
		amount     sql.NullInt64
		staffId    sql.NullString
		createdAt  sql.NullString
		updatedAt  sql.NullString
		deleted    sql.NullBool
		deletedAt  sql.NullString
	)

	query = `
		SELECT
			id,
			sales_id,
			type,
			source_type,
			text,
			amount,
			staff_id,
			created_at,
			updated_at,
			deleted,
			deleted_at
		FROM staff_transaction
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&saleId,
		&Type,
		&sourceType,
		&text,
		&amount,
		&staffId,
		&createdAt,
		&updatedAt,
		&deleted,
		&deletedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.StaffTransaction{
		Id:         id.String,
		SaleId:     saleId.String,
		Type:       Type.String,
		SourceType: sourceType.String,
		Text:       text.String,
		Amount:     amount.Int64,
		StaffId:    staffId.String,
		CreatedAt:  createdAt.String,
		UpdatedAt:  updatedAt.String,
		Deleted:    deleted.Bool,
		DeletedAt:  deletedAt.String,
	}, nil
}

func (r *StaffTransactionRepo) GetList(ctx context.Context, req *models.StaffTransactionGetListRequest) (*models.StaffTransactionGetListResponse, error) {
	var (
		resp   = &models.StaffTransactionGetListResponse{}
		query  string
		where  = " WHERE deleted = false"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			sales_id,
			type,
			source_type,
			text,
			amount,
			staff_id,
			created_at,
			updated_at,
			deleted,
			deleted_at
		FROM staff_transaction
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
			id         sql.NullString
			saleId     sql.NullString
			Type       sql.NullString
			sourceType sql.NullString
			text       sql.NullString
			amount     sql.NullInt64
			staffId    sql.NullString
			createdAt  sql.NullString
			updatedAt  sql.NullString
			deleted    sql.NullBool
			deletedAt  sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&saleId,
			&Type,
			&sourceType,
			&text,
			&amount,
			&staffId,
			&createdAt,
			&updatedAt,
			&deleted,
			&deletedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.StaffTransactions = append(resp.StaffTransactions, &models.StaffTransaction{
			Id:         id.String,
			SaleId:     saleId.String,
			Type:       Type.String,
			SourceType: sourceType.String,
			Text:       text.String,
			Amount:     amount.Int64,
			StaffId:    staffId.String,
			CreatedAt:  createdAt.String,
			UpdatedAt:  updatedAt.String,
			Deleted:    deleted.Bool,
			DeletedAt:  deletedAt.String,
		})
	}

	return resp, nil
}

func (r *StaffTransactionRepo) Update(ctx context.Context, req *models.StaffTransactionUpdate) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			staff_transaction
		SET
			id = :id,
			sales_id = :sales_id,
			type = :type,
			source_type = :source_type,
			text = :text,
			amount = :amount,
			staff_id = :staff_id,
			updated_at = NOW()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"sales_id":    req.SaleId,
		"type":        req.Type,
		"source_type": req.SourceType,
		"text":        req.Text,
		"amount":      req.Amount,
		"staff_id":    req.StaffId,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *StaffTransactionRepo) Delete(ctx context.Context, req *models.StaffTransactionPrimaryKey) (int64, error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			staff_transaction
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
