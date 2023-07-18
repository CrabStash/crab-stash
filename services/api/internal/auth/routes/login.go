package routes

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"

	pb "github.com/CrabStash/crab-stash-protofiles/auth/proto"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context, c pb.AuthServiceClient) {
	payload := pb.LoginRequest{}

	if err := ctx.BindJSON(&payload); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if payload.Email == "" || payload.Passwd == "" {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("email or password missing"))
		return
	}

	res, err := c.Login(context.Background(), &payload)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ref_exp, err := strconv.ParseInt(os.Getenv("REFRESH_EXP"), 10, 64)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("error parsing env"))
	}

	token_exp, err := strconv.ParseInt(os.Getenv("REFRESH_EXP"), 10, 64)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("error parsing env"))
	}

	ctx.SetCookie("refresh_token", res.Refresh, 24*60*60*int(ref_exp), "/", os.Getenv("DOMAIN"), false, true)
	ctx.SetCookie("access_token", res.Token, 24*60*60*int(token_exp), "/", os.Getenv("DOMAIN"), false, true)

	ctx.JSON(http.StatusOK, gin.H{"token": res.Token})
}
