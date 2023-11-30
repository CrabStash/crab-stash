package routes

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	pb "github.com/CrabStash/crab-stash-protofiles/core/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/structpb"
)

func CreateEntity(ctx *gin.Context, c pb.CoreServiceClient) {
	payload := pb.CreateEntityRequest{}

	CategoryID := strings.Split(ctx.Param("categoryID"), "/")[0]
	WarehouseID := strings.Split(ctx.Param("warehouseID"), "/")[0]

	byteBody, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "response": gin.H{"error": err.Error()}})
	}
	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(byteBody))

	bytesToMap := make(map[string]interface{})

	if err := json.Unmarshal(byteBody, &bytesToMap); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "response": gin.H{"error": err.Error()}})
	}

	formData, err := structpb.NewStruct(bytesToMap)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "response": gin.H{"error": err.Error()}})
		return
	}

	payload.FormData = formData
	payload.CategoryID = CategoryID
	payload.WarehouseID = WarehouseID

	if len(bytesToMap) == 0 || CategoryID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "response": gin.H{"error": "missing values"}})
		return
	}

	res, _ := c.CreateEntity(context.Background(), &payload)

	ctx.JSON(int(res.Status), res)
}
