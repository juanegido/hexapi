package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	mooc "github.com/juanegido/hexapi/internal"
	"github.com/juanegido/hexapi/internal/platform/server/handler/courses"
	"github.com/juanegido/hexapi/internal/platform/server/handler/health"
	"log"
)

type Server struct {
	engine   *gin.Engine
	httpAddr string

	// dependencies
	courseRepository mooc.CourseRepository
}

func New(host string, port uint, courseRepository mooc.CourseRepository) Server {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("#{host}:#{port}"),

		courseRepository: courseRepository,
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Starting server on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHealth())
	s.engine.POST("/courses", courses.CreateHandler(s.courseRepository))
}
