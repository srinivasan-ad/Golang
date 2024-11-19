package main
import "gofr.dev/pkg/gofr"
//Add your own config folder , i have removed mine :)
const payload = "Trying go for the first time !";
 func main(){
	app := gofr.New();
	app.GET("/ping", func(c *gofr.Context) (interface{}, error) {
		return "PONG" , nil
   })
   app.PUT("/data",func(c *gofr.Context) (interface{}, error) {
	return payload,nil
   })
app.Run()
 }

