package handler

import (
	"net/http"
	"strconv"

	"github.com/c-major/article_manager/common"
	"github.com/c-major/article_manager/model"
	"github.com/gin-gonic/gin"
)

// ShowIndexPage .
func ShowIndexPage(c *gin.Context) {
	articleList := model.GetAllArticles()

	common.Render(c,
		gin.H{
			"title":   "Home Page",
			"payload": articleList,
		},
		"index.html")
}

// GetArticle .
func GetArticle(c *gin.Context) {
	articleID, err := strconv.ParseUint(c.Param("article_id"), 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	article, err := model.GetArticleByID(articleID)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}

	common.Render(
		c,
		gin.H{
			"title":   article.Title,
			"payload": article,
		},
		"article.index")
}
