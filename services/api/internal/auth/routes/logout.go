package routes

import (
	"context"
	"net/http"
	"os"
	"strings"

	pb "github.com/CrabStash/crab-stash-protofiles/auth/proto"
	"github.com/gin-gonic/gin"
)

func Logout(ctx *gin.Context, c pb.AuthServiceClient) {

	refresh_token, err := ctx.Cookie("refresh_token")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "response": gin.H{"error": err.Error()}})
		return
	}

	auth := ctx.Request.Header.Get("Authorization")

	if auth == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "response": gin.H{"error": err.Error()}})
		return
	}

	token := strings.Split(auth, "Bearer ")

	if len(token) < 2 {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "response": gin.H{"error": err.Error()}})
		return
	}

	res, _ := c.Logout(context.Background(), &pb.LogoutRequest{Token: token[1], Refresh: refresh_token})
	if res.Status >= 300 {
		ctx.JSON(int(res.Status), res)
		return
	}

	ctx.SetCookie("refresh_token", "", -1, "/", os.Getenv("DOMAIN"), false, true)
	ctx.SetCookie("access_token", "", -1, "/", os.Getenv("DOMAIN"), false, true)

	ctx.JSON(int(res.Status), res)
}
