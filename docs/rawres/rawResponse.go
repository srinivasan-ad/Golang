package main

import (
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/http/response"
)
type username struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
  }
 func main(){
	app := gofr.New()
	app.GET("/normalres",func(c *gofr.Context) (interface{}, error) {
		users := []username{{Id: 1, Name: "Daria"}, {Id: 2, Name: "Ihor"}}
		return users,nil
	})
	 
	app.GET("/rawres",func(c *gofr.Context) (interface{}, error) {
		users := []username{{Id: 1, Name: "Daria"}, {Id: 2, Name: "Ihor"}}
		return response.Raw{Data : users},nil
	})
	app.Run()
 }