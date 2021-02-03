package util

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// GetTestRouter is a helper function to create a router during testing
func GetTestRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("../template/*")
	}

	return r
}

// TestHTTPResp is a helper function to process a request and test its response
func TestHTTPResp(t *testing.T, r *gin.Engine, req *http.Request, f func(rr *httptest.ResponseRecorder) bool) {
	// Create a response recorder
	recorder := httptest.NewRecorder()

	// Create the service and process the above request.
	r.ServeHTTP(recorder, req)

	if !f(recorder) {
		t.Fail()
	}
}
