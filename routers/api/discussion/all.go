package qna

import "github.com/gin-gonic/gin"

func GetAll(c *gin.Context) {
	c.String(200, "discussion")
}
