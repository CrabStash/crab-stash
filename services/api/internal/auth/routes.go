package auth

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(),
	}

	routes := r.Group("/auth")
	routes.POST("/register", svc.Register)
	routes.POST("/login", svc.Login)
	routes.POST("/logout", svc.Logout)
	routes.POST("/refresh", svc.Refresh)

	return svc
}

func (svc *ServiceClient) Register(ctx *gin.Context) {

}

func (svc *ServiceClient) Login(ctx *gin.Context) {

}

func (svc *ServiceClient) Logout(ctx *gin.Context) {

}

func (svc *ServiceClient) Refresh(ctx *gin.Context) {

}
