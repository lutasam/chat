package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/lutasam/chat/biz/bo"
	"github.com/lutasam/chat/biz/common"
	"github.com/lutasam/chat/biz/service"
	"github.com/lutasam/chat/biz/utils"
)

type UserController struct{}

func RegisterUserRouter(r *gin.RouterGroup) {
	userController := &UserController{}
	{
		r.GET("/get_detail", userController.GetUserDetail)
		r.POST("/update_user_info", userController.UpdateUserInfo)
	}
}

func (ins *UserController) GetUserDetail(c *gin.Context) {
	logger := utils.GetCtxLogger(c)
	jwtStruct, err := utils.GetCtxUserInfoJWT(c)
	if err != nil {
		logger.DoError(err.Error())
		utils.ResponseClientError(c, common.USERNOTLOGIN)
		return
	}
	resp, err := service.GetUserService().GetUserDetail(c, jwtStruct.UserID)
	if err != nil {
		logger.DoError(err.Error())
		if errors.Is(err, common.USERDOESNOTEXIST) {
			utils.ResponseClientError(c, common.USERDOESNOTEXIST)
			return
		} else {
			utils.ResponseServerError(c, common.UNKNOWNERROR)
			return
		}
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *UserController) UpdateUserInfo(c *gin.Context) {
	logger := utils.GetCtxLogger(c)
	jwtStruct, err := utils.GetCtxUserInfoJWT(c)
	if err != nil {
		logger.DoError(err.Error())
		utils.ResponseClientError(c, common.USERNOTLOGIN)
		return
	}
	req := &bo.UpdateUserInfoRequest{}
	err = c.ShouldBind(req)
	if err != nil {
		logger.DoError(err.Error())
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	err = service.GetUserService().UpdateUserInfo(c, req, jwtStruct.UserID)
	if err != nil {
		logger.DoError(err.Error())
		if errors.Is(err, common.USERDOESNOTEXIST) {
			utils.ResponseClientError(c, common.USERDOESNOTEXIST)
			return
		} else {
			utils.ResponseServerError(c, common.UNKNOWNERROR)
			return
		}
	}
	utils.ResponseSuccess(c, nil)
}
