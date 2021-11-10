package routes

import (
	"auth_server/models"
	"auth_server/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"time":    time.Now,
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
func RegisterNewUser(c *gin.Context, db *gorm.DB, cfg utils.Configuration) (int, map[string]interface{}) {

	if !CheckPostParameters([]string{"username", "password", "email"}, c) {
		return http.StatusBadRequest, gin.H{"message": "Missing parameters"}
	}

	// gets parameters from the POST request
	var email = c.PostForm("email")
	// save new user
	var _, salt = _SaveNewUserInDB(c.PostForm("username"), email, c.PostForm("password"), db)
	// send email to the user for verification
	utils.SendEmailVerficationMail(email, salt, cfg)

	return http.StatusCreated, gin.H{"message": "User created"}
}

// save a new user in the Db
// the new user is not verified yet, and has no api key
// then returns both user id and salt
func _SaveNewUserInDB(username string, email string, password string, db *gorm.DB) (uint, string) {
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
	return u.ID, pwdSalt
}

func UnregisterUser(c *gin.Context, db *gorm.DB, cfg utils.Configuration) bool {
	// gets parameters from the POST request
	email := c.PostForm("email")
	plainPassword := c.PostForm("password")

	// get the user from the database
	var user = models.User{}
	db.Preload("Password").Take(user, "email = ?", email)

	// check if password matches
	var tx = db.Begin()
	if utils.DoPasswordsMatch(plainPassword, user.Password.HashedPassword, user.Password.Salt) {
		tx = db.Delete(user)
	}
	defer tx.Commit()

	return tx.RowsAffected > 0

}
