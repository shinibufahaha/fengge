package handlers

import (
    "github.com/gin-gonic/gin"
)

func DefaultHandler(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"title": "Application Analysis",
	})
}