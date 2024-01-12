package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"

	pb "github.com/CrabStash/crab-stash-protofiles/warehouse/proto"
	"github.com/CrabStash/crab-stash/api/internal/utils"
	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Create(ctx *gin.Context, c pb.WarehouseServiceClient, utils *utils.Utils) {
	payload := pb.CreateRequest{}
	if err := ctx.Bind(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "response": gin.H{"error": err.Error()}})
		return
	}

	userID, _ := ctx.Get("uuid")
	payload.OwnerID = userID.(string)
	warehouseID, err := uuid.NewV7()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "response": gin.H{"error": err.Error()}})
		return
	}

	payload.ID = fmt.Sprintf("warehouse:⟨%s⟩", warehouseID)

	logo, _ := ctx.FormFile("logo")

	if logo != nil {
		logoURL, err := utils.UploadFile(logo, payload.ID)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "response": gin.H{"error": err.Error()}})
			return
		}

		payload.Logo = logoURL
	}

	_, err = valid.ValidateStruct(&payload)

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "response": gin.H{"error": err.Error()}})
		return
	}

	res, _ := c.CreateWarehouse(context.Background(), &payload)
	if res.Status >= 300 {
		if logo != nil {
			err = utils.DeleteFile(payload.ID)

			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "response": gin.H{"error": fmt.Sprintf("error while creating warehouse: %s, error while deleting file after invalid warehouse creation: %e", res.Response, err)}})
				return
			}
		}
		ctx.JSON(int(res.Status), res)
		return
	}

	ctx.JSON(int(res.Status), res)
}
