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

type SaleRepo struct {
	db *pgxpool.Pool
}

func NewSaleRepo(db *pgxpool.Pool) *SaleRepo {
	return &SaleRepo{
		db: db,
	}
}

func (r *SaleRepo) Create(ctx context.Context, req *models.SaleCreate) (string, error) {
	var (
		id    = uuid.New().String()
		query string
	)
	query = `
		INSERT INTO sales(id, branch_id, shop_assistent_id, cashier_id, price, payment_type, client_name, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7,NOW())
	`
	_, err := r.db.Exec(ctx, query,
		id,
		req.BranchId,
		req.ShopAssistentId,
		req.CashierId,
		req.Price,
		req.PaymentType,
		req.ClientName,
	)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	return id, nil
}

func (r *SaleRepo) GetByID(ctx context.Context, req *models.SalePrimaryKey) (*models.Sale, error) {

	var (
		query string

		id              sql.NullString
		branchId        sql.NullString
		shopAssistentId sql.NullString
		cashierId       sql.NullString
		price           sql.NullInt64
		paymentType     sql.NullString
		clientName      sql.NullString
		status          sql.NullString
		createdAt       sql.NullString
		updatedAt       sql.NullString
		deleted         sql.NullBool
		deletedAt       sql.NullString
	)

	query = `
		SELECT
			id,
			branch_id,
			shop_assistent_id,
			cashier_id,
			price,
			payment_type,
			client_name,
			status,
			created_at,
			updated_at,
			deleted,
			deleted_at
		FROM sales
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&branchId,
		&shopAssistentId,
		&cashierId,
		&price,
		&paymentType,
		&clientName,
		&status,
		&createdAt,
		&updatedAt,
		&deleted,
		&deletedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.Sale{
		Id:              id.String,
		BranchId:        branchId.String,
		ShopAssistentId: shopAssistentId.String,
		CashierId:       cashierId.String,
		Price:           price.Int64,
		PaymentType:     paymentType.String,
		ClientName:      clientName.String,
		Status:          status.String,
		CreatedAt:       createdAt.String,
		UpdatedAt:       updatedAt.String,
		Deleted:         deleted.Bool,
		DeletedAt:       deletedAt.String,
	}, nil
}

func (r *SaleRepo) GetList(ctx context.Context, req *models.SaleGetListRequest) (*models.SaleGetListResponse, error) {

	var (
		resp    = &models.SaleGetListResponse{}
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
			branch_id,
			shop_assistent_id,
			cashier_id,
			price,
			payment_type,
			client_name,
			status,
			created_at,
			updated_at,
			deleted,
			deleted_at
		FROM sales
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
			id              sql.NullString
			branchId        sql.NullString
			shopAssistentId sql.NullString
			cashierId       sql.NullString
			price           sql.NullInt64
			paymentType     sql.NullString
			clientName      sql.NullString
			status          sql.NullString
			createdAt       sql.NullString
			updatedAt       sql.NullString
			deleted         sql.NullBool
			deletedAt       sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&branchId,
			&shopAssistentId,
			&cashierId,
			&price,
			&paymentType,
			&clientName,
			&status,
			&createdAt,
			&updatedAt,
			&deleted,
			&deletedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Sales = append(resp.Sales, &models.Sale{
			Id:              id.String,
			BranchId:        branchId.String,
			ShopAssistentId: shopAssistentId.String,
			CashierId:       cashierId.String,
			Price:           price.Int64,
			PaymentType:     paymentType.String,
			ClientName:      clientName.String,
			Status:          status.String,
			CreatedAt:       createdAt.String,
			UpdatedAt:       updatedAt.String,
			Deleted:         deleted.Bool,
			DeletedAt:       deletedAt.String,
		})
	}

	return resp, nil
}

func (r *SaleRepo) Update(ctx context.Context, req *models.SaleUpdate) (int64, error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			sales
		SET
			id = :id,
			shop_assistent_id = :shop_assistent_id,
			cashier_id = :cashier_id,
			status = :status,
			updated_at = NOW()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":                req.Id,
		"shop_assistent_id": req.ShopAssistentId,
		"cashier_id":        req.CashierId,
		"status":            req.Status,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		log.Println("Here:", err.Error())
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *SaleRepo) Delete(ctx context.Context, req *models.SalePrimaryKey) (int64, error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			sales
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
