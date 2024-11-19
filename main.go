package main
import (
	"gofr.dev/pkg/gofr"
	"log"
	"fmt"
	"sync"	
)
//Add your own config folder , i have removed mine :)
var payload map[string]interface{};
var lock sync.Mutex
 func main(){
	app := gofr.New();
	app.GET("/ping", func(c *gofr.Context) (interface{}, error) {
		return "PONG" , nil
   })
   app.POST("/data",func(c *gofr.Context) (interface{}, error) {
	   var requestdata map[string]interface{}
	   if err := c.Bind(&requestdata); err != nil {
		log.Println("Error binding data:", err)
		return nil, err 
	}
	fmt.Println("Data entered is :",requestdata);
	  lock.Lock()
	  defer lock.Unlock()
      payload = requestdata;
	   return payload,nil
   })
   app.GET("/data",func(c *gofr.Context) (interface{}, error) {
	lock.Lock()
	defer lock.Unlock()
	if payload == nil {
		return "No data has been recieved yet :(" , nil
	}
    return  payload , nil
   })
app.Run()
 }

