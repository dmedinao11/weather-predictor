package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Ping struct{}

func (p *Ping) Pong() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	}
}
