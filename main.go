package main

import (
	"auth_server/models"
	"auth_server/routes"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	server := gin.Default()
	server.GET("/ping", routes.Ping)

	db, err := gorm.Open(sqlite.Open("auth.db"))
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("Connected to database")
	}

	// Migrate the schema and create the table
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.ApiKey{})
	db.AutoMigrate(&models.RequstLog{})

	var user models.User
	db.Preload("RequstLog").Take(&user, 1)

	fmt.Println(user)

	/*
			db.Create(&models.RequstLog{
			UserID: 1,
			Method: "GET",
			Path:   "/api/v1/users",
		})
		----
				db.Create(&models.User{
					UserName:     "teo",
					Email:        "teo_sca@byschii.com",
					PasswordHash: "123456",
					ApiKey: models.ApiKey{
						Key:       "123456789",
						CodeReset: "siohb",
						Resetting: false,
					},
				})
			----
				var k models.ApiKey
				db.Take(&k)

				fmt.Println(k)

	*/

}
