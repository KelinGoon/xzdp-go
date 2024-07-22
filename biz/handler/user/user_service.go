package user

import (
	"context"
	"fmt"

	model "xzdp/biz/model/user"
	"xzdp/biz/service"
	"xzdp/biz/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// SendCode .
// @router /user/code [POST]
func SendCode(ctx context.Context, c *app.RequestContext) {
	var err error
	var req model.UserLoginFrom
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewSendCodeService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UserLogin .
// @router /user/login [POST]
func UserLogin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req model.UserLoginFrom
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewUserLoginService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UserInfo .
// @router /user/:id [GET]
func UserInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req model.UserLoginFrom
	err = c.BindAndValidate(&req)
	id := c.Param("id")
	fmt.Println(id)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewUserInfoService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UserMethod .
// @router /user/me [GET]
func UserMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req string
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewUserMethodService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
