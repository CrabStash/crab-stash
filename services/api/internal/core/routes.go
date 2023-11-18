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
		Client:    InitServiceClient(),
		Warehouse: warehouseSvc.Client,
	}

	routes := r.Group("core")
	routes.Use(a.AuthRequired)
	schemas := routes.Group("schemas")
	category := routes.Group("category")
	{
		schemas.GET("/category", svc.NewCategorySchema)
		schemas.GET("/field", svc.NewFieldSchema)
		category.GET("/inheritance", svc.FieldsInheritance)
		category.GET("/category", svc.GetCategorySchema)
	}

	return svc
}

func (svc *ServiceClient) NewCategorySchema(ctx *gin.Context) {
	routes.NewCategorySchema(ctx, svc.Client)
}

func (svc *ServiceClient) NewFieldSchema(ctx *gin.Context) {
	routes.NewFieldSchema(ctx, svc.Client)
}

func (svc *ServiceClient) GetCategorySchema(ctx *gin.Context) {
	code, err := warehouse.PermissionHandler(0, svc.Warehouse, ctx)
	if err != nil {
		ctx.JSON(code, gin.H{"status": code, "response": gin.H{"error": err.Error()}})
		return
	}

	code, err = CoreMiddleware(svc.Client, ctx, "categories_to_warehouses")
	if err != nil {
		ctx.JSON(code, gin.H{"status": code, "response": gin.H{"error": err.Error()}})
		return
	}
	routes.GetCategorySchema(ctx, svc.Client)
}

func (svc *ServiceClient) FieldsInheritance(ctx *gin.Context) {
	code, err := warehouse.PermissionHandler(2, svc.Warehouse, ctx)
	if err != nil {
		ctx.JSON(code, gin.H{"status": code, "response": gin.H{"error": err.Error()}})
		return
	}

	code, err = CoreMiddleware(svc.Client, ctx, "categories_to_warehouses")
	if err != nil {
		ctx.JSON(code, gin.H{"status": code, "response": gin.H{"error": err.Error()}})
		return
	}
	routes.FieldsInheritance(ctx, svc.Client)
}
