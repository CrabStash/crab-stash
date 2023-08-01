package routes

import (
	"context"
	"log"
	"net/http"

	pb "github.com/CrabStash/crab-stash-protofiles/warehouse/proto"
	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func AddUser(ctx *gin.Context, c pb.WarehouseServiceClient) {
	payload := pb.AddUsersRequest{}
	if err := ctx.BindJSON(&payload); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	_, err := valid.ValidateStruct(&payload)

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "response": err.Error()})
		return
	}

	res, err := c.AddUsersToWarehouse(context.Background(), &payload)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "response": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, res)
}
