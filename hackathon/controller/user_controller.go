package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hackathon/models"
	"hackathon/response"
	"hackathon/service"
	"net/http"
)

func Register(ctx *gin.Context) {
	p := new(models.ParamRegister)
	if err := ctx.ShouldBind(p); err != nil {
		response.Response(ctx, http.StatusOK, response.CodeParamError, nil, response.GetErrMsg(response.CodeParamError))
		return
	}
	//数据验证
	code := service.Register(p)
	if code != 0 {
		response.Response(ctx, http.StatusOK, code, nil, response.GetErrMsg(code))
		return
	}
	response.Success(ctx, nil, "注册成功")
}

func Login(ctx *gin.Context) {
	//获取参数
	p := new(models.ParamLogin)
	if err := ctx.ShouldBind(p); err != nil {
		response.Response(ctx, http.StatusOK, response.CodeParamError, nil, response.GetErrMsg(response.CodeParamError))
		return
	}

	//数据验证
	token, code := service.Login(p)
	if code != 0 {
		response.Response(ctx, http.StatusOK, code, nil, response.GetErrMsg(code))
		return
	}
	//返回结果
	response.Success(ctx, gin.H{"token": token}, "登入成功")
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	//ctx.JSON(http.StatusOK,gin.H{"code":200,"data":gin.H{"user":dto.ToUserDto(user.(model.User))}})
	fmt.Println(user)
}
