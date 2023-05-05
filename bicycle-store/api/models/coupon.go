package models

//CREATE TYPE discount_type AS ENUM ('fix', 'percent');
// fix or percent

type Coupon struct {
	CouponID        int    `json:"coupon_id"`
	Name            string `json:"name"`
	Discount        int    `json:"discount"`
	Discount_Type   string `json:"discount_type"`
	OrderLimitPrice int    `json:"order_limit_price"`
}

type CreateCoupon struct {
	Name            string `json:"name"`
	Discount        int    `json:"discount"`
	Discount_Type   string `json:"discount_type"`
	OrderLimitPrice int    `json:"order_limit_price"`
}

type CouponPrimaryKey struct {
	CouponID int `json:"coupon_id"`
}

type GetListCouponRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

// - name
//   - discount
//   - discount_type => Фикс | Процент
//   - order_limit_price
