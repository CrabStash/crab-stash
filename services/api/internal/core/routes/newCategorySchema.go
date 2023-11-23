package routes

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	pb "github.com/CrabStash/crab-stash-protofiles/core/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
)

func NewCategorySchema(ctx *gin.Context, c pb.CoreServiceClient) {
	schema, _ := c.NewCategorySchema(context.Background(), &emptypb.Empty{})

	var placeholder map[string]interface{}

	if err := json.Unmarshal(schema.FileContent, &placeholder); err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "response": gin.H{"data": placeholder}})
	return
}
