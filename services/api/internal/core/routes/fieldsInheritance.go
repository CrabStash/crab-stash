package routes

import (
	"context"
	"log"
	"net/http"

	pb "github.com/CrabStash/crab-stash-protofiles/core/proto"
	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func FieldsInheritance(ctx *gin.Context, c pb.CoreServiceClient) {
	payload := pb.InheritanceRequest{}

	if err := ctx.BindJSON(&payload); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "response": gin.H{"error": err.Error()}})
		return
	}

	_, err := valid.ValidateStruct(&payload)

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "response": gin.H{"error": err.Error()}})
		return
	}

	res, _ := c.FieldsInheritance(context.Background(), &payload)

	ctx.JSON(int(res.Status), res)
}
