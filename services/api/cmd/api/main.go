package main

import (
	"os"
	"strings"

	"github.com/CrabStash/crab-stash/api/internal/auth"
	"github.com/CrabStash/crab-stash/api/internal/core"
	"github.com/CrabStash/crab-stash/api/internal/user"
	"github.com/CrabStash/crab-stash/api/internal/utils"
	"github.com/CrabStash/crab-stash/api/internal/warehouse"
	valid "github.com/asaskevich/govalidator"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	valid.SetFieldsRequiredByDefault(true)
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Split(os.Getenv("ALLOWED_ORIGINS"), ","),
		AllowHeaders:     []string{"Content-Type", "X-XSRF-TOKEN", "accept", "origin", "X-Requested-With", "Authorization", "Accept-Encoding", "Content-Length", "Cache-Control"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowCredentials: true,
	}))

	r.MaxMultipartMemory = 8 << 20
	utils := utils.InitBucket()
	authSvc := *auth.RegisterRoutes(r)
	_ = user.RegisterRoutes(r, &authSvc, &utils)
	warehouseSvc := warehouse.RegisterRoutes(r, &authSvc, &utils)
	_ = core.RegisterRoutes(r, &authSvc, warehouseSvc)

	r.Run(":8080")
}
