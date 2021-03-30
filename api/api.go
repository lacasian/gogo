package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lacasian/gogo/db"
	"github.com/sirupsen/logrus"
)

var log = logrus.WithField("module", "api")

type Config struct {
	Port           string
	DevCorsEnabled bool
	DevCorsHost    string
}

type API struct {
	config Config
	engine *gin.Engine
	db     *db.DB
}

func New(config Config, db *db.DB) *API {
	return &API{
		config: config,
		db:     db,
	}
}

func (a *API) Run() {
	a.engine = gin.Default()

	if a.config.DevCorsEnabled {
		a.engine.Use(cors.New(cors.Config{
			AllowOrigins:     []string{a.config.DevCorsHost},
			AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
			AllowHeaders:     []string{"Origin"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
		}))
	}

	a.setRoutes()

	err := a.engine.Run(":" + a.config.Port)
	if err != nil {
		log.Fatal(err)
	}
}

func (a *API) Close() {
}
