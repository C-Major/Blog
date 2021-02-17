package handler

import (
	"net/http"
	"strconv"

	"github.com/c-major/blog/common"
	"github.com/c-major/blog/db"
	"github.com/gin-gonic/gin"
)

// ShowIndexPage .
func ShowIndexPage(c *gin.Context) {
	common.TextLog.WithContext(c)

	articleList, err := db.GetAllArticles()
	if err != nil {
		common.TextLog.Error("[ShowIndexPage] failed to get all articles from db")
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	common.Render(c,
		gin.H{
			"title":   "Home Page",
			"payload": articleList,
		},
		"index.html")
}

// GetArticle .
func GetArticle(c *gin.Context) {
	common.TextLog.WithContext(c)

	articleID, err := strconv.ParseUint(c.Param("article_id"), 10, 64)
	if err != nil {
		common.TextLog.Error("[GetArticle] failed to parse url param")
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	article, err := db.GetArticleByID(articleID)
	if err != nil {
		common.TextLog.Error("[GetArticle] failed to get article from db")
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	if article == nil {
		common.Render(
			c,
			gin.H{
				"title": "article not found",
			},
			"article.html",
		)
		return
	}

	common.Render(
		c,
		gin.H{
			"title":   article.Title,
			"payload": article,
		},
		"article.html")
}
