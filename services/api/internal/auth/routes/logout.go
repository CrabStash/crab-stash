package routes

import (
	"context"
	"net/http"
	"os"
	"strings"
	"time"

	pb "github.com/CrabStash/crab-stash-protofiles/auth/proto"
	"github.com/gin-gonic/gin"
)

func Logout(ctx *gin.Context, c pb.AuthServiceClient) {

	refresh_token, err := ctx.Cookie("refresh_token")
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	auth := ctx.Request.Header.Get("authorization")

	if auth == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := strings.Split(auth, "Bearer ")

	if len(token) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	res, err := c.Logout(context.Background(), &pb.LogoutRequest{Token: token[1], Refresh: refresh_token})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	exp := time.Now().Add(-time.Hour * 24)

	ctx.SetCookie("refresh_token", "", int(exp.Unix()), "/", os.Getenv("DOMAIN"), false, true)
	ctx.SetCookie("access_token", "", int(exp.Unix()), "/", os.Getenv("DOMAIN"), false, true)

	ctx.JSON(http.StatusOK, res)
}
