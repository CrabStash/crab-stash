package warehouse

import (
	"github.com/CrabStash/crab-stash/api/internal/auth"
	"github.com/CrabStash/crab-stash/api/internal/warehouse/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, authSvc *auth.ServiceClient) *ServiceClient {
	a := auth.InitAuthMiddleware(authSvc)
	svc := &ServiceClient{
		Client: InitServiceClient(),
	}

	routes := r.Group("warehouse")
	routes.Use(a.AuthRequired)
	{
		routes.POST("/users/add", svc.AddUser)
		routes.POST("/create", svc.Create)
		routes.DELETE("/delete/:id", svc.Delete)
		routes.GET("/info/:id", svc.GetInfo)
		routes.DELETE("/users/delete/:warehouseID/:userID", svc.RemoveUser)
		routes.PUT("/update/:id", svc.Update)
	}

	return svc
}

func (svc *ServiceClient) Create(ctx *gin.Context) {
	routes.Create(ctx, svc.Client)
}

func (svc *ServiceClient) GetInfo(ctx *gin.Context) {
	routes.GetInfo(ctx, svc.Client)
}

func (svc *ServiceClient) Update(ctx *gin.Context) {
	routes.Update(ctx, svc.Client)
}

func (svc *ServiceClient) AddUser(ctx *gin.Context) {
	routes.AddUser(ctx, svc.Client)
}

func (svc *ServiceClient) RemoveUser(ctx *gin.Context) {
	routes.RemoveUser(ctx, svc.Client)
}

func (svc *ServiceClient) Delete(ctx *gin.Context) {
	routes.Delete(ctx, svc.Client)
}
