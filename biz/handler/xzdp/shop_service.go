package xzdp

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	xzdp "xzdp/biz/model/xzdp"
	"xzdp/biz/service"
	"xzdp/biz/utils"
)

// ShopList .
// @router /shop-type/list [GET]
func ShopList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req xzdp.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewShopListService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
