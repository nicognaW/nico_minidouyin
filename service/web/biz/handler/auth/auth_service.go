// Code generated by hertz generator.

package auth

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"nico_minidouyin/gen/douyin/auth"
	pb "nico_minidouyin/gen/douyin/auth"
)

var (
	conn   *grpc.ClientConn
	client pb.AuthServiceClient
)

func init() {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	var err error
	conn, err = grpc.Dial("localhost:40128", opts...)
	if err != nil {
		panic(err)
	}
	client = pb.NewAuthServiceClient(conn)
}

// Register .
// @router /douyin/user/register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.RegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, err)
		return
	}

	resp, err := client.Register(ctx, &req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, err.Error())
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// Login .
// @router /douyin/user/login/ [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.LoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, err)
		return
	}

	resp, err := client.Login(ctx, &req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, err.Error())
		return
	}

	c.JSON(consts.StatusOK, resp)
}
