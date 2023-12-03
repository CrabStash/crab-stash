package routes

import (
	"context"
	"net/http"

	pb "github.com/CrabStash/crab-stash-protofiles/user/proto"
	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func ChangePassword(ctx *gin.Context, c pb.UserServiceClient) {
	payload := pb.ChangePasswordRequest{}
	uuid, _ := ctx.Get("uuid")
	payload.UserID = uuid.(string)

	if err := ctx.BindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "response": gin.H{"error": err.Error()}})
		return
	}
	_, err := valid.ValidateStruct(&payload)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "response": gin.H{"error": err.Error()}})
		return
	}
	res, _ := c.ChangePassword(context.Background(), &payload)
	if res.Status >= 300 {
		ctx.JSON(int(res.Status), res)
		return
	}

	ctx.JSON(int(res.Status), res)

}
