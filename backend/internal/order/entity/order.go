package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type Order struct {
	ID                  int             `db:"id" json:"id"`
	OrderName           string          `db:"order_name" json:"order_name"`
	CompanyName         string          `db:"company_name" json:"company_name"`
	CustomerName        string          `db:"customer_name" json:"customer_name"`
	CreatedAtTime       time.Time       `db:"created_at" json:"-"`
	CreatedAt           string          `json:"created_at"`
	TotalPrice          decimal.Decimal `db:"total_price" json:"total_price"`
	TotalDeliveredPrice decimal.Decimal `db:"total_delivered_price" json:"total_delivered_price"`
}

func ConvertDateToString(orders []Order) {
	for _, order := range orders {
		order.ConvertDateToString()
	}
}

func (o *Order) ConvertDateToString() {
	o.CreatedAt = o.CreatedAtTime.Format("Jan 2nd, 3:04 PM")
}

type Option struct {
	Search    string    `form:"search" json:"search"`
	StartDate time.Time `form:"start_date" json:"start_date"`
	EndDate   time.Time `form:"end_date" json:"end_date"`
	Page      int       `form:"page" json:"page"`
}

type OrderSummary struct {
	TotalData   int             `db:"total_rows"`
	TotalAmount decimal.Decimal `db:"total_amount"`
}

type OrderResponse struct {
	Data        []*Order        `json:"data"`
	TotalRows   int             `json:"total_rows"`
	TotalAmount decimal.Decimal `json:"total_amount"`
}
