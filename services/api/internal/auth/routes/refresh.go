package routes

import (
	"context"
	"net/http"
	"os"
	"strconv"

	pb "github.com/CrabStash/crab-stash-protofiles/auth/proto"
	"github.com/gin-gonic/gin"
)

func Refresh(ctx *gin.Context, c pb.AuthServiceClient) {

	refresh_token, err := ctx.Cookie("refresh_token")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "response": gin.H{"error": err.Error()}})
		return
	}

	if refresh_token == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "response": gin.H{"error": "no refresh cookie present"}})
		return
	}

	res, _ := c.Refresh(context.Background(), &pb.RefreshRequest{Token: refresh_token})
	if res.Status >= 300 {
		ctx.JSON(int(res.Status), res)
		return
	}

	token_exp, err := strconv.ParseInt(os.Getenv("TOKEN_EXP"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "response": gin.H{"error": err.Error()}})
	}

	ctx.SetCookie("access_token", res.GetData().GetToken(), 24*60*60*int(token_exp), "/", os.Getenv("DOMAIN"), false, true)
	ctx.JSON(int(res.Status), res)
}
