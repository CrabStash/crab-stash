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

func ListCategories(ctx *gin.Context, c pb.CoreServiceClient) {
	payload := pb.PaginatedCategoriesFetchRequest{}

	payload.WarehouseID = strings.Split(ctx.Param("warehouseID"), "/")[0]
	payload.Id = ctx.DefaultQuery("category_id", "")

	_, err := valid.ValidateStruct(&payload)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "response": gin.H{"error": err.Error()}})
		return
	}

	res, _ := c.ListCategories(context.Background(), &payload)

	if res.Status >= 300 {
		ctx.JSON(int(res.Status), res)
		return
	}
	ctx.JSON(int(res.Status), res)
}
