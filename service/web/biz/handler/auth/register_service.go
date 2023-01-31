// Code generated by hertz generator.

package auth

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	auth "nico_minidouyin/gen/douyin/auth"
)

// Register .
// @router /douyin/user/register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.RegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(auth.RegisterResponse)

	c.JSON(consts.StatusOK, resp)
}