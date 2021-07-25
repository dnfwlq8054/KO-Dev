package routers

import (
	"github.com/dnfwlq8054/KO-Dev/routers/api/qna"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	//router.Group("/api/qna")

	router.GET("/qna/all", qna.GetAll)

	return router
}
