package routers

import (
	"github.com/dnfwlq8054/KO-Dev/routers/api/discussion"
	"github.com/dnfwlq8054/KO-Dev/routers/api/event"
	"github.com/dnfwlq8054/KO-Dev/routers/api/qna"
	"github.com/dnfwlq8054/KO-Dev/routers/api/toyproject"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	//	router.Group("/api/qna")

	router.GET("/qna/all", qna.GetAll)
	router.GET("/event/all", event.GetAll)
	router.GET("/discussion/all", discussion.GetAll)
	router.GET("/toyproject/all", toyproject.GetAll)

	router.GET("/qna/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return router
}
