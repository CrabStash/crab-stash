package routes

import (
	"context"
	"log"
	"net/http"
	"strconv"

	pb "github.com/CrabStash/crab-stash-protofiles/warehouse/proto"
	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func ListWarehouses(ctx *gin.Context, c pb.WarehouseServiceClient) {
	payload := pb.ListWarehousesRequest{}

	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "15"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusInternalServerError, "response": gin.H{"error": "could not convert query parameter limit to int"}})
		return
	}

	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusInternalServerError, "response": gin.H{"error": "could not convert query parameter page to int"}})
		return
	}

	if page < 1 {
		page = 1
	}

	uuid, _ := ctx.Get("uuid")

	payload.Limit = int32(limit)
	payload.Page = int32(page)
	payload.Uuid = uuid.(string)

	_, err = valid.ValidateStruct(&payload)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "response": gin.H{"error": err.Error()}})
		return
	}

	res, _ := c.ListWarehouses(context.Background(), &payload)

	if res.Status >= 300 {
		ctx.JSON(int(res.Status), res)
		return
	}
	ctx.JSON(int(res.Status), res)
}
