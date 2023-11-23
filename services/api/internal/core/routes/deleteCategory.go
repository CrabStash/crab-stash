package routes

import (
	"context"
	"log"
	"net/http"
	"strings"

	pb "github.com/CrabStash/crab-stash-protofiles/core/proto"
	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func DeleteCategory(ctx *gin.Context, c pb.CoreServiceClient) {
	payload := pb.GenericFetchRequest{}

	EntityID := strings.Split(ctx.Param("id"), "/")[0]
	WarehouseID := strings.Split(ctx.Param("warehouseID"), "/")[0]

	payload.EntityID = EntityID
	payload.WarehouseID = WarehouseID

	_, err := valid.ValidateStruct(&payload)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "response": gin.H{"error": err.Error()}})
		return
	}

	res, _ := c.DeleteCategory(context.Background(), &payload)

	ctx.JSON(int(res.Status), res)
}
