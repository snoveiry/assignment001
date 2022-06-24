package twitter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	error1 "github.com/snoveiry/assignment001/error"
	authservice "github.com/snoveiry/assignment001/v1/services/auth"
	thirdparty "github.com/snoveiry/assignment001/v1/services/thirdparty"
)

// GetStream handles getting a real time filter stream tweets
// @Summary get filter stream tweets
// @Description Handles getting filtered stream tweets of a special user with special rules
// @Tags V1 Twitter
// @Success 200
// @Security Bearer
// @Router /v1.0/tweet/get/stream [get]
func (c *Controller) GetStream(ctx *gin.Context) {

	url := "https://api.twitter.com/2/tweets/search/stream"

	// Create a Bearer string by appending string access token
	var bearer = "Bearer "
	auth := authservice.New(c.Assignment001)
	token := auth.UniversalAuthenticator(ctx)
	if token != nil {
		bearer += *token
	} else {
		error1.JSON(ctx, http.StatusUnauthorized, &error1.E{
			Type:        "CLIENT UNRECOGNIZED",
			Description: "Authorization error.",
		})
		return
	}

	thpService := thirdparty.New(c.Assignment001)
	thpService.GetStreamService(ctx, url, bearer)
}
