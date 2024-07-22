package interceptor

import (
	"context"
	"time"
	"xzdp/biz/dal/redis"
	model "xzdp/biz/model/user"
	"xzdp/biz/pkg/constants"
	"xzdp/biz/utils"

	"github.com/cloudwego/hertz/pkg/app"
)

func CheckToken(ctx context.Context, c *app.RequestContext) {
	token := c.GetHeader("authorization")
	if token == nil {
		c.Next(ctx)
	}
	if len(token) == 0 {
		c.Next(ctx)
	}
	var userdto model.UserDTO
	if err := redis.RedisClient.HGetAll(ctx, constants.LOGIN_USER_KEY+string(token)).Scan(&userdto); err != nil {
		c.Next(ctx)
	}
	redis.RedisClient.Expire(ctx, constants.LOGIN_USER_KEY+string(token), time.Minute*1)
	utils.SaveUser(ctx, &userdto)
	c.Next(ctx)
}
