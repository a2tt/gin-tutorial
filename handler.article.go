package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	//// Define the route for the index page and display the index.html template
	//// To start with, we'll use an inline route handler. Later on, we'll create
	//// standalone functions that will be used as route handlers.
	//// Call the HTML method of the Context to render a template
	//c.HTML(
	//	// Set the HTTP status to 200 (OK)
	//	http.StatusOK,
	//	// Use the index.html template
	//	"index.html",
	//	// Pass the data that the page uses
	//	gin.H{
	//		"title":   "Home Page",
	//		"payload": articles,
	//	},
	//)

	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title":   "Home Page",
		"payload": articles,
	}, "index.html")
}

func getArticle(c *gin.Context) {
	// Check if the article ID is valid
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// Check id the article exists
		if article, err := getArticleByID(articleID); err == nil {
			//// Call teh HTML method of the Context to render a tmeplate
			//c.HTML(
			//	// Set the HTTP status to 200 (OK)
			//	http.StatusOK,
			//	// Use the article.html template
			//	"article.html",
			//	// Pass the data that the page uses
			//	gin.H{
			//		"title":   article.Title,
			//		"payload": article,
			//	},
			//)

			// Call the render function with the name of the template to render
			render(c, gin.H{
				"title":   article.Title,
				"payload": article,
			}, "article.html")
		} else {
			// If the article is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		// If an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}
