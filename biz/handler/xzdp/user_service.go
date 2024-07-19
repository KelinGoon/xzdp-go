package xzdp

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	xzdp "xzdp/biz/model/xzdp"
	"xzdp/biz/service"
	"xzdp/biz/utils"
)

// UserMethod .
// @router /user [GET]
func UserMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req xzdp.UserReq
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

// UserLogin .
// @router /user/login [POST]
func UserLogin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req xzdp.UserReq
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
