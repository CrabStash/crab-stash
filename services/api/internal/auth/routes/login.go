package routes

import (
	"context"
	"net/http"
	"os"
	"strconv"

	pb "github.com/CrabStash/crab-stash-protofiles/auth/proto"
	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context, c pb.AuthServiceClient) {
	payload := pb.LoginRequest{}

	if err := ctx.BindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "response": err.Error()})
		return
	}

	_, err := valid.ValidateStruct(&payload)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "response": err.Error()})
		return
	}

	res, err := c.Login(context.Background(), &payload)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "response": err.Error()})
		return
	}

	ref_exp, err := strconv.ParseInt(os.Getenv("REFRESH_EXP"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "response": "error parsing env"})
		return
	}

	token_exp, err := strconv.ParseInt(os.Getenv("TOKEN_EXP"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "response": "error parsing env"})
		return
	}

	ctx.SetCookie("refresh_token", res.Refresh, 24*60*60*int(ref_exp), "/", os.Getenv("DOMAIN"), false, true)
	ctx.SetCookie("access_token", res.Token, 24*60*60*int(token_exp), "/", os.Getenv("DOMAIN"), false, true)

	ctx.JSON(http.StatusOK, gin.H{"status": "ok", "response": res.Token})
}
