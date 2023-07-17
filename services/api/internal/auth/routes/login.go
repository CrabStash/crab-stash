package routes

import (
	"context"
	"net/http"

	pb "github.com/CrabStash/crab-stash-protofiles/auth/proto"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context, c pb.AuthServiceClient) {
	payload := pb.LoginRequest{}

	if err := ctx.BindJSON(&payload); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Login(context.Background(), &payload)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
