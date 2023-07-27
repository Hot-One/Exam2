package postgres

import (
	"context"
	"database/sql"
	"fmt"

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
		return "", err
	}
	var (
		total        sql.NullInt64
		total_amount sql.NullInt64
		date         sql.NullString
		where1       string
		group        string
	)

	query1 := `
		SELECT
			COUNT(id) as total,
			SUM(amount),
			DATE(created_at) as dates
		FROM
			staff_transaction
	`
	where1 = fmt.Sprintf(" WHERE deleted = false AND staff_id = '%s'AND source_type = 'Sales'AND DATE(created_at) = CURRENT_DATE", req.StaffId)
	group = fmt.Sprintf(" GROUP BY dates")
	query1 += where1 + group
	err = r.db.QueryRow(ctx, query1).Scan(
		&total,
		&total_amount,
		&date,
	)
	if err != nil {
		fmt.Println("Query 1:", err.Error())
		return "", err
	}
	var (
		const_total        = int(total.Int64)
		const_total_amount = int(total_amount.Int64)
	)
	if const_total >= 10 && const_total_amount >= 1_500_000 {
		// Update Staff Transaction source_type
		where2 := ""
		query2 := `
			UPDATE 
				staff_transaction
			SET
				source_type = 'Bonus'
		`
		where2 = fmt.Sprintf(" WHERE staff_id = '%s'", req.StaffId)
		query2 += where2
		_, err = r.db.Exec(ctx, query2)
		if err != nil {
			fmt.Println("Query 2:", err.Error())
			return "", err
		}
		// Update staff balace
		where3 := fmt.Sprintf(" WHERE deleted = false AND id = '%s'", req.StaffId)
		var (
			balance int
		)
		query3 := `
			SELECT
				balace
			FROM
				staff
		`
		query3 += where3
		err = r.db.QueryRow(ctx, query3).Scan(&balance)
		if err != nil {
			fmt.Println("Query 3:", err.Error())
			return "", err
		}
		balance += 50000
		var (
			query4 string
			params map[string]interface{}
		)

		query4 = `
			UPDATE
				staff
			SET
				id = :id,
				balace = :balace
			WHERE id = :id
		`

		params = map[string]interface{}{
			"id":     req.StaffId,
			"balace": balance,
		}

		query, args := helper.ReplaceQueryParams(query4, params)
		fmt.Println(query)
		_, err := r.db.Exec(ctx, query, args...)
		if err != nil {
			fmt.Println("Query 4:", err.Error())
			return "", err
		}
	}

	if err != nil {
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
		resp    = &models.StaffTransactionGetListResponse{}
		query   string
		where   = " WHERE deleted = false"
		offset  = " OFFSET 0"
		limit   = " LIMIT 10"
		ordered = " ORDER BY amount"
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

	if req.SearchSales != "" {
		where += ` AND sales_id ILIKE '%' || '` + req.SearchSales + `' || '%'`
	}

	if req.SearchType != "" {
		where += ` AND type ILIKE '%' || '` + req.SearchType + `' || '%'`
	}

	if req.SearchStaff != "" {
		where += ` AND staff_id ILIKE '%' || '` + req.SearchStaff + `' || '%'`
	}

	if req.Order != "" {
		ordered += fmt.Sprintf(" ORDER BY amount %s", req.Order)
	}

	query += where + ordered + offset + limit
	fmt.Println(query)
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
		"id":          req.Id,
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
