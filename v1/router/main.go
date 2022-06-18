// Package router defines all routes for the v1 API
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/snoveiry/assignment001/config"
	"github.com/snoveiry/assignment001/v1/controllers/twitter"
)

type Router struct {
	Group         *gin.RouterGroup
	Assignment001 config.Assignment001
}

func New() Router {
	return Router{}
}

func (r *Router) Create() {
	twitterV1 := twitter.NewController()
	mainGroup := r.Group.Group("/")
	{
		mainGroup.GET("/v1.0/twitter/:id", twitterV1.GetTwitt)
	}
}
