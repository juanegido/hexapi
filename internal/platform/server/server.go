package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/juanegido/hexapi/internal/platform/server/handler/courses"
	"github.com/juanegido/hexapi/internal/platform/server/handler/health"
	"github.com/juanegido/hexapi/kit/bus"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	// deps
	bus bus.Bus
}

func New(host string, port uint, bus bus.Bus) Server {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),

		bus: bus,
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
	s.engine.POST("/courses", courses.CreateHandler(s.bus))
	s.engine.GET("/courses", courses.GetHandler(s.bus))
}
