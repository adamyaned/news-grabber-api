package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"news-grabber-api/news"
)

func indexHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello world")
}

func collectHandler(c *gin.Context) {
	category := c.Param("category")

	news.Collect(category)

	c.String(http.StatusOK, "Search initialized")
}

func resultHandler(c *gin.Context) {
	category := c.Param("category")

	topics := news.Result(category)

	c.JSON(http.StatusOK, topics)
}
