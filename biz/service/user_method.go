package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	xzdp "xzdp/biz/model/xzdp"
)

type UserMethodService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUserMethodService(Context context.Context, RequestContext *app.RequestContext) *UserMethodService {
	return &UserMethodService{RequestContext: RequestContext, Context: Context}
}

func (h *UserMethodService) Run(req *xzdp.UserReq) (resp *xzdp.UserResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
