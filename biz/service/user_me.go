package service

import (
	"context"

	user "xzdp/biz/model/user"

	"github.com/cloudwego/hertz/pkg/app"
)

type UserMeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUserMeService(Context context.Context, RequestContext *app.RequestContext) *UserMeService {
	return &UserMeService{RequestContext: RequestContext, Context: Context}
}

func (h *UserMeService) Run() (resp *user.UserResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
