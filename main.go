package main

import (
	"auth_server/models"
	"auth_server/routes"
	"auth_server/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	confFileName = "conf.yml"
)

func main() {
	var conf = utils.LoadConfiguration(confFileName)

	db, err := gorm.Open(sqlite.Open(conf.DbFileName))
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("Connected to database")
	}

	server := gin.Default()
	server.GET("/ping", routes.Ping)
	server.POST("/register", func(c *gin.Context) {
		c.JSON(routes.RegisterNewUser(c, db))
	})

	server.POST("/unregister", func(c *gin.Context) {
		var allParams = routes.CheckPostParameters([]string{"email", "password"}, c)
		if !allParams {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Missing parameters"})
		} else {
			if routes.UnregisterUser(c, db) {
				c.JSON(http.StatusOK, gin.H{"message": "User unregistered"})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"message": "User not found"})
			}
		}
	})

	// Migrate the schema and create the table
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Password{})
	db.AutoMigrate(&models.ApiKey{})
	db.AutoMigrate(&models.RequstLog{})

	// server.Run(":8080")

}
