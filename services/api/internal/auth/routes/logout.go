package routes

import (
	"context"
	"net/http"

	pb "github.com/CrabStash/crab-stash-protofiles/auth/proto"
	"github.com/gin-gonic/gin"
)

func Logout(ctx *gin.Context, c pb.AuthServiceClient) {
	payload := pb.LogoutRequest{}

	if err := ctx.BindJSON(&payload); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Logout(context.Background(), &payload)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
