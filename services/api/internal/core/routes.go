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
	entity := routes.Group("entity")
	field := routes.Group("field")
	{
		schemas.GET("/category", svc.NewCategorySchema)
		schemas.GET("/field", svc.NewFieldSchema)
		schemas.GET("/:id/warehouse/:warehouseID", svc.GetCategorySchema)
		schemas.GET("/:id/warehouse/:warehouseID/inheritance", svc.FieldsInheritance)
		// categories
		category.GET("/:id/warehouse/:warehouseID", svc.GetCategoryData)
		category.POST("/:warehouseID", svc.CreateCategory)
		category.PATCH("/:id/warehouse/:warehouseID", svc.EditCategory)
		category.DELETE("/:id/warehouse/:warehouseID", svc.DeleteCategory)
		category.GET("/warehouse/:warehouseID", svc.ListCategories)
		// fields
		field.GET("/:id/warehouse/:warehouseID", svc.GetFieldData)
		field.POST("/:warehouseID", svc.CreateField)
		field.PATCH("/:id/warehouse/:warehouseID", svc.EditField)
		field.DELETE("/:id/warehouse/:warehouseID", svc.DeleteField)
		field.GET("/warehouse/:warehouseID", svc.ListFields)
		// entities
		entity.GET("/:id/category/:categoryID/warehouse/:warehouseID", svc.GetEntityData)
		entity.POST("/:categoryID/warehouse/:warehouseID", svc.CreateEntity)
		entity.PATCH("/:id/category/:categoryID/warehouse/:warehouseID", svc.EditEntity)
		entity.DELETE("/:id/category/:categoryID/warehouse/:warehouseID", svc.DeleteEntity)
		entity.GET("/warehouse/:warehouseID", svc.ListFields)
	}

	return svc
}

// Create
func (svc *ServiceClient) CreateCategory(ctx *gin.Context) {
	code, err := warehouse.PermissionHandler(2, svc.Warehouse, ctx)
	if err != nil {
		ctx.JSON(code, gin.H{"status": code, "response": gin.H{"error": err.Error()}})
		return
	}
	routes.CreateCategory(ctx, svc.Client)
}

func (svc *ServiceClient) CreateField(ctx *gin.Context) {
	code, err := warehouse.PermissionHandler(2, svc.Warehouse, ctx)
	if err != nil {
		ctx.JSON(code, gin.H{"status": code, "response": gin.H{"error": err.Error()}})
		return
	}

	routes.CreateField(ctx, svc.Client)
}

func (svc *ServiceClient) CreateEntity(ctx *gin.Context) {
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
	routes.CreateEntity(ctx, svc.Client)
}

// Edit
func (svc *ServiceClient) EditCategory(ctx *gin.Context) {
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
	routes.EditCategory(ctx, svc.Client)
}

func (svc *ServiceClient) EditField(ctx *gin.Context) {
	code, err := warehouse.PermissionHandler(2, svc.Warehouse, ctx)
	if err != nil {
		ctx.JSON(code, gin.H{"status": code, "response": gin.H{"error": err.Error()}})
		return
	}

	code, err = CoreMiddleware(svc.Client, ctx, "fields_to_warehouses")
	if err != nil {
		ctx.JSON(code, gin.H{"status": code, "response": gin.H{"error": err.Error()}})
		return
	}
	routes.EditField(ctx, svc.Client)
}

func (svc *ServiceClient) EditEntity(ctx *gin.Context) {
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

	code, err = CoreMiddleware(svc.Client, ctx, "entities_to_categories")
	if err != nil {
		ctx.JSON(code, gin.H{"status": code, "response": gin.H{"error": err.Error()}})
		return
	}

	routes.EditEntity(ctx, svc.Client)
}

// Delete
func (svc *ServiceClient) DeleteCategory(ctx *gin.Context) {
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
	routes.DeleteCategory(ctx, svc.Client)
}

func (svc *ServiceClient) DeleteField(ctx *gin.Context) {
	code, err := warehouse.PermissionHandler(2, svc.Warehouse, ctx)
	if err != nil {
		ctx.JSON(code, gin.H{"status": code, "response": gin.H{"error": err.Error()}})
		return
	}

	code, err = CoreMiddleware(svc.Client, ctx, "fields_to_warehouses")
	if err != nil {
		ctx.JSON(code, gin.H{"status": code, "response": gin.H{"error": err.Error()}})
		return
	}
	routes.DeleteField(ctx, svc.Client)
}

func (svc *ServiceClient) DeleteEntity(ctx *gin.Context) {
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

	code, err = CoreMiddleware(svc.Client, ctx, "entities_to_categories")
	if err != nil {
		ctx.JSON(code, gin.H{"status": code, "response": gin.H{"error": err.Error()}})
		return
	}

	routes.DeleteEntity(ctx, svc.Client)
}

// Fetch Data
func (svc *ServiceClient) NewCategorySchema(ctx *gin.Context) {
	routes.NewCategorySchema(ctx, svc.Client)
}

func (svc *ServiceClient) NewFieldSchema(ctx *gin.Context) {
	routes.NewFieldSchema(ctx, svc.Client)
}

func (svc *ServiceClient) GetCategorySchema(ctx *gin.Context) {
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
	routes.GetCategorySchema(ctx, svc.Client)
}

func (svc *ServiceClient) GetCategoryData(ctx *gin.Context) {
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
	routes.GetCategoryData(ctx, svc.Client)
}

func (svc *ServiceClient) GetFieldData(ctx *gin.Context) {
	code, err := warehouse.PermissionHandler(2, svc.Warehouse, ctx)
	if err != nil {
		ctx.JSON(code, gin.H{"status": code, "response": gin.H{"error": err.Error()}})
		return
	}

	code, err = CoreMiddleware(svc.Client, ctx, "fields_to_warehouses")
	if err != nil {
		ctx.JSON(code, gin.H{"status": code, "response": gin.H{"error": err.Error()}})
		return
	}
	routes.GetFieldData(ctx, svc.Client)
}

func (svc *ServiceClient) GetEntityData(ctx *gin.Context) {
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

	code, err = CoreMiddleware(svc.Client, ctx, "entities_to_categories")
	if err != nil {
		ctx.JSON(code, gin.H{"status": code, "response": gin.H{"error": err.Error()}})
		return
	}

	routes.GetEntityData(ctx, svc.Client)
}

// List

func (svc *ServiceClient) ListFields(ctx *gin.Context) {
	code, err := warehouse.PermissionHandler(2, svc.Warehouse, ctx)
	if err != nil {
		ctx.JSON(code, gin.H{"status": code, "response": gin.H{"error": err.Error()}})
		return
	}

	routes.ListFields(ctx, svc.Client)
}

func (svc *ServiceClient) ListCategories(ctx *gin.Context) {
	code, err := warehouse.PermissionHandler(0, svc.Warehouse, ctx)
	if err != nil {
		ctx.JSON(code, gin.H{"status": code, "response": gin.H{"error": err.Error()}})
		return
	}

	routes.ListCategories(ctx, svc.Client)
}

func (svc *ServiceClient) ListEntities(ctx *gin.Context) {
	code, err := warehouse.PermissionHandler(0, svc.Warehouse, ctx)
	if err != nil {
		ctx.JSON(code, gin.H{"status": code, "response": gin.H{"error": err.Error()}})
		return
	}

	routes.ListEntities(ctx, svc.Client)
}

// Misc

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
