package bootstrap

import "github.com/juanegido/hexapi/internal/platform/server"

const (
	HOST = "localhost"
	PORT = 8080
)

func Run() error {
	srv := server.New(HOST, PORT)
	return srv.Run()
}
