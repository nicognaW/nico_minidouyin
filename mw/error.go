package mw

import (
	"context"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server/binding"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type ErrorResponse struct {
	StatusCode uint32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty" form:"status_code" query:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty" form:"status_msg" query:"status_msg"`       // 返回状态描述
}

type BizBindError struct {
	ErrType, FailField, Msg string
}

func (e BizBindError) Error() string {
	errorMessage := fmt.Sprintf("bind error: %s, cause: %s", e.FailField, e.Msg)
	errorResponse := ErrorResponse{
		StatusCode: 114514,
		StatusMsg:  &errorMessage,
	}
	jsonResponse, err := sonic.Marshal(&errorResponse)
	if err != nil {
		panic(err)
	}
	return string(jsonResponse)
}

type BizValidateError struct {
	ErrType, FailField, Msg string
}

func (e BizValidateError) Error() string {
	errorMessage := fmt.Sprintf("validate error: %s, cause: %s", e.FailField, e.Msg)
	errorResponse := ErrorResponse{
		StatusCode: 114514,
		StatusMsg:  &errorMessage,
	}
	jsonResponse, err := sonic.Marshal(&errorResponse)
	if err != nil {
		panic(err)
	}
	return string(jsonResponse)
}

func bizValidateErrFunc(field string, msg string) error {
	err := BizValidateError{
		ErrType:   "参数错误",
		FailField: "[validateFailField]: " + field,
		Msg:       "[validateErrMsg]: " + msg,
	}

	return &err
}

func bizBindErrFunc(field string, msg string) error {
	err := BizBindError{
		ErrType:   "bindErr",
		FailField: "[bindFailField]: " + field,
		Msg:       "[bindErrMsg]: " + msg,
	}

	return &err
}

func init() {
	hlog.Info("using bizerror")
	binding.SetErrorFactory(bizBindErrFunc, bizValidateErrFunc)
}

func BizErrorMiddleware() app.HandlerFunc {
	return func(ctx context.Context, rc *app.RequestContext) {
		rc.Next(ctx)
	}
}
