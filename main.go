package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {

	// Connection Database
	dsn := "root:@tcp(127.0.0.1:3306)/go-db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error Connecting database")
		fmt.Println(db) // TODO remove
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/posts", Posts)
	r.GET("/posts/:id", Show)
	r.POST("/posts", Store)
	r.PUT("/posts/:id", Update)
	r.DELETE("/posts/:id", Delete)
	r.Run(":9090") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
