package router

import (
	"github.com/gin-gonic/gin"

	"github.com/bragfoo/saman/backend/api/controller"
	"github.com/bragfoo/saman/backend/api/global"
)

func Server(router *gin.RouterGroup, g *global.G) {
	conf := g.Conf

	v1 := router.Group("/" + conf.Api.Version)
	// GET
	//v1.GET("/platformType",controller.GetPlatType(g))
	v1.GET("/platformFans",controller.GetPlatFans(g))
	//v1.GET("/platformFans",controller.GetPlatFansTemp(g))
	v1.GET("/events",controller.GetEvents(g))
	v1.GET("/mobileData",controller.GetMobileData(g))
	v1.GET("/appUGC",controller.GetUGC(g))
	v1.GET("/videoPlayAmount",controller.GetVideoPlayAmount(g))
	v1.GET("/sil",controller.GetPing(g))

	// POST
	v1.POST("/platformFans",controller.PostPlatFans(g))
	v1.POST("/events",controller.PostEvents(g))
	v1.POST("/mobileData",controller.PostMobileData(g))
	v1.POST("/appUGC",controller.PostUGC(g))
	v1.POST("/videoPlayAmount",controller.PostVideoPlayAmount(g))

	//test
	test:=router.Group("/test")
	test.GET("/platformType",controller.GetPlatType(g))
}
