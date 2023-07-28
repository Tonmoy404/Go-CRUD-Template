package rest

import (
	"fmt"
	"net/http"

	"github.com/Tonmoy404/project/config"
	"github.com/Tonmoy404/project/service"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router    *gin.Engine
	appConfig *config.Application
	svc       service.Service
	salt      *config.Salt
}

func NewServer(appConfig *config.Application, svc service.Service, salt *config.Salt) (*Server, error) {
	server := &Server{
		appConfig: appConfig,
		svc:       svc,
		salt:      salt,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	server.router = gin.Default()
	server.router.GET("/api/test", server.test)

	server.router.POST("/api/signup", server.CreateUser)
	server.router.GET("/api/user/:id", server.GetUser)
	server.router.PATCH("/api/user/:id", server.UpdateUser)
	server.router.DELETE("/apip/user/:id", server.DeleteUser)

}

func (server *Server) Start() error {
	return server.router.Run(fmt.Sprintf("%s:%s", server.appConfig.Host, server.appConfig.Port))
}

func (server *Server) test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "testing")
}
