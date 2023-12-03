package user

import (
	"github.com/CrabStash/crab-stash/api/internal/auth"
	"github.com/CrabStash/crab-stash/api/internal/user/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, authSvc *auth.ServiceClient) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(),
	}
	a := auth.InitAuthMiddleware(authSvc)
	routes := r.Group("/user")
	routes.Use(a.AuthRequired)

	routes.GET("/me", svc.MeInfo)
	routes.PUT("/update", svc.UpdateUserInfo)
	routes.GET("/:id", svc.GetUserInfo)
	routes.DELETE("/delete", svc.DeleteUser)
	routes.POST("/changePassword", svc.ChangePassword)

	return svc
}

func (svc *ServiceClient) MeInfo(ctx *gin.Context) {
	routes.MeInfo(ctx, svc.Client)
}

func (svc *ServiceClient) UpdateUserInfo(ctx *gin.Context) {
	routes.UpdateUserInfo(ctx, svc.Client)
}

func (svc *ServiceClient) GetUserInfo(ctx *gin.Context) {
	routes.GetUserInfo(ctx, svc.Client)
}

func (svc *ServiceClient) DeleteUser(ctx *gin.Context) {
	routes.DeleteUser(ctx, svc.Client)
}

func (svc *ServiceClient) ChangePassword(ctx *gin.Context) {
	routes.ChangePassword(ctx, svc.Client)
}
