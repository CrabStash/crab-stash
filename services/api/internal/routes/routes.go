package routes

import (
	pb "github.com/CrabStash/crab-stash-protofiles/auth/proto"
	"github.com/CrabStash/crab-stash/api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine, c pb.AuthServiceClient) {
	v1 := r.Group("/v1")
	{
		v1.POST("/login", handlers.LoginHandler(c))
		v1.POST("/register", handlers.RegisterHandler(c))
	}
	return
}
