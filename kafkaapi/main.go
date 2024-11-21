package main

import (
	"gofr-server/kafkaapi/handler"
	"log"

	"gofr.dev/pkg/gofr"
)
func main(){
	app := gofr.New()
app.POST("/produce" ,func (ctx *gofr.Context) (interface{}, error) {
handler.ProducerHandler(ctx)
 return "data added",nil
})
	log.Println("Starting Gofr API server on port 8080...")
	app.Run(); 
}