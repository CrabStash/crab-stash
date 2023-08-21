package routes

import (
	"context"
	"fmt"
	"net/http"

	pb "github.com/CrabStash/crab-stash-protofiles/user/proto"
	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func DeleteUser(ctx *gin.Context, c pb.UserServiceClient) {
	payload := pb.DeleteUserRequest{}
	uuid, _ := ctx.Get("uuid")
	payload.UserID = uuid.(string)

	_, err := valid.ValidateStruct(&payload)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "response": gin.H{"error": err.Error()}})
		return
	}

	res, err := c.DeleteUser(context.Background(), &payload)

	fmt.Println(res)
	fmt.Println(err)
	if err != nil {
		fmt.Println("huj")
		ctx.JSON(int(res.Status), res)
		return
	}

	ctx.JSON(int(res.Status), res)

}
