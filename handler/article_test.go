package handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/c-major/blog/caller"
	"github.com/c-major/blog/common"
	"github.com/c-major/blog/util"
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

// Test that a GET request to the home page returns the home page with
// the HTTP code 200 for an unauthenticated user
func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := util.GetTestRouter(true)
	r.GET("/", ShowIndexPage)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/", nil)

	util.TestHTTPResp(t, r, req, func(rr *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200
		statusOK := rr.Code == http.StatusOK

		// Test that the page title is "Home Page"
		// You can carry out a lot more detailed tests using libraries that can
		// parse and process HTML pages
		b, err := ioutil.ReadAll(rr.Body)
		pageOK := (err == nil && strings.Index(string(b), "<title>Home Page</title>") > 0)

		return statusOK && pageOK
	})
}

func TestGetArticleUnauthenticated(t *testing.T) {
	r := util.GetTestRouter(true)
	r.GET("/article/view/:article_id", GetArticle)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/article/view/1", nil)

	util.TestHTTPResp(t, r, req, func(rr *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200
		statusOK := rr.Code == http.StatusOK

		// Test that the page title is "Article 1"
		// You can carry out a lot more detailed tests using libraries that can
		// parse and process HTML pages
		b, err := ioutil.ReadAll(rr.Body)
		pageOK := (err == nil && strings.Index(string(b), "<title>test title 1</title>") > 0)

		return statusOK && pageOK
	})
}
