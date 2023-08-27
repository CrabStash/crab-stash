package main

import (
	"github.com/CrabStash/crab-stash/api/internal/auth"
	"github.com/CrabStash/crab-stash/api/internal/user"
	"github.com/CrabStash/crab-stash/api/internal/warehouse"
	valid "github.com/asaskevich/govalidator"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	valid.SetFieldsRequiredByDefault(true)
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowHeaders:     []string{"Origin, Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	authSvc := *auth.RegisterRoutes(r)
	_ = user.RegisterRoutes(r, &authSvc)
	_ = warehouse.RegisterRoutes(r, &authSvc)

	r.Run(":8080")
}
