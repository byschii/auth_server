package routes

import (
	"auth_server/models"
	"auth_server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// check ehat evey parameter in the list is present in the request
func CheckPostParameters(parameters []string, c *gin.Context) bool {
	for _, param := range parameters {
		if c.PostForm(param) == "" {
			return false
		}
	}
	return true
}

// Gets data from POST request, creates and save a USER
func RegisterNewUser(c *gin.Context, db *gorm.DB) {
	// gets parameters from the POST request
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")

	// salt and hash the password
	var pwdHash, pwdSalt = utils.HashAndSaltPassword(password)
	// create a new user without apikey
	var u = models.User{
		UserName: username,
		Email:    email,
		Verified: false,
		Password: models.Password{
			HashedPassword: pwdHash,
			Salt:           pwdSalt,
			Resettable:     models.Resettable{},
		},
		ApiKey: models.ApiKey{
			Key:        "",
			Resettable: models.Resettable{},
		},
	}

	db.Create(u)
}

func UnregisterUser(c *gin.Context, db *gorm.DB) bool {
	// gets parameters from the POST request
	email := c.PostForm("email")
	plainPassword := c.PostForm("password")

	// get the user from the database
	var user = models.User{}
	db.Take(user, "email = ?", email)

	// check if password matches
	var tx = db.Begin()
	if utils.DoPasswordsMatch(plainPassword, user.Password.HashedPassword, user.Password.Salt) {
		tx = db.Delete(user)
	}
	defer tx.Commit()

	return tx.RowsAffected > 0

}
