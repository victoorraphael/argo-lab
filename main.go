package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	app.POST("/", example)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: app,
	}

	log.Fatal(srv.ListenAndServe())
}

func example(c *gin.Context) {
	req := struct {
		Name string `json:"name" binding:"required"`
	}{}
	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, req)
}
