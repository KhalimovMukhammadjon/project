package postgresql

import (
	"app/api/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type couponRepo struct {
	db *pgxpool.Pool
}

func NewCouponRepo(db *pgxpool.Pool) *couponRepo {
	return &couponRepo{
		db: db,
	}
}

func (r *couponRepo) CreateCoupon(ctx context.Context, req *models.CreateCoupon) (int, error) {
	var (
		id    int
		query string
	)

	query = `
		INSERT INTO coupons(
			coupon_id,
			name,
			discount,
			discount_type,
			order_limit_price
		)
		VALUES (
			$1,
			$2,
			$3,
			$4)
			RETURNING coupon_id
			
	`
	err := r.db.QueryRow(ctx, query,
		req.Name,
		req.Discount,
		req.Discount_Type,
		req.OrderLimitPrice,
	).Scan(id)
	if err != nil {
		return 0, err
	}

	return 0, nil
}

func (r *couponRepo) GetByID(ctx context.Context, req *models.CouponPrimaryKey) (*models.Coupon, error) {

	var (
		query string
		resp  models.Coupon
	)

	query = `
		SELECT
			coupon_id, 
			name,
			discount,
			discount_type,
			order_limit_price 
		FROM coupons
		WHERE coupon_id = $1
	`

	err := r.db.QueryRow(ctx, query,
		req.CouponID,
	).Scan(
		&resp.Name,
		&resp.Discount,
		&resp.Discount_Type,
		&resp.OrderLimitPrice,
	)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (r *couponRepo) GetList(ctx context.Context, req models.GetListCouponRequest) (resp []models.Coupon, err error) {

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			coupon_id, 
			name,
			discount,
			discount_type,
			order_limit_price 
		FROM coupons
	`

	if len(req.Search) > 0 {
		filter += " AND name ILIKE '%' || '" + req.Search + "' || '%' "
	}

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += filter + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c models.Coupon
		err = rows.Scan(
			&c.CouponID,
			&c.Name,
			&c.Discount,
			&c.Discount_Type,
			&c.OrderLimitPrice,
		)
		if err != nil {
			return nil, err
		}
		resp = append(resp, c)
	}

	return resp, nil
}

func (r *couponRepo) Delete(ctx context.Context, req *models.CouponPrimaryKey) (int64, error) {
	query := `
		DELETE 
		FROM coupons
		WHERE coupon_id = $1
	`

	result, err := r.db.Exec(ctx, query, req.CouponID)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}
