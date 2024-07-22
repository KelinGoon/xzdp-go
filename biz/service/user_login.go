package service

import (
	"context"
	"fmt"

	"xzdp/biz/dal/mysql"
	"xzdp/biz/dal/redis"
	model "xzdp/biz/model/user"
	"xzdp/biz/utils"
	"xzdp/pkg/constants"

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

func (h *UserLoginService) Run(req *model.UserLoginFrom) (resp *model.Result, err error) {
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
	var user model.User
	result := mysql.DB.First(&user, "phone = ?", phone)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("user not found")
	}
	redisCode, err := redis.RedisClient.Get(h.Context, constants.LOGIN_CODE_KEY+phone).Result()
	if err != nil {
		hlog.CtxErrorf(h.Context, "err = %s", err.Error())
		return nil, err
	}
	if redisCode != code {
		return nil, fmt.Errorf("code not match")
	}

	token, err := utils.RandomUUID()
	if err != nil {
		return nil, err
	}
	fmt.Println(token)
	return &model.Result{Success: true}, nil
}

func (h *UserLoginService) createNewUser(phone string) error {
	user := model.User{
		Phone: phone,
	}
	result := mysql.DB.Create(&user)
	return result.Error
}
