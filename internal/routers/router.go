package routers

import "github.com/gin-gonic/gin"

//NewRouter 文件路由
func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/tags")
		apiv1.GET("/tags")
		apiv1.DELETE("/tags/:id")
		apiv1.PUT("/tags/:id")

	}
	return r
}
