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
		ordered = " ORDERED BY"
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
		where += ` AND type = ` + req.Search
	}
	if req.From != "" {
		from = fmt.Sprintf(" AND created_at BETWEEN %s", req.From)
	}

	if req.To != "" {
		to = fmt.Sprintf(" AND %s", req.To)
	}

	if req.Ordered != "" {
		ordered = fmt.Sprintf(" ORDER BY balace %s", req.Ordered)
	}

	query += where + from + to + ordered
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
			&Branch,
			&name,
			&Balance,
		)

		if err != nil {
			return nil, err
		}

		resp.Staffes = append(resp.Staffes, &models.Staff{
			Name:     name.String,
			BranchId: Branch.String,
			Balance:  Balance.Int64,
		})
	}

	return resp, nil
}
