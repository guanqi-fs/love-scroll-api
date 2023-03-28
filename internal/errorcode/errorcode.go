package errorcode

type ErrorCode struct {
	Code    int
	Message string
}

var (
	SuccessCode = &ErrorCode{
		Code: 10000,
		Message: "操作成功",
	}
	ServiceErr = &ErrorCode{
		Code: 10001,
		Message: "服务错误",
	}
	GenerateTokenErr = &ErrorCode{
		Code: 10002,
		Message: "生成Token错误",
	}
	DbErr = &ErrorCode{
		Code: 10003,
		Message: "数据库操作错误",
	}
	BadRequestArgs = &ErrorCode{
		Code: 10004,
		Message: "错误的请求参数",
	}
	Unauthorized = &ErrorCode{
		Code: 10005,
		Message: "权限认证失败",
	}
)

