package handlers

import (
	"context"
	"net/http"
	"time"
"github.com/pruthvimax/user-service/database"
"github.com/pruthvimax/user-service/models"


	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	var user models.User

	// get json from request
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := database.UserCollection.InsertOne(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
