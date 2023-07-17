package routes

import (
	"context"
	"fmt"
	"net/http"

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

	ctx.JSON(http.StatusOK, res)
}
