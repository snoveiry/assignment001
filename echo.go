package assignment001

import (
	"github.com/gin-gonic/gin"
)

type echoResponse struct {
	Message string `json:"message"`
}

// echo should always be accessible and serves as an availability check of the API
// @Summary Uptime endpoint
// @Tags Global
// @Description Sending a get request to the echo endpoint should always return a response
// @Produce  json
// @Success 200 {object} echoResponse
// @Router /echo [get]
func echo(ctx *gin.Context) {
	ctx.JSON(200, echoResponse{
		Message: "Welcome to twitter service.",
	})
}
