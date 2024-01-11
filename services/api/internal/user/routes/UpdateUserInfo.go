package routes

import (
	"context"
	"fmt"
	"net/http"

	pb "github.com/CrabStash/crab-stash-protofiles/user/proto"
	"github.com/CrabStash/crab-stash/api/internal/utils"
	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func UpdateUserInfo(ctx *gin.Context, c pb.UserServiceClient, utils *utils.Utils) {
	payload := pb.UpdateUserInfoRequest{}
	uuid, _ := ctx.Get("uuid")
	payload.UserID = uuid.(string)

	if err := ctx.Bind(&payload.Data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "response": gin.H{"error": err.Error()}})
		return
	}

	avatar, err := ctx.FormFile("avatar")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "response": gin.H{"error": err.Error()}})
		return
	}

	avatarURL, err := utils.UploadFile(avatar, uuid.(string))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "response": gin.H{"error": err.Error()}})
		return
	}

	payload.Data.Avatar = avatarURL

	_, err = valid.ValidateStruct(&payload)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "response": gin.H{"error": err.Error()}})
		return
	}

	res, _ := c.UpdateUserInfo(context.Background(), &payload)
	if res.Status >= 300 {
		err = utils.RestoreFile(uuid.(string))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "response": gin.H{"error": fmt.Sprintf("error while updating user info: %s, error while restoring file after a invalid userUpdate: %e", res.Response, err)}})
			return
		}

		ctx.JSON(int(res.Status), res)
		return
	}

	ctx.JSON(int(res.Status), res)

}
