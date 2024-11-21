package main

import (

	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

app.GET("/checkout",func(ctx *gofr.Context) (interface{}, error) {
	
		ctx.Logger.Info("Checkout process initiated")
		return "hello" , nil
})

	app.Run()
}
