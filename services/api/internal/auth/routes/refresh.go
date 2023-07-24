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

func Refresh(ctx *gin.Context, c pb.AuthServiceClient) {

	refresh_token, err := ctx.Cookie("refresh_token")
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if refresh_token == "" {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("no refresh cookie present"))
		return
	}

	res, err := c.Refresh(context.Background(), &pb.RefreshRequest{Token: refresh_token})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	token_exp, err := strconv.ParseInt(os.Getenv("TOKEN_EXP"), 10, 64)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("error parsing env"))
	}

	ctx.SetCookie("access_token", res.Token, 24*60*60*int(token_exp), "/", os.Getenv("DOMAIN"), false, true)
	ctx.JSON(http.StatusOK, res)
}
