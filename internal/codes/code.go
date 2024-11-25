package codes

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Code int

var (
	Success  Code = 200
	ParamErr Code = 400

	DatabaseErr Code = 99998
	Finally     Code = 99999
)

const (
	UserNotFound Code = 11001 + iota
	UserPasswordErr
	UserStatusErr
	UserAuthExpired
	UserJWTErr
)

const (
	AccountEmpty Code = 12001 + iota
	AccountError
)

const (
	CouponNotFound Code = 13001 + iota
	CouponCannotUse
)

var _codeMsg = map[Code]string{
	Success:  "success",
	ParamErr: "参数错误",

	UserNotFound:    "用户不存在",
	UserPasswordErr: "用户密码错误",
	UserStatusErr:   "用户状态异常",
	UserAuthExpired: "用户授权过期，请重新登录",
	UserJWTErr:      "用户授权信息错误",

	AccountEmpty: "账号为空 请先配置",
	AccountError: "账号信息错误",

	CouponNotFound:  "优惠券不存在",
	CouponCannotUse: "优惠券不可用",

	DatabaseErr: "数据库错误",
	Finally:     "finally",
}

func (c Code) Msg() string {
	if msg, ok := _codeMsg[c]; ok {
		return msg
	}

	return _codeMsg[Finally]
}

type rsp struct {
	Code Code `json:"code"`
	Msg  any  `json:"msg"`
	Data any  `json:"data"`
}

func (c Code) Response(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, rsp{
		Code: c,
		Msg:  c.Msg(),
		Data: data,
	})
}

func (c Code) WithResponse(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, rsp{
		Code: c,
		Msg:  c.Msg(),
		Data: data,
	})
}

func (c Code) EmptyResponse(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, rsp{
		Code: c,
		Msg:  c.Msg(),
	})
}

func (c Code) Empty(ctx *gin.Context) {
	c.EmptyResponse(ctx)
}

func (c Code) Error() string {
	return c.Msg()
}

func As(err error) (Code, bool) {
	var errCode Code
	if errors.As(err, &errCode) {
		return errCode, true
	}
	return Finally, false
}
