package auth

import (
	"context"
	"net/http"
	"strings"

	pb "github.com/CrabStash/crab-stash-protofiles/auth/proto"
	"github.com/gin-gonic/gin"
)

type AuthMiddlewareConfig struct {
	svc *ServiceClient
}

func InitAuthMiddleware(svc *ServiceClient) AuthMiddlewareConfig {
	return AuthMiddlewareConfig{svc}
}

func (c *AuthMiddlewareConfig) AuthRequired(ctx *gin.Context) {
	auth := ctx.Request.Header.Get("Authorization")

	if auth == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "response": gin.H{"error": "wrong auth header"}})
		return
	}

	token := strings.Split(auth, "Bearer ")

	if len(token) < 2 {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "response": gin.H{"error": "wrong auth header"}})
		return
	}

	uuid, err := c.svc.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: token[1],
	})

	if err != nil {
		ctx.JSON(int(uuid.Status), uuid)
		return
	}

	ctx.Set("uuid", uuid.GetData().GetUuid())
	ctx.Next()
}
