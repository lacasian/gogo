package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *API) setRoutes() {
	a.engine.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "It works!")
	})
}
