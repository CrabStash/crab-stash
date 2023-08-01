package routes

import (
	"context"
	"log"
	"net/http"
	"strings"

	pb "github.com/CrabStash/crab-stash-protofiles/warehouse/proto"
	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func RemoveUser(ctx *gin.Context, c pb.WarehouseServiceClient) {
	payload := pb.RemoveUserRequest{}
	payload.UserIds = strings.Split(ctx.Param("userID"), "/")[0]
	payload.WarehouseID = strings.Split(ctx.Param("warehouseID"), "/")[0]

	_, err := valid.ValidateStruct(&payload)

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "response": err.Error()})
		return
	}

	res, err := c.RemoveUserFromWarehouse(context.Background(), &payload)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "response": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, res)
}
