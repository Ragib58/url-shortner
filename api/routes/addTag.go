package routes

import (
	"encoding/json"
	"net/http"
	"url-shortner/api/database"

	"github.com/gin-gonic/gin"
)

type TagRequest struct {
	ShortID string `json:"shortID"`
	Tag     string `json:"tag" `
}

func AddTag(c *gin.Context) {
	var tagRequest TagRequest
	if err := c.ShouldBindJSON(&tagRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	shortID := tagRequest.ShortID
	tag := tagRequest.Tag

	r := database.CreateClient(0)
	defer r.Close()

	val, err := r.Get(database.Ctx, shortID).Result()
	if err != nil || val == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data not found for given shortID"})
		return
	}

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		//if the data is not in JSON format, we can assume it's a simple string
		data = make(map[string]interface{})
		data["data"] = val
	}

	//check if the tag already exists and it's a slice of strings
	var tags []string
	if existingTags, ok := data["tags"].([]interface{}); ok {
		for _, t := range existingTags {
			if tagStr, ok := t.(string); ok {
				tags = append(tags, tagStr)
			}
		}
	}

	//check for duplicate tag
	for _, existingTag := range tags {
		if existingTag == tag {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Tag already exists"})
			return
		}
	}

	// Add the new tag
	tags = append(tags, tag)
	data["tags"] = tags

	//Marshal the updated data back to JSON
	updatedData, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Marshal updated data"})
		return
	}

	err = r.Set(database.Ctx, shortID, updatedData, 0).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the database"})
		return
	}

	//Response with the updated data
	c.JSON(http.StatusOK, data)
}
