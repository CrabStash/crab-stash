package main

import (
	pb "github.com/CrabStash/crab-stash/auth/proto"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine, c pb.AuthServiceClient) {
	v1 := r.Group("/v1")
	{
		v1.POST("/login", LoginHandler(c))
		v1.POST("/register", RegisterHandler(c))
	}
	return
}
