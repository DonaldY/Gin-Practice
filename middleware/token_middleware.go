package middleware

import (
	"Gin-Practice/utils"
	"github.com/gin-gonic/gin"
	)

func AuthTokenMiddleware() gin.HandlerFunc {

	return func(context *gin.Context) {

		accessToken := getAccessToken(context)

		if stringUtils.IsEmpty(accessToken) {

		}
	}
}

func getAccessToken(context *gin.Context) string {

	return ""
}
