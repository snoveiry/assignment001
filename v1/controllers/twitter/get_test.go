package twitter

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetTweetNoAuth(t *testing.T) {

	w := httptest.NewRecorder()
	gin.SetMode(gin.ReleaseMode)
	c, _ := gin.CreateTestContext(w)

	c.Params = []gin.Param{
		{
			Key:   "id",
			Value: "20",
		},
	}

	c.Request = httptest.NewRequest("GET", "/v1.0/tweet/get/:id", nil)
	controller := NewController()

	controller.GetTweet(c)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	expectedStatusCode := http.StatusUnauthorized

	if resp.StatusCode != expectedStatusCode {
		t.Errorf("Response code '%d' does not match expected response code '%d'", resp.StatusCode, expectedStatusCode)
		return
	}

	t.Log(string(body))
}

func TestGetTweetNoParam(t *testing.T) {

	w := httptest.NewRecorder()
	gin.SetMode(gin.ReleaseMode)
	c, _ := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest("GET", "/v1.0/tweet/get/:id", nil)
	controller := NewController()

	controller.GetTweet(c)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	expectedStatusCode := http.StatusBadRequest

	if resp.StatusCode != expectedStatusCode {
		t.Errorf("Response code '%d' does not match expected response code '%d'", resp.StatusCode, expectedStatusCode)
		return
	}

	t.Log(string(body))
}
