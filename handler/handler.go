package handler

import (
	"github.com/gin-gonic/gin"
	. "main/common"
	"net/http"
)

/*
身份验证
*/
func AuthHandler(context *gin.Context) {
	context.Next()
}

/*
404处理
*/
func NotFoundHandler(context *gin.Context) {
	context.JSON(http.StatusNotFound, Result{
		Message: NotFound,
	})
}

/*
500处理
*/
func ServerErrorHandler(context *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			context.JSON(http.StatusInternalServerError, Result{
				Data:    r,
				Message: ServerError,
			})
		}
	}()
	context.Next()
}
