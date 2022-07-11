package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	// HTTP format
	// c.HTML(
	// 	http.StatusOK,
	// 	"index.html",
	// 	gin.H{
	// 		"title":   "Home Page",
	// 		"payload": articles,
	// 	},
	// )

	// render the context
	render(c, gin.H{
		"title":   "Home Page",
		"payload": articles}, "index.html")
}

func getArticle(c *gin.Context) {
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		if article, err := getArticleByID(articleID); err == nil {
			// if the corresponding article is found then call a HTML response to render the template
			c.HTML(
				http.StatusOK, // http status code
				"article.html",
				gin.H{
					"title":   article.ID,
					"payload": article,
				},
			)
		} else {
			// abort if the article is not found
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		// abort if the article is not found
		c.AbortWithStatus(http.StatusNotFound)
	}
}

// render the response content as the required format
func render(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, data)
	}
}
