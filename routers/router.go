package routers

import (
	"github.com/dnfwlq8054/KO-Dev/routers/api/qna"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	apiv1 := router.Group("/api/qna")

	apiv1.GET("/qna/all", qna.GetAll)

	return router
}
