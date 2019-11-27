package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (a *API) setRoutes() {
	a.engine.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "It works!")
	})
}
