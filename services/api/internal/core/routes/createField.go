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

func CreateField(ctx *gin.Context, c pb.CoreServiceClient) {
	payload := pb.CreateFieldRequest{}

	WarehouseID := strings.Split(ctx.Param("warehouseID"), "/")[0]

	if err := ctx.BindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "response": gin.H{"error": err.Error()}})
		return
	}

	payload.WarehouseID = WarehouseID

	_, err := valid.ValidateStruct(&payload)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "response": gin.H{"error": err.Error()}})
		return
	}

	res, _ := c.CreateField(context.Background(), &payload)

	ctx.JSON(int(res.Status), res)
}
