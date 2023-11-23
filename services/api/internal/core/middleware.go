package core

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	pb "github.com/CrabStash/crab-stash-protofiles/core/proto"
	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func CoreMiddleware(client pb.CoreServiceClient, ctx *gin.Context, target string) (int, error) {

	payload := pb.CoreMiddlewareRequest{}

	EntityID := strings.Split(ctx.Param("id"), "/")[0]
	CategoryID := strings.Split(ctx.Param("categoryID"), "/")[0]
	WarehouseID := strings.Split(ctx.Param("warehouseID"), "/")[0]

	payload.Type = target
	if target == "entities_to_categories" {
		payload.Out = CategoryID
	} else {
		payload.Out = WarehouseID
	}
	// create -> cat_id, whID -> categories_to_warehouses
	// edit i fetch -> id, cat_id, whID -> categories_to_warehouses && entities_to_categories
	if target == "categories_to_warehouses" && EntityID == "" {
		payload.In = CategoryID
	} else if EntityID != "" && CategoryID != "" && target == "categories_to_warehouses" {
		payload.In = CategoryID
	} else {
		payload.In = EntityID
	}

	_, err := valid.ValidateStruct(&payload)
	if err != nil {
		return http.StatusBadRequest, fmt.Errorf("%v", err.Error())
	}

	res, err := client.CoreMiddleware(context.Background(), &payload)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	if !res.DoesItBelong {
		return http.StatusNotFound, fmt.Errorf("requested resource does not exist in this warehouse")
	}

	return http.StatusOK, nil
}
