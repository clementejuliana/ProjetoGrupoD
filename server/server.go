package server

import (
	"log"

	"github.com/clementejuliana/grupoDprojetoGO/server/routes"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server{
	return Server{
		port: "5000",
		server: gin.Default(),
	}
}

func(s *Server) Run(){
	router := routes.Config(s.server)

	log.Print("server is running at port:", s.port)
	log.Fatal(router.Run(":" + s.port))
}