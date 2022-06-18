package twitter

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	error1 "github.com/snoveiry/assignment001/error"

	authservice "github.com/snoveiry/assignment001/v1/services/auth"

	"github.com/gin-gonic/gin"

	twittmodel "github.com/snoveiry/assignment001/v1/models/twitter"
)

// GetTwitt handles getting a special twitt from a special user
// @Summary get twitt
// @Description Handles getting a special twitt
// @Tags V1 Twitter
// @Param id path string true "a valid twitt id value"
// @Success 200
// @Security Bearer
// @Router /v1.0/twitter/{id} [get]
func (c *Controller) GetTwitt(ctx *gin.Context) {
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

	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		error1.JSON(ctx, http.StatusNotFound, &error1.E{
			Type:        "END POINT ERROR",
			Description: "Some description goes here.",
		})
		return
	}

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		error1.JSON(ctx, http.StatusInternalServerError, &error1.E{
			Type:        "SERVER ERROR",
			Description: "Calling from thirdparty has error.",
		})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		error1.JSON(ctx, http.StatusBadRequest, &error1.E{
			Type:        "READ RESPONSE ERROR",
			Description: "Could not read service response.",
		})
		return
	}

	var response twittmodel.GetTwittResponse
	if err := json.Unmarshal(body, &response); err != nil {
		error1.JSON(ctx, http.StatusInternalServerError, &error1.E{
			Type:        "RESPONSE UNMARSHAL ERROR",
			Description: "Could not unmarshal twitter service result.",
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
