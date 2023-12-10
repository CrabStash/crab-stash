package auth

import (
	"github.com/CrabStash/crab-stash/api/internal/auth/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(),
	}

	a := InitAuthMiddleware(svc)

	routes := r.Group("/auth")
	routes.POST("/register", svc.Register)
	routes.POST("/login", svc.Login)
	routes.GET("/logout", a.AuthRequired(), svc.Logout)
	routes.GET("/refresh", svc.Refresh)

	return svc
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}

func (svc *ServiceClient) Logout(ctx *gin.Context) {
	routes.Logout(ctx, svc.Client)
}

func (svc *ServiceClient) Refresh(ctx *gin.Context) {
	routes.Refresh(ctx, svc.Client)
}
