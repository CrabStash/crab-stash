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
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "response": gin.H{"error": err.Error()}})
		return
	}

	auth := ctx.Request.Header.Get("authorization")

	if auth == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "response": gin.H{"error": err.Error()}})
		return
	}

	token := strings.Split(auth, "Bearer ")

	if len(token) < 2 {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "response": gin.H{"error": err.Error()}})
		return
	}

	res, err := c.Logout(context.Background(), &pb.LogoutRequest{Token: token[1], Refresh: refresh_token})
	if err != nil {
		ctx.JSON(int(res.Status), res)
		return
	}

	exp := time.Now().Add(-time.Hour * 24)

	ctx.SetCookie("refresh_token", "", int(exp.Unix()), "/", os.Getenv("DOMAIN"), false, true)
	ctx.SetCookie("access_token", "", int(exp.Unix()), "/", os.Getenv("DOMAIN"), false, true)

	ctx.JSON(int(res.Status), res)
}
