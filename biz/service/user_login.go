package service

import (
	"context"
	"fmt"

	xzdp "xzdp/biz/model/xzdp"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type UserLoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUserLoginService(Context context.Context, RequestContext *app.RequestContext) *UserLoginService {
	return &UserLoginService{RequestContext: RequestContext, Context: Context}
}

func (h *UserLoginService) Run(req *xzdp.UserLoginFrom) (resp *xzdp.Result, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()
	// todo edit your code
	phone := req.Phone
	code := req.Code
	if phone == "" || code == "" {
		return nil, fmt.Errorf("phone or code can't be empty")
	}
	return &xzdp.Result{Success: true}, nil
}
