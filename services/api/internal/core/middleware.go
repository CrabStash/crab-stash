package core

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	pb "github.com/CrabStash/crab-stash-protofiles/core/proto"
	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func CoreMiddleware(client pb.CoreServiceClient, ctx *gin.Context, target string) (int, error) {

	payload := pb.GenericFetchRequest{}

	byteBody, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("error while reading request body")
	}
	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(byteBody))

	if err := json.Unmarshal(byteBody, &payload); err != nil {
		return http.StatusBadRequest, fmt.Errorf("%v", err.Error())
	}
	payload.Type = target

	_, err = valid.ValidateStruct(&payload)
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
