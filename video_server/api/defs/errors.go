package defs

type Err struct {
	Error string `josn:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrResponse struct {
	HttpSC int
	Error Err
}

var (
	ErrorRequestBodyParseFaild = ErrResponse{HttpSC: 400, Error: Err{Error: "发送的请求内容无法解析", ErrorCode: "001"}}
	ErrorNotAuthUser = ErrResponse{HttpSC: 401, Error: Err{Error: "权限错误", ErrorCode: "002"}}
)