package controller

import (
	"gin-learn-todo/pkg/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Logger(ctx *gin.Context) *zap.SugaredLogger {
	if l, ok := ctx.Get("logger"); ok {
		return l.(*zap.SugaredLogger)
	}
	return log.Sugar()
}
