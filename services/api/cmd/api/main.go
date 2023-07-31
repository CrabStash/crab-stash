package main

import (
	"github.com/CrabStash/crab-stash/api/internal/auth"
	"github.com/CrabStash/crab-stash/api/internal/warehouse"
	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func main() {
	valid.SetFieldsRequiredByDefault(true)
	r := gin.Default()
	authSVC := *auth.RegisterRoutes(r)

	warehouse.RegisterRoutes(r, &authSVC)

	r.Run(":8080")
}
