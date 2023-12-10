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

func (conf *AuthMiddlewareConfig) AuthRequired() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		_, err := ctx.Cookie("access_token")

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "user not logged in"})
			return
		}

		auth := ctx.Request.Header.Get("Authorization")

		if auth == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "response": gin.H{"error": "wrong auth header"}})
			return
		}

		token := strings.Split(auth, "Bearer ")

		if len(token) < 2 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "response": gin.H{"error": "wrong auth header"}})
			return
		}

		uuid, err := conf.svc.Client.Validate(context.Background(), &pb.ValidateRequest{
			Token: token[1],
		})

		if err != nil {
			ctx.AbortWithStatusJSON(int(uuid.Status), uuid)
			return
		}

		if uuid.Status >= 400 {
			ctx.AbortWithStatusJSON(int(uuid.Status), uuid.Response)
			return
		}

		ctx.Set("uuid", uuid.GetData().GetUuid())
		ctx.Next()
	}
}
