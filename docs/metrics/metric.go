package main

import (

	"gofr.dev/pkg/gofr"
	"time"
)

func main() {
	app := gofr.New()
	app.Metrics().NewCounter("transaction_success", "used to track the count of successful transactions")

	app.POST("/transaction", func(ctx *gofr.Context) (interface{}, error) {
		ctx.Metrics().IncrementCounter(ctx, "transaction_success")

		return "Transaction Successful", nil
	})
 
	app.Metrics().NewUpDownCounter("total_credit_day_sale", "used to track the total credit sales in a day")

	app.POST("/sale", func(ctx *gofr.Context) (interface{}, error) {
		ctx.Metrics().DeltaUpDownCounter(ctx, "total_credit_day_sale", 1000)

		return "Sale Completed", nil
	})

	app.Metrics().NewHistogram("transaction_time", "used to track the time taken by a transaction",
		5, 10, 15, 20, 25, 35)

	app.POST("/transaction", func(ctx *gofr.Context) (interface{}, error) {
		transactionStartTime := time.Now()

		// transaction logic

		tranTime := time.Now().Sub(transactionStartTime).Milliseconds()

		ctx.Metrics().RecordHistogram(ctx, "transaction_time", float64(tranTime))

		return "Transaction Completed", nil
	})

	app.Metrics().NewGauge("product_stock", "used to track the number of products in stock")

	app.POST("/sale", func(ctx *gofr.Context) (interface{}, error) {
		ctx.Metrics().SetGauge("product_stock", 10)

		return "Sale Completed", nil
	})

	app.Run()
}
