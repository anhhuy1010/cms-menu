package routes

import (
	"net/http"

	"github.com/anhhuy1010/cms-menu/controllers"

	docs "github.com/anhhuy1010/cms-menu/docs"
	"github.com/anhhuy1010/cms-menu/middleware"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RouteInit(engine *gin.Engine) {
	productCtr := new(controllers.ProductController)

	engine.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Auth Service API")
	})
	engine.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	engine.Use(middleware.Recovery())
	docs.SwaggerInfo.BasePath = "/v1"
	apiV1 := engine.Group("/v1")

	//apiV1.Use(middleware.ValidateHeader())
	// apiV1.Use(middleware.VerifyAuth())
	apiV1.Use(middleware.RequestLog())
	{
		apiV1.POST("/menu-list", productCtr.Create)
		apiV1.GET("/menu-list", productCtr.List)
		apiV1.GET("/menu-list/:uuid", productCtr.Detail)
		apiV1.PUT("/menu-list/:uuid", productCtr.Update)
		apiV1.DELETE("/menu-list/:uuid", productCtr.Delete)
		apiV1.POST("/menu-detail", productCtr.CreateDetail)
		apiV1.PUT("/menu-detail/:uuid", productCtr.UpdateDetail)
	}
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

}
