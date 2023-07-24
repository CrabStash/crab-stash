package routes

import (
	"context"
	"fmt"
	"net/http"

	pb "github.com/CrabStash/crab-stash-protofiles/auth/proto"
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context, c pb.AuthServiceClient) {
	payload := pb.RegisterRequest{}

	if err := ctx.BindJSON(&payload); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if payload.Email == "" || payload.Passwd == "" {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("email or password missing"))
		return
	}

	res, err := c.Register(context.Background(), &payload)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}