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
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "response": gin.H{"error": err.Error()}})
		return
	}

	res, _ := c.RemoveUserFromWarehouse(context.Background(), &payload)
	if res.Status >= 300 {
		ctx.JSON(int(res.Status), res)
		return
	}
	ctx.JSON(int(res.Status), res)
}
