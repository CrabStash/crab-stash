package main

import (
	"github.com/CrabStash/crab-stash/api/internal/auth"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	_ = *auth.RegisterRoutes(r)

	r.Run(":8080")
}
