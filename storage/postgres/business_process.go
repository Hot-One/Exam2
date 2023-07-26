package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

	"exam/api/models"
)

type BusinessProcessRepo struct {
	db *pgxpool.Pool
}

func NewBusinessProcessRepo(db *pgxpool.Pool) *BusinessProcessRepo {
	return &BusinessProcessRepo{
		db: db,
	}
}

func (r *BusinessProcessRepo) GetTopWorker(ctx context.Context, req *models.BusinessProcessGetRequest) (*models.BusinessProcessGetResponse, error) {

	var (
		resp    = &models.BusinessProcessGetResponse{}
		query   string
		where   = " WHERE deleted = false"
		from    = ""
		to      = ""
		ordered = " ORDER BY"
	)

	query = `
			SELECT
				name,
				branch_id,
				balace
			FROM
				staff
		`
	if req.Search != "" {
		where = fmt.Sprintf(" WHERE deleted = false AND type = '%s'", req.Search)
	}
	if req.From != "" {
		from = fmt.Sprintf(" AND created_at BETWEEN '%s'", req.From)
	}

	if req.To != "" {
		to = fmt.Sprintf(" AND '%s'", req.To)
	}

	if req.Ordered != "" {
		ordered = fmt.Sprintf(" ORDER BY balace %s", req.Ordered)
	}

	query += where + from + to + ordered
	fmt.Println(query)
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			name    sql.NullString
			Branch  sql.NullString
			Balance sql.NullInt64
		)

		err := rows.Scan(
			&name,
			&Branch,
			&Balance,
		)

		if err != nil {
			return nil, err
		}

		resp.Staffes = append(resp.Staffes, &models.BusinessProcess{
			Name:    name.String,
			Branch:  Branch.String,
			Balance: Balance.Int64,
		})
	}

	return resp, nil
}

func (r *BusinessProcessRepo) GetTopBranch(ctx context.Context,
	req *models.BusinessProcessGetRequestBranch) (*models.BusinessProcessGetResponseBranch,
	error) {
	var (
		resp    = &models.BusinessProcessGetResponseBranch{}
		query   string
		ordered = " ORDER BY"
	)

	query = `
		SELECT
			b.name as branch,
			SUM(s.price) as total_sum,
			CURRENT_DATE as DAY
		FROM branch as b
		JOIN sales as s ON s.branch_id = b.id
		GROUP BY b.name
	`

	if req.Ordered != "" {
		ordered = fmt.Sprintf(" ORDER BY total_sum %s", req.Ordered)
	}

	query += ordered
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var (
			BranchName sql.NullString
			TotalPrice sql.NullInt64
			Date       sql.NullString
		)

		err := rows.Scan(
			&BranchName,
			&TotalPrice,
			&Date,
		)

		if err != nil {
			return nil, err
		}

		resp.Branches = append(resp.Branches, &models.BusinessProcessBranch{
			Name:       BranchName.String,
			TotalPrice: TotalPrice.Int64,
			Date:       Date.String,
		})
	}

	return resp, nil
}
