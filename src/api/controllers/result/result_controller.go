package result

import (
	"github.com/gin-gonic/gin"
)



import (

	ResponseService "github.com/mercadolibre/ejercicio1/src/api/services/result"
	"github.com/mercadolibre/myml/src/api/utils/apierrors"
	"net/http"
	"strconv"
)

func Ping(context *gin.Context){

	context.String(200, "pong")
}

const (
	paramUserID = "id"
	paramUrlType = "urlType"
)

func GetResponse(ctx *gin.Context){

	userID := ctx.Param(paramUserID)
	urlType := ctx.Param(paramUrlType)

	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil{
		apiErr := &apierrors.ApiError{
			Message: err.Error(),
			Status: http.StatusBadRequest,
		}
		ctx.JSON(apiErr.Status, apiErr)
		return
	}

	response, apiErr := ResponseService.GetResponse(id, urlType)
	if apiErr != nil{
		ctx.JSON(apiErr.Status, apiErr)
		return
	}

	ctx.JSON(200, response)
}

func GetResponseWg(ctx *gin.Context){

	userID := ctx.Param(paramUserID)
	urlType := ctx.Param(paramUrlType)

	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil{
		apiErr := &apierrors.ApiError{
			Message: err.Error(),
			Status: http.StatusBadRequest,
		}
		ctx.JSON(apiErr.Status, apiErr)
		return
	}

	response, apiErr := ResponseService.GetResponseWg(id, urlType)
	if apiErr != nil{
		ctx.JSON(apiErr.Status, apiErr)
		return
	}

	ctx.JSON(200, response)
}

func GetResponseCh(ctx *gin.Context){

	userID := ctx.Param(paramUserID)
	urlType := ctx.Param(paramUrlType)

	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil{
		apiErr := &apierrors.ApiError{
			Message: err.Error(),
			Status: http.StatusBadRequest,
		}
		ctx.JSON(apiErr.Status, apiErr)
		return
	}

	response, apiErr := ResponseService.GetResponseCh(id, urlType)
	if apiErr != nil{
		ctx.JSON(apiErr.Status, apiErr)
		return
	}

	ctx.JSON(200, response)
}


/*
func GetResponseCh2(ctx *gin.Context){

	userID := ctx.Param(paramUserID)
	urlType := ctx.Param(paramUrlType)

	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil{
		apiErr := &apierrors.ApiError{
			Message: err.Error(),
			Status: http.StatusBadRequest,
		}
		ctx.JSON(apiErr.Status, apiErr)
		return
	}

	response, apiErr := ResponseService.GetResponseCh2(id, urlType)
	if apiErr != nil{
		ctx.JSON(apiErr.Status, apiErr)
		return
	}

	ctx.JSON(200, response)
}
 */

