package servers

import (
	"go-clean-arch/configs"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Server struct {
	Engine *gin.Engine
	Db     *gorm.DB
	Cfg    *configs.Configs
}

func NewServer(config *configs.Configs, db *gorm.DB) *Server {
	return &Server{
		Engine: gin.Default(),
		Db:     db,
		Cfg:    config,
	}
}

func (s *Server) Run() {
	// ? Mapping Route
	s.MapHandlers()

	log.Printf("%+v", s.Cfg)
	s.Engine.Run(":" + s.Cfg.Gin.Port)

}
