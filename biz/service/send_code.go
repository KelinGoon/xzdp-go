package service

import (
	"context"
	"fmt"

	xzdp "xzdp/biz/model/xzdp"

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
	return &xzdp.Result{Success: true}, nil
}
