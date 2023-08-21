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
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "response": gin.H{"error": err.Error()}})
		return
	}

	_, err := valid.ValidateStruct(&payload)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "response": gin.H{"error": err.Error()}})
		return
	}

	res, _ := c.Login(context.Background(), &payload)
	if res.Status >= 300 {
		ctx.JSON(int(res.Status), res)
		return
	}

	ref_exp, err := strconv.ParseInt(os.Getenv("REFRESH_EXP"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "response": gin.H{"error": "error parsing env"}})
		return
	}

	token_exp, err := strconv.ParseInt(os.Getenv("TOKEN_EXP"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "response": gin.H{"error": "error parsing env"}})
		return
	}

	ctx.SetCookie("refresh_token", res.GetData().GetRefresh(), 24*60*60*int(ref_exp), "/", os.Getenv("DOMAIN"), false, true)
	ctx.SetCookie("access_token", res.GetData().GetToken(), 24*60*60*int(token_exp), "/", os.Getenv("DOMAIN"), false, true)

	ctx.JSON(int(res.Status), gin.H{"status": "ok", "response": gin.H{"data": gin.H{"token": res.GetData().GetToken()}}})
}
