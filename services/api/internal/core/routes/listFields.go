package routes

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"strings"

	pb "github.com/CrabStash/crab-stash-protofiles/core/proto"
	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func ListFields(ctx *gin.Context, c pb.CoreServiceClient) {
	payload := pb.PaginatedFieldFetchRequest{}

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

	payload.Limit = int32(limit)
	payload.Page = int32(page)
	payload.WarehouseID = strings.Split(ctx.Param("warehouseID"), "/")[0]

	_, err = valid.ValidateStruct(&payload)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "response": gin.H{"error": err.Error()}})
		return
	}

	res, _ := c.ListFields(context.Background(), &payload)

	if res.Status >= 300 {
		ctx.JSON(int(res.Status), res)
		return
	}
	ctx.JSON(int(res.Status), res)
}
