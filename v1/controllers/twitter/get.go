package twitter

import (
	"encoding/json"
	"net/http"
	"strconv"

	error1 "github.com/snoveiry/assignment001/error"

	authservice "github.com/snoveiry/assignment001/v1/services/auth"
	thirdparty "github.com/snoveiry/assignment001/v1/services/thirdparty"

	"github.com/gin-gonic/gin"

	tweetmodel "github.com/snoveiry/assignment001/v1/models/twitter"
)

// GetTweet handles getting a special tweet from a special user
// @Summary get tweet
// @Description Handles getting a special tweet
// @Tags V1 Twitter
// @Param id path string true "a valid tweet id value"
// @Success 200 {object} tweetmodel.GetTweetResponse
// @Security Bearer
// @Router /v1.0/tweet/get/{id} [get]
func (c *Controller) GetTweet(ctx *gin.Context) {
	tid := ctx.Param("id")
	url := "https://api.twitter.com/2/tweets/"

	if _, err := strconv.Atoi(tid); err == nil {
		url += tid
	} else {
		error1.JSON(ctx, http.StatusBadRequest, &error1.E{
			Type:        "PARAMETER ERROR",
			Description: "Input parameter is not valid",
		})
		return
	}

	// Create a Bearer string by appending string access token
	var bearer = "Bearer "
	auth := authservice.New(c.Assignment001)
	token := auth.UniversalAuthenticator(ctx)
	if token != nil && *token != "" && *token == "AAAAAAAAAAAAAAAAAAAAAM5PdwEAAAAAgXwbaN0ExfG7lytY18p8Xvk9fGA%3DhSCNeiaFgahz5FpY3BMQ4RwCXoP8GubTg9C4YLYUU0MzD6d937" {
		bearer += *token
	} else {
		error1.JSON(ctx, http.StatusUnauthorized, &error1.E{
			Type:        "CLIENT UNRECOGNIZED",
			Description: "Authorization error.",
		})
		return
	}

	thpService := thirdparty.New(c.Assignment001)
	res, body := thpService.GetTweetService(ctx, url, bearer)
	if !res {
		return
	}

	var response tweetmodel.GetTweetResponse
	if err := json.Unmarshal(body, &response); err != nil {
		error1.JSON(ctx, http.StatusInternalServerError, &error1.E{
			Type:        "RESPONSE UNMARSHAL ERROR",
			Description: "Could not unmarshal twitter service result.",
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
