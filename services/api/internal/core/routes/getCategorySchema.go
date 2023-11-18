package routes

import (
	"context"
	"log"
	"net/http"

	pb "github.com/CrabStash/crab-stash-protofiles/core/proto"
	"github.com/gin-gonic/gin"
)

func GetCategorySchema(ctx *gin.Context, c pb.CoreServiceClient) {
	payload := pb.GenericFetchRequest{}

	if err := ctx.BindJSON(&payload); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "response": gin.H{"error": err.Error()}})
		return
	}

	res, _ := c.GetCategorySchema(context.Background(), &payload)

	ctx.JSON(int(res.Status), res)
}
