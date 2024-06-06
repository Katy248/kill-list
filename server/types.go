package server

import "github.com/gin-gonic/gin"

type Options struct {
	Name string
}

func DefaultOptions() Options {
	return Options{
		Name: "default",
	}
}

type Server struct {
	Options
	r *gin.Engine
}

type OptionsFunc func(Options)

func NewServer(opts ...OptionsFunc) Server {
	options := DefaultOptions()
	for _, fn := range opts {
		fn(options)
	}

	routes := gin.Default()
	setupRoutes(routes)
	return Server{
		Options: options,
		r:       routes,
	}
}

func setupRoutes(routes *gin.Engine) {
	routes.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

func (server *Server) Listen(addr string) error {
	return server.r.Run(addr)
}
