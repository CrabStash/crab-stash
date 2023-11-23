package warehouse

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	pb "github.com/CrabStash/crab-stash-protofiles/warehouse/proto"
	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func PermissionHandler(permissionLevel int, client pb.WarehouseServiceClient, ctx *gin.Context) (int, error) {
	UserID := strings.Split(ctx.Param("userID"), "/")[0]
	WarehouseID := strings.Split(ctx.Param("warehouseID"), "/")[0]
	uuid, _ := ctx.Get("uuid")

	if UserID != "" && WarehouseID != "" && UserID == WarehouseID && uuid.(string) == UserID {
		return 200, nil
	}

	payload := pb.InternalFetchWarehouseRoleRequest{}
	byteBody, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("error while reading request body")
	}
	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(byteBody))

	if WarehouseID != "" {
		payload.WarehouseID = WarehouseID
	} else if err := json.Unmarshal(byteBody, &payload); err != nil {
		return http.StatusBadRequest, fmt.Errorf("%v", err.Error())
	}

	payload.UserID = uuid.(string)

	_, err = valid.ValidateStruct(&payload)
	if err != nil {
		return http.StatusBadRequest, fmt.Errorf("%v", err.Error())
	}

	res, _ := client.InternalFetchWarehouseRole(context.Background(), &payload)

	if res.Status >= 300 {
		return int(res.Status), fmt.Errorf("%s", res.GetError())
	}

	if int32(res.GetData().Role) < int32(permissionLevel) {
		return http.StatusUnauthorized, fmt.Errorf("insufficient permissions")
	}

	return http.StatusOK, nil

}
