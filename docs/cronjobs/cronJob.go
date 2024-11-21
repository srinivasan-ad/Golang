package main 
import(
	"gofr.dev/pkg/gofr"
	"time"
)
func main(){
 app:= gofr.New()
 app.AddCronJob("*/10 * * * * *","" , func(ctx *gofr.Context) {
	ctx.Logger.Infof("Server is online at time : ",time.Now())
 })
 app.Run()
}