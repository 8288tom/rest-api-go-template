package utils

import "github.com/gin-gonic/gin"

func GetToken(context *gin.Context) string {
	token := context.Request.Header.Get("Authorization")
	return token
}
