package errwrap

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type WebError struct {
	// webserver内部错误编码，用于检索具体的错误类型
	//
	// * 0 成功
	// * 10000-19999 公共错误。一般包括非法参数、内容不存在、服务器内部错误描述等
	// * 20000-29999 用户模块错误
	// * 30000-39999 项目模块错误
	Code int

	// 错误消息
	//
	// * 成功        表示成功
	// * 其它中文     表示错误描述
	Message string

	// 字段提示
	//
	// * 接口返回错误后的字段提示信息
	Field string

	// 返回的数据对象
	//
	// * 具体的返回结果，`JSON`对象，具体接口会提供说明描述
	Data interface{}
}

// NewWebError returns a webError
func NewWebError(code int, message string) *WebError {
	return &WebError{
		Code:    code,
		Message: message,
	}
}

func (e *WebError) WithField(f string) *WebError {
	e.Field = f
	return e
}

const errFormat = "code: %d, message: %s"

func (e *WebError) Error() string {
	return fmt.Sprintf(errFormat, e.Code, e.Message)
}

// 成功
const (
	CodeOK = 0
)

// 公共错误，模块内从10000开始递增
const (
	CodeComInvalidJSONParam = 10000 + iota
	CodeComInvalidURLParam
	CodeComInvalidFormData
	CodeComInternalServerErr
	CodeComDBServerErr
	CodeComCommitDataFailed
	CodeComGetDataFailed
)

// 用户模块错误，模块内从20000开始递增
const (
	CodeUserNotFound = 20000 + iota
	CodeUserPwdNotRight
	CodeUserCreate
	CodeUserAlreadyExist
	CodeUserUpdate
	CodeUserAgainPwdNotSame
	CodeUserOrPwdNotRight
)

// 项目模块错误，模块内从30000开始递增
const (
	CodeProjectNotFound = 30000 + iota
	CodeProjectAlreadyExist
	CodeProjectCreate
)

// 默认错误内容
var (
	ErrSuccess = NewWebError(CodeOK, "成功")

	// 公共错误
	ErrComInvalidJSONParam = NewWebError(CodeComInvalidJSONParam, "body为非法的JSON格式")
	ErrComInvalidURLParam  = NewWebError(CodeComInvalidURLParam, "非法的URL查询参数")
	ErrComInvalidFormData  = NewWebError(CodeComInvalidFormData, "表单数据不正确")
	ErrComCommitDataFailed = NewWebError(CodeComCommitDataFailed, "提交数据失败")
	ErrComGetDataFailed    = NewWebError(CodeComGetDataFailed, "获取数据失败")
	ErrComInternalServer   = NewWebError(CodeComInternalServerErr, "服务器内部错误")
	ErrComDBServer         = NewWebError(CodeComDBServerErr, "数据库服务异常")

	// 用户模块错误
	ErrUserNotFound        = NewWebError(CodeUserNotFound, "用户不存在")
	ErrUserPwdNotRight     = NewWebError(CodeUserPwdNotRight, "密码不正确")
	ErrUserCreate          = NewWebError(CodeUserCreate, "创建用户数据失败")
	ErrUserAlreadyExist    = NewWebError(CodeUserAlreadyExist, "用户已存在")
	ErrUserUpdate          = NewWebError(CodeUserUpdate, "更新用户信息失败")
	ErrUserAgainPwdNotSame = NewWebError(CodeUserAgainPwdNotSame, "两次新密码不相同")
	ErrUserOrPwdNotRight   = NewWebError(CodeUserOrPwdNotRight, "用户或密码不正确")

	// 项目模块错误
	ErrProjectNotFound     = NewWebError(CodeProjectNotFound, "项目不存在")
	ErrProjectAlreadyExist = NewWebError(CodeProjectAlreadyExist, "项目已存在")
	ErrProjectCreate       = NewWebError(CodeProjectCreate, "创建项目数据失败")
)

func WriteError(ctx *gin.Context, code int, data interface{}) {
	ctx.JSON(code, data)
}

func WriteOK(ctx *gin.Context) {
	ctx.JSON(ErrSuccess.Code, ErrSuccess.Message)
}

func WriteData(ctx *gin.Context, data interface{}) {
	ctx.JSON(ErrSuccess.Code, data)
}

func WriteList(ctx *gin.Context, offset, limit int, total int64, data interface{}) {
	//
}
