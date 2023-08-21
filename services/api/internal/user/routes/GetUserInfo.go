package routes

import (
	"context"
	"net/http"
	"strings"

	pb "github.com/CrabStash/crab-stash-protofiles/user/proto"
	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func GetUserInfo(ctx *gin.Context, c pb.UserServiceClient) {
	payload := pb.GetUserInfoRequest{}

	payload.Id = strings.Split(ctx.Param("id"), "/")[0]

	_, err := valid.ValidateStruct(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "response": gin.H{"error": err.Error()}})
		return
	}
	res, err := c.GetUserInfo(context.Background(), &payload)
	if err != nil {
		ctx.JSON(int(res.Status), res)
		return
	}

	ctx.JSON(int(res.Status), res)

}
