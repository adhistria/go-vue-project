package datastore

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/adhistria/backend/go-vue-project/internal/order/entity"
	"github.com/jmoiron/sqlx"
)

const perPage int = 5

type OrderRepository struct {
	sqlClient *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) OrderRepository {
	return OrderRepository{sqlClient: db}
}

func (o OrderRepository) Search(ctx context.Context, option entity.Option) ([]*entity.Order, error) {
	result := []*entity.Order{}

	paginationParams := ""
	if option.Page > 1 {
		paginationParams = fmt.Sprintf("LIMIT %v OFFSET %v", perPage, (option.Page-1)*(perPage))
	} else {
		paginationParams = fmt.Sprintf("LIMIT %v", perPage)
	}

	columns := `o.id as id,
	o.order_name as order_name,
	cc.company_name as company_name,
	c.name as customer_name,
	o.created_at as created_at,
	SUM(oi.price_per_unit * oi.quantity) AS total_price,
	COALESCE(SUM(oi.price_per_unit * d.delivered_quantity), 0) AS total_delivered_price`
	query := o.baseQuery(columns, option) + paginationParams + ";"
	// fmt.Println(query)
	err := o.sqlClient.Select(&result, query)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (o OrderRepository) GetTotal(option entity.Option) (*entity.OrderSummary, error) {
	query := fmt.Sprintf("SELECT SUM(total) as total_rows, SUM(total_price) as total_amount FROM (%v ", o.baseQuery("COUNT(DISTINCT (o.id)) as total, SUM(oi.price_per_unit * oi.quantity) AS total_price", option)) + ");"
	result := entity.OrderSummary{}
	err := o.sqlClient.Get(&result, query)
	fmt.Println(query)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (o OrderRepository) baseQuery(columns string, option entity.Option) string {
	whereArr := []string{}
	whereParams := ""

	australiaTimeZone := "Australia/Sydney"
	fmt.Println(australiaTimeZone, option.EndDate)

	if option.Search != "" {
		whereArr = append(whereArr, fmt.Sprintf("oi.product ILIKE '%v'", "%"+option.Search+"%"))
		whereArr = append(whereArr, fmt.Sprintf("o.order_name ILIKE '%v'", "%"+option.Search+"%"))
		whereArr = append(whereArr, fmt.Sprintf("o.customer_id ILIKE '%v'", "%"+option.Search+"%"))

	}

	emptyTime := time.Time{}
	fmt.Println(emptyTime, "empty time")
	whereTimeArr := []string{}
	if option.StartDate != emptyTime {
		whereTimeArr = append(whereTimeArr, fmt.Sprintf("o.created_at >= '%v'", option.StartDate.Format("2006-01-02")))
	}
	if option.EndDate != emptyTime {
		whereTimeArr = append(whereTimeArr, fmt.Sprintf("o.created_at <= '%v'", option.EndDate.Format("2006-01-02")))
	}

	if len(whereArr) > 0 {
		whereParams += "WHERE " + strings.Join(whereArr, " OR ")
	}
	if len(whereTimeArr) > 0 {
		if len(whereParams) > 0 {
			whereParams += " AND " + strings.Join(whereTimeArr, "AND ")
		} else {
			whereParams += "WHERE " + strings.Join(whereTimeArr, "AND ")
		}

	}

	query := fmt.Sprintf(`SELECT
			%v
		FROM
			public.orders o
		JOIN
			customers c ON o.customer_id = c.user_id
		JOIN
			customer_companies cc ON cc.company_id = c.company_id
		LEFT JOIN
			order_items oi ON oi.order_id = o.id
		LEFT JOIN
			(
				SELECT
					SUM(d.delivered_quantity) AS delivered_quantity,
					d.order_item_id
				FROM
					deliveries d
				GROUP BY
					d.order_item_id
			) d ON d.order_item_id = oi.id
		%v
		GROUP BY
			o.id, o.order_name, cc.company_name, c."name", o.created_at
		ORDER BY
			o.created_at `, columns, whereParams)
	return query
}
