package routes

import (
	"context"
	"net/http"
	"strings"

	pb "github.com/CrabStash/crab-stash-protofiles/warehouse/proto"
	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func Update(ctx *gin.Context, c pb.WarehouseServiceClient) {
	payload := pb.UpdateRequest{}
	payload.WarehouseID = strings.Split(ctx.Param("id"), "/")[1]

	_, err := valid.ValidateStruct(&payload)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "response": err.Error()})
		return
	}

	res, err := c.UpdateWarehouse(context.Background(), &payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	ctx.JSON(http.StatusCreated, res)
}
