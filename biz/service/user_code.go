package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	xzdp "xzdp/biz/model/xzdp"
)

type UserCodeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUserCodeService(Context context.Context, RequestContext *app.RequestContext) *UserCodeService {
	return &UserCodeService{RequestContext: RequestContext, Context: Context}
}

func (h *UserCodeService) Run(req *xzdp.UserLoginFrom) (resp *xzdp.UserResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
