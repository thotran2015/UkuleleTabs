package main

import (
	"UkuleleTabs/tabGenerator"
	"github.com/gin-gonic/gin"
	"net/http"
)

func generateTabHandler(c *gin.Context) {
	chordName := c.Query("chord")
	result := tabGenerator.CreateTab(chordName)
	c.JSON(http.StatusOK, gin.H{"generatedTab": result})
}

func main() {
	router := gin.Default()
	router.GET("ukulele/generate-tabs", generateTabHandler)
	_ = router.Run(":8080")
}
