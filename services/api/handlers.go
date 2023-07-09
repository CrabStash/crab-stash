package main

import (
	"fmt"
	"net/http"

	pb "github.com/CrabStash/crab-stash/auth/proto"
	"github.com/gin-gonic/gin"
)

func LoginHandler(c pb.AuthServiceClient) gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		reqBody := new(pb.User)
		err := ctx.BindJSON(reqBody)

		if reqBody.Email == "" || reqBody.Passwd == "" {
			err = fmt.Errorf("email or password missing")
		}

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		token, err := GrpcLogin(c, *reqBody)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, token)
	}
	return fn
}

func RegisterHandler(c pb.AuthServiceClient) gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		reqBody := new(pb.User)
		err := ctx.BindJSON(reqBody)

		if reqBody.Email == "" || reqBody.Passwd == "" {
			err = fmt.Errorf("email or password missing")
		}

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = GrpcCreateUser(c, *reqBody)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": "user created"})
	}
	return fn
}
