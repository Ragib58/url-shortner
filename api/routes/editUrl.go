package routes

import (
	"net/http"
	"time"
	"url-shortner/api/database"
	"url-shortner/api/models"

	"github.com/gin-gonic/gin"
)

func EditURL(c *gin.Context) {
	shortID := c.Param("shortID")
	var body models.Request
	if err := c.ShouldBind(&body); err != nil {
		c.JSON((http.StatusBadRequest), gin.H{"error": "Can't Parse JSON"})

	}
	r := database.CreateClient(0)
	defer r.Close()

	// Check if the shortID exists
	val, err := r.Get(database.Ctx, shortID).Result()
	if err != nil || val == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data not found for given shortID"})
		return
	}
	//Update the content of the URL, expiry time with shortID
	err = r.Set(database.Ctx, shortID, body.URL, body.Expiry*3600*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "URL updated successfully"})
}
