package warehouse

import (
	"github.com/CrabStash/crab-stash/api/internal/auth"
	"github.com/CrabStash/crab-stash/api/internal/utils"
	"github.com/CrabStash/crab-stash/api/internal/warehouse/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, authSvc *auth.ServiceClient, utils *utils.Utils) *ServiceClient {
	a := auth.InitAuthMiddleware(authSvc)
	svc := &ServiceClient{
		Client: InitServiceClient(),
		Utils:  utils,
	}

	routes := r.Group("warehouse")
	routes.Use(a.AuthRequired())
	{
		routes.POST("/users/add", svc.AddUser)
		routes.POST("/create", svc.Create)
		routes.DELETE("/delete/:warehouseID", svc.Delete)
		routes.GET("/info/:warehouseID", svc.GetInfo)
		routes.GET("", svc.ListWarehouses)
		routes.GET("/dashboard/:warehouseID", svc.Dashboard)
		routes.DELETE("/users/delete/:warehouseID/:userID", svc.RemoveUser)
		routes.PUT("/update/:warehouseID", svc.Update)
		routes.PUT("/users/role", svc.ChangeRole)
		routes.GET("/users/:warehouseID", svc.ListUsers)
	}

	return svc
}

func (svc *ServiceClient) Create(ctx *gin.Context) {
	routes.Create(ctx, svc.Client, svc.Utils)
}

func (svc *ServiceClient) GetInfo(ctx *gin.Context) {
	code, err := PermissionHandler(0, svc.Client, ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(code, gin.H{"status": code, "response": gin.H{"error": err.Error()}})
		return
	}
	routes.GetInfo(ctx, svc.Client)
}

func (svc *ServiceClient) Update(ctx *gin.Context) {
	code, err := PermissionHandler(3, svc.Client, ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(code, gin.H{"status": code, "response": gin.H{"error": err.Error()}})
		return
	}
	routes.Update(ctx, svc.Client, svc.Utils)
}

func (svc *ServiceClient) AddUser(ctx *gin.Context) {
	code, err := PermissionHandler(2, svc.Client, ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(code, gin.H{"status": code, "response": gin.H{"error": err.Error()}})
		return
	}

	routes.AddUser(ctx, svc.Client)
}

func (svc *ServiceClient) Dashboard(ctx *gin.Context) {
	code, err := PermissionHandler(0, svc.Client, ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(code, gin.H{"status": code, "response": gin.H{"error": err.Error()}})
		return
	}

	routes.Dashboard(ctx, svc.Client)
}

func (svc *ServiceClient) ChangeRole(ctx *gin.Context) {
	code, err := PermissionHandler(2, svc.Client, ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(code, gin.H{"status": code, "response": gin.H{"error": err.Error()}})
		return
	}
	routes.ChangeRole(ctx, svc.Client)
}

func (svc *ServiceClient) RemoveUser(ctx *gin.Context) {
	code, err := PermissionHandler(2, svc.Client, ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(code, gin.H{"status": code, "response": gin.H{"error": err.Error()}})
		return
	}
	routes.RemoveUser(ctx, svc.Client)
}

func (svc *ServiceClient) Delete(ctx *gin.Context) {
	code, err := PermissionHandler(4, svc.Client, ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(code, gin.H{"status": code, "response": gin.H{"error": err.Error()}})
		return
	}
	routes.Delete(ctx, svc.Client)

}

func (svc *ServiceClient) ListUsers(ctx *gin.Context) {
	code, err := PermissionHandler(0, svc.Client, ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(code, gin.H{"status": code, "response": gin.H{"error": err.Error()}})
		return
	}
	routes.ListUsers(ctx, svc.Client)
}

func (svc *ServiceClient) ListWarehouses(ctx *gin.Context) {
	routes.ListWarehouses(ctx, svc.Client)
}
