package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getWhitelist(c *gin.Context) {
	orgID := c.Param("orgID")

	var whitelist []string
	if orgID == "exampleOrg" {
		whitelist = []string{"1.1.178.95", "192.168.1.1"}
	}

	if len(whitelist) == 0 {
		c.JSON(http.StatusOK, nil)
	} else {
		c.JSON(http.StatusOK, whitelist)
	}
}

func main() {
	router := gin.Default()
	router.GET("/whitelist/:orgID", getWhitelist)
	router.Run(":8080")
}
