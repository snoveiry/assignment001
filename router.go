package assignment001

import (
	"github.com/gin-gonic/gin"
	"github.com/snoveiry/assignment001/logger"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	routerv1 "github.com/snoveiry/assignment001/v1/router"
)

func (m *Assignment001) router() {
	r := gin.New()

	r.Use(logger.Recovery())
	r.Use(logger.Default())
	r.Use(logger.JSON())



	v2r := routerv1.New()
	v2r.Group = r.Group("/")
	v2r.Monolith = m.Config
	v2r.Create()

	echoGroup := r.Group("/")
	echoGroup.Use(m.Config.AuthService.Universal)
	//echoGroup.GET("/echo", echo)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	logger.FatalError(nil, r.Run(m.Config.Port))
}
