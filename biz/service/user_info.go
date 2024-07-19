package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	xzdp "xzdp/biz/model/xzdp"
)

type UserInfoService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUserInfoService(Context context.Context, RequestContext *app.RequestContext) *UserInfoService {
	return &UserInfoService{RequestContext: RequestContext, Context: Context}
}

func (h *UserInfoService) Run(req *string) (resp *xzdp.UserResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
