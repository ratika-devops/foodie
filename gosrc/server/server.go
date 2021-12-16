package server

import (
	"net/http"

	"foodies/gosrc/cache"
	"foodies/gosrc/controller"
	"foodies/gosrc/docs"

	"github.com/gin-gonic/gin"
	"github.com/markbates/pkger"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartServer() {
	cache.InitCache()

	var UIWebPath = "/foodies/ui"
	var APIWebPath = "/foodies/api/v1"
	var SwaggerPath = "/foodies/swagger"

	router := gin.Default()

	router.StaticFS(UIWebPath, pkger.Dir("/webui-docroot"))

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, UIWebPath)
		c.AbortWithStatus(http.StatusTemporaryRedirect)
	})

	router.GET("/foodies", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, UIWebPath)
		c.AbortWithStatus(http.StatusTemporaryRedirect)
	})

	router.GET("/foodies/swagger", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, SwaggerPath+"/index.html")
		c.AbortWithStatus(http.StatusTemporaryRedirect)
	})

	c := controller.NewController()

	v1 := router.Group(APIWebPath)
	{
		searchby := v1.Group("/searchby")
		{
			searchby.GET("truck/:truck", c.SearchByTruck)
			searchby.GET("cousine/:cousine", c.SearchByCousine)
			searchby.GET("address/:address", c.SearchByAddress)

		}
		cache := v1.Group("/cache")
		{
			cache.GET("refresh", c.CacheRefresh)
		}
	}

	docs.SwaggerInfo.BasePath = APIWebPath
	docs.SwaggerInfo.Title = "foodies-api"
	router.GET(SwaggerPath+"/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":" + "8084")
}
