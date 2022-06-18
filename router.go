package assignment001

import (
	"github.com/gin-gonic/gin"
	"github.com/snoveiry/assignment001/logger"
	routerv1 "github.com/snoveiry/assignment001/v1/router"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (a *Assignment001) router() {
	r := gin.New()

	v1r := routerv1.New()
	v1r.Group = r.Group("/")
	v1r.Assignment001 = a.Config
	v1r.Create()

	echoGroup := r.Group("/")
	echoGroup.GET("/echo", echo)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	logger.FatalError(nil, r.Run(a.Config.Port))
}
