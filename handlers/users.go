package handlers

import (
	"net/http"
	"rest-api/models"
	"rest-api/utils"

	"github.com/gin-gonic/gin"
)


func SignupUser(c *gin.Context) {
var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	
	err := user.Save()
	if err != nil {
		print(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user"})
		return
	}	
	c.JSON(http.StatusCreated, gin.H{"message": "User saved successfully"})
}
 

func LoginUser(c *gin.Context) {
	var user models.User
	err :=c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": "Required fields were not given"})
		return
	}
	// serach for the user so that we can compare
	err =user.ValidateCredentials()
	if err != nil {
		// if you want it to be more realistic, it should just use err.Error() instead of this very generic (and wrong in some cases) message
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}
	token, err := utils.GenerateJWT(user.Email, user.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate token"})
		return
	}
	// set the cookie
	c.SetCookie("rest-event-cookie", token, 60*60*2, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"token": token, "message": "Login successful"})
}