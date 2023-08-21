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

func Delete(ctx *gin.Context, c pb.WarehouseServiceClient) {
	id := strings.Split(ctx.Param("id"), "/")[0]

	payload := pb.DeleteRequest{}
	payload.WarehouseID = id

	_, err := valid.ValidateStruct(&payload)

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "response": err.Error()})
		return
	}

	res, err := c.DeleteWarehouse(context.Background(), &payload)
	if err != nil {
		log.Println(err)
		ctx.JSON(int(res.Status), res)
		return
	}
	ctx.JSON(int(res.Status), res)
}
