package http

import (
	"conscience-backend/model"
	"conscience-backend/service"
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

const RequestParametersError = "请求参数错误"

//var txId = ""

func RegisterFaceID(ctx *gin.Context) {
	var req model.RequestRegisterFaceID

	err := bindBody(ctx, &req)
	if err != nil {
		logger.Errorf("client %s request parameters error:%v", ctx.ClientIP(), err)
		ctx.String(http.StatusBadRequest, RequestParametersError)
		return
	}

	var txId string
	txId, _, err = service.DefaultFaceIDService.RegisterFaceID(&req)
	if err != nil {
		logger.Errorf("client %s request register faceid error: %v", ctx.ClientIP(), err)
		ctx.JSON(http.StatusOK, model.NewInternalServerErrorJsonResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, model.NewSuccessfulJsonResponse(txId))
}

func RegisterCertificate(ctx *gin.Context) {
	var req model.RequestRegisterCertificate

	err := bindBody(ctx, &req)
	if err != nil {
		logger.Errorf("client %s request parameters error:%v", ctx.ClientIP(), err)
		ctx.String(http.StatusBadRequest, RequestParametersError)
		return
	}

	var txId string
	txId, _, err = service.DefaultFaceIDService.RegisterCertificate(&req)
	if err != nil {
		logger.Errorf("client %s request register certificate error: %v", ctx.ClientIP(), err)
		ctx.JSON(http.StatusOK, model.NewInternalServerErrorJsonResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, model.NewSuccessfulJsonResponse(txId))
}

func Record(ctx *gin.Context) {
	var req model.RequestRecord

	err := bindBody(ctx, &req)
	if err != nil {
		logger.Errorf("client %s request parameters error:%v", ctx.ClientIP(), err)
		ctx.String(http.StatusBadRequest, RequestParametersError)
		return
	}

	var txId string
	txId, _, err = service.DefaultFaceIDService.Record(&req)
	if err != nil {
		logger.Errorf("client %s request record error: %v", ctx.ClientIP(), err)
		ctx.JSON(http.StatusOK, model.NewInternalServerErrorJsonResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, model.NewSuccessfulJsonResponse(txId))
}

func GetUser(ctx *gin.Context) {
	var req model.RequestGetUser

	err := bindBody(ctx, &req)
	if err != nil {
		logger.Errorf("client %s request parameters error:%v", ctx.ClientIP(), err)
		ctx.String(http.StatusBadRequest, RequestParametersError)
		return
	}

	var txIdData string
	txIdData, err = service.DefaultFaceIDService.GetUser(&req)
	if err != nil {
		logger.Errorf("client %s request get user error: %v", ctx.ClientIP(), err)
		ctx.JSON(http.StatusOK, model.NewInternalServerErrorJsonResponse(err))
		return
	}

	var r interface{}
	err2 := json.Unmarshal([]byte(txIdData), &r)
	if err2 != nil {
		logger.Errorf("client %s request get user json error: %v", ctx.ClientIP(), err2)
		ctx.JSON(http.StatusOK, model.NewInternalServerErrorJsonResponse(err2))
		return
	} else {
		//fmt.Println(r)
		ctx.JSON(http.StatusOK, model.NewSuccessfulJsonResponse(r))
		return
	}
}

func HistoryFaceIDs(ctx *gin.Context) {
	var req model.RequestHistoryFaceIDs

	err := bindBody(ctx, &req)
	if err != nil {
		logger.Errorf("client %s request parameters error:%v", ctx.ClientIP(), err)
		ctx.String(http.StatusBadRequest, RequestParametersError)
		return
	}

	var txIdData string
	txIdData, err = service.DefaultFaceIDService.HistoryFaceIDs(&req)
	if err != nil {
		logger.Errorf("client %s request history faceids error: %v", ctx.ClientIP(), err)
		ctx.JSON(http.StatusOK, model.NewInternalServerErrorJsonResponse(err))
		return
	}

	var r interface{}
	err2 := json.Unmarshal([]byte(txIdData), &r)
	if err2 != nil {
		logger.Errorf("client %s request history faceids json error: %v", ctx.ClientIP(), err2)
		ctx.JSON(http.StatusOK, model.NewInternalServerErrorJsonResponse(err2))
		return
	} else {
		//fmt.Println(r)
		ctx.JSON(http.StatusOK, model.NewSuccessfulJsonResponse(r))
		return
	}
}

func bind(ctx *gin.Context, obj interface{}) error {
	err := ctx.ShouldBind(obj)
	if err != nil {
		return err
	}

	validate, err := govalidator.ValidateStruct(obj)
	if validate {
		return nil
	}
	return err
}

func bindBody(ctx *gin.Context, obj interface{}) error {
	err := ctx.ShouldBindBodyWith(obj, binding.JSON)
	if err != nil {
		return err
	}
	validate, err := govalidator.ValidateStruct(obj)
	if validate {
		return nil
	}
	return err
}

func bindQuery(ctx *gin.Context, obj interface{}) error {
	err := ctx.BindQuery(obj)
	if err != nil {
		return err
	}
	validate, err := govalidator.ValidateStruct(obj)
	if validate {
		return nil
	}
	return err
}
