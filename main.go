package main
import "gofr.dev/pkg/gofr"
//Add your own config folder , i have removed mine :)
 func main(){
	app := gofr.New();
	app.GET("/ping", func(c *gofr.Context) (interface{}, error) {
		return "PONG" , nil

	})
app.Run()
 }

