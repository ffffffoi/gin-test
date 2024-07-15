package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Ip string `json:"ip"`
}


func main(){
	r := gin.Default()

	r.GET("/",func(c *gin.Context) {
		ip := c.Request.RemoteAddr

		c.JSON(200,User{
			ID : "12345",
			Name: "Tomas",
			Ip: ip,
		})
	})

	

	r.Run("localhost:8080")
}