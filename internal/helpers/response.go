package helpers

import (
	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, code int, data interface{}, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"data":    data,
		"messege": message,
	})
}

func Error(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, gin.H{
		"code":    code,
		"data":    data,
		"messege": message,
	})
}
