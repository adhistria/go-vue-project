package main

import (
	"net/http"
	"os"

	"github.com/adhistria/backend/go-vue-project/internal/order/datastore"
	"github.com/adhistria/backend/go-vue-project/internal/order/entity"
	"github.com/adhistria/backend/go-vue-project/internal/order/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	dbDriver := os.Getenv("DB_DRIVER")
	db, err := sqlx.Connect(dbDriver, os.Getenv("DB_URL"))
	if err != nil {
		panic(err)
	}

	or := datastore.NewOrderRepository(db)
	os := service.NewOrderService(or)
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/orders", func(c *gin.Context) {
		option := entity.Option{}
		err := c.BindQuery(&option)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": err,
			})
		}
		res, err := os.Search(c, option)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"data":         res.Data,
				"total_rows":   res.TotalRows,
				"total_amount": res.TotalAmount,
			})
		}
	})
	r.Run()
}
