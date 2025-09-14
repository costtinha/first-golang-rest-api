package http

import (
	"net/http"

	"github.com/costtinha/first-golang-rest-api/internal/config"
	"github.com/costtinha/first-golang-rest-api/internal/http/middleware"
	"github.com/costtinha/first-golang-rest-api/internal/logger"
	"github.com/costtinha/first-golang-rest-api/internal/user"
	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine
	cfg    *config.Config
	log    *logger.Logger
}

func NewRouter(cfg *config.Config, log *logger.Logger) *Router {
	if cfg.AppEnv == "dev" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(gin.Logger(), middleware.Recover(), middleware.RequestID())

	return &Router{engine: r, cfg: cfg, log: log}

}

func (r *Router) RegisterHealth() {
	r.engine.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
}

func (r *Router) RegisterUserRouter(h *user.Handler) {
	v1 := r.engine.Group("/v1")
	users := v1.Group("/users")
	{
		users.POST("", h.Create)
		users.GET("/:id", h.GetById)
		users.GET("", h.List)
		users.PUT("/:id", h.Update)
		users.DELETE("/:id", h.Delete)

	}
}

func (r *Router) Run(addr string) error {
	return r.engine.Run(addr)
}
