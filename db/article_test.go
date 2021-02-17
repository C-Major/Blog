package db

import (
	"os"
	"testing"

	"github.com/c-major/blog/caller"
	"github.com/c-major/blog/common"
	"github.com/c-major/blog/model"
	"github.com/stretchr/testify/assert"
)

func init() {
	common.InitLog()

	rootDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	os.Setenv("IS_TEST_ENV", "1")
	config, err := common.GetConfig(rootDir, "../", "conf")
	if err != nil {
		common.TextLog.Error("failed to initialize config")
	}

	err = caller.InitCaller(config)
	if err != nil {
		common.TextLog.Error("failed to initialize caller")
	}
}

func TestCreateArticleByEntity(t *testing.T) {
	article := model.BlogArticle{
		Title:    "test title",
		Content:  "test content",
		AuthorID: 1,
		Status:   0,
	}

	err := CreateArticleByEntity(&article)
	assert.Nil(t, err)
}

func TestUpdateArticleByFilter(t *testing.T) {
	var (
		id       uint64 = 1
		newTitle string = "updated title"
	)

	filter := ArticleUpdateFilter{
		Title: &newTitle,
	}
	err := UpdateArticleByFilter(id, &filter)
	assert.Nil(t, err)
}

func TestGetArticleByID(t *testing.T) {
	var id uint64 = 1

	article, err := GetArticleByID(id)
	assert.Nil(t, err)
	t.Logf("article: %v", article)
}

func TestGetArticleByAuthorID(t *testing.T) {
	var authorID uint64 = 1

	articleList, err := GetArticlesByAuthorID(authorID)
	assert.Nil(t, err)
	for _, article := range articleList {
		t.Logf("article: %v", article)
	}
}

func TestGetArticleByFilter(t *testing.T) {
	var (
		title  string = "test"
		status int8   = 0
	)

	filter := ArticleQueryFilter{
		Title:  &title,
		Status: &status,
	}

	articleList, err := GetArticlesByFilter(&filter)
	assert.Nil(t, err)
	for _, article := range articleList {
		t.Logf("article: %v", article)
	}
}

func TestGetAllArticles(t *testing.T) {
	articleList, err := GetAllArticles()
	assert.Nil(t, err)
	for _, article := range articleList {
		t.Logf("article: %v", article)
	}
}
