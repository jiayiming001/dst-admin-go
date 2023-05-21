package api

import (
	"dst-admin-go/service"
	"dst-admin-go/utils/dstConfigUtils"
	"dst-admin-go/vo"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func StartSpecifiedGame(ctx *gin.Context) {

	opType, _ := strconv.Atoi(ctx.DefaultQuery("type", "0"))
	cluster := dstConfigUtils.GetDstConfig().Cluster
	log.Println("正在启动指定游戏服务 type:", cluster, opType)
	service.StartSpecifiedGame(cluster, opType)

	ctx.JSON(http.StatusOK, vo.Response{
		Code: 200,
		Msg:  "start " + cluster + " game success",
		Data: nil,
	})
}

func StopSpecifiedGame(ctx *gin.Context) {

	opType, _ := strconv.Atoi(ctx.DefaultQuery("type", "0"))
	cluster := dstConfigUtils.GetDstConfig().Cluster
	log.Println("正在停止指定游戏服务 type:", cluster, opType)

	service.StopSpecifiedGame(cluster, opType)

	ctx.JSON(http.StatusOK, vo.Response{
		Code: 200,
		Msg:  "stop " + cluster + " game success",
		Data: nil,
	})
}

func GetSpecifiedDashboardInfo(ctx *gin.Context) {

	cluster := dstConfigUtils.GetDstConfig().Cluster
	ctx.JSON(http.StatusOK, vo.Response{
		Code: 200,
		Msg:  "success",
		Data: service.GetSpecifiedClusterDashboard(cluster),
	})
}
