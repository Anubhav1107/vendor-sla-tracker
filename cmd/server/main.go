package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Event struct {
	Vendor    string `json:"vendor"`
	ErrorCode string `json:"error_code"`
	TimeStamp string `json:"timestamp"`
}

func main() {
	fmt.Println("Vendor SLA Tracker started!")
	r := gin.Default()

	r.POST("/events", func(c *gin.Context) {
		var e Event

		if err := c.BindJSON(&e); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}

		fmt.Printf("Received Event — Vendor: %s, Error: %s, Time: %s\n",
			e.Vendor, e.ErrorCode, e.TimeStamp)

		c.JSON(http.StatusOK, gin.H{
			"message": "Event logged successfully",
		})
	})
	fmt.Println("SLA Tracker API running on :8080")
	r.Run(":8080")
}
