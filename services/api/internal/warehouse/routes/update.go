package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	pb "github.com/CrabStash/crab-stash-protofiles/warehouse/proto"
	"github.com/CrabStash/crab-stash/api/internal/utils"
	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func Update(ctx *gin.Context, c pb.WarehouseServiceClient, utils *utils.Utils) {
	payload := pb.UpdateRequest{}

	if err := ctx.Bind(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "response": gin.H{"error": err.Error()}})
		return
	}
	payload.WarehouseID = strings.Split(ctx.Param("warehouseID"), "/")[0]

	logo, _ := ctx.FormFile("logo")

	if logo != nil {
		logoURL, err := utils.UploadFile(logo, payload.WarehouseID)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "response": gin.H{"error": err.Error()}})
			return
		}

		payload.Logo = logoURL
	}

	_, err := valid.ValidateStruct(&payload)

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "response": gin.H{"error": err.Error()}})
		return
	}

	res, _ := c.UpdateWarehouse(context.Background(), &payload)
	if res.Status >= 300 {

		err = utils.RestoreFile(payload.WarehouseID)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "response": gin.H{"error": fmt.Sprintf("error while updating warehouse info: %s, error while restoring file after a invalid warehouse update: %e", res.Response, err)}})
			return
		}

		ctx.JSON(int(res.Status), res)
		return
	}

	ctx.JSON(int(res.Status), res)
}
