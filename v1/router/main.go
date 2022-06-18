// Package router defines all routes for the v1 API
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/snoveiry/assignment001/config"
)

type Router struct {
	Group         *gin.RouterGroup
	Assignment001 config.Assignment001
}

func New() Router {
	return Router{}
}

func (r *Router) Create() {

}
