package core

import (
	"github.com/CrabStash/crab-stash/api/internal/auth"
	"github.com/CrabStash/crab-stash/api/internal/core/routes"
	"github.com/CrabStash/crab-stash/api/internal/warehouse"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, authSvc *auth.ServiceClient, warehouseSvc *warehouse.ServiceClient) *ServiceClient {
	a := auth.InitAuthMiddleware(authSvc)
	svc := &ServiceClient{
		Client: InitServiceClient(),
	}

	routes := r.Group("core")
	schemas := routes.Group("schemas")
	routes.Use(a.AuthRequired)
	{
		schemas.GET("/category", svc.GetCategorySchema)
		schemas.GET("/field", svc.GetFieldSchema)
	}

	return svc
}

func (svc *ServiceClient) GetCategorySchema(ctx *gin.Context) {
	routes.GetCategorySchema(ctx, svc.Client)
}

func (svc *ServiceClient) GetFieldSchema(ctx *gin.Context) {
	routes.GetFieldSchema(ctx, svc.Client)
}
