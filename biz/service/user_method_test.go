package service

import (
	"context"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	xzdp "xzdp/biz/model/xzdp"
)

func TestUserMethodService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewUserMethodService(ctx, c)
	// init req and assert value
	req := &xzdp.UserReq{}
	resp, err := s.Run(req)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
