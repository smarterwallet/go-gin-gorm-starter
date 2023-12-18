package controller

import (
	"github.com/gin-gonic/gin"
	"go-gin-gorm-starter/models"
	"go-gin-gorm-starter/service"
	"go-gin-gorm-starter/utils/gin2"
	"strconv"
)

func GetSimpleStrategy(ctx *gin.Context) {
	address := ctx.Query("address")

	data, err := service.FindSimpleStrategy(&models.StrategyInfo{
		Name: address,
	})
	if err != nil {
		gin2.HttpResponse(ctx, "", err)
		return
	}

	gin2.HttpResponse(ctx, data, err)
}

func SaveSimpleStrategy(ctx *gin.Context) {
	var (
		reqUser models.StrategyInfo
	)

	ctx.Bind(&reqUser)

	err := service.SaveSimpleStrategy(&reqUser)
	if err != nil {
		gin2.HttpResponse(ctx, "", err)
		return
	}

	gin2.HttpResponse(ctx, err == nil, err)
}

func DeleteSimpleStrategy(ctx *gin.Context) {
	var (
		params models.StrategyInfo
	)

	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		gin2.HttpResponse(ctx, "", err)
		return
	}

	params.ID = uint(id)

	err = service.DeleteSimpleStrategy(&params)
	if err != nil {
		gin2.HttpResponse(ctx, "", err)
		return
	}

	gin2.HttpResponse(ctx, err == nil, err)
}
