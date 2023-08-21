package main

import (
	"github.com/CrabStash/crab-stash/api/internal/auth"
	"github.com/CrabStash/crab-stash/api/internal/user"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	authSvc := *auth.RegisterRoutes(r)
	_ = user.RegisterRoutes(r, &authSvc)

	r.Run(":8080")
}
