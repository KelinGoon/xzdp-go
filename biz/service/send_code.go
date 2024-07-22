package service

import (
	"context"
	"fmt"

	"xzdp/biz/dal/redis"
	xzdp "xzdp/biz/model/xzdp"
	"xzdp/biz/utils"
	"xzdp/pkg/constants"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type SendCodeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSendCodeService(Context context.Context, RequestContext *app.RequestContext) *SendCodeService {
	return &SendCodeService{RequestContext: RequestContext, Context: Context}
}

func (h *SendCodeService) Run(req *xzdp.UserLoginFrom) (resp *xzdp.Result, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()
	// todo edit your code
	phone := req.Phone
	if phone == "" {
		return nil, fmt.Errorf("phone can't be empty")
	}
	code := utils.GenerateDigits(6)
	err = redis.RedisClient.Set(h.Context, constants.LOGIN_CODE_KEY+phone, code, 0).Err()
	if err != nil {
		hlog.CtxErrorf(h.Context, "err = %s", err.Error())
		return nil, err
	}

	hlog.CtxInfof(h.Context, "code = %s", code)
	return &xzdp.Result{Success: true}, nil
}
