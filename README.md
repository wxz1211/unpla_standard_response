## 使用方法

```
go get -u github.com/wxz1211/unpla_standard_response
```

在gin中使用

```go
import (
    ...

    unpla "github.com/wxz1211/unpla_standard_response"
    
    ...
)
type Data struct {
    ...
}

func(c *gin.Context) {
    data := &Data{}
    c.JSON(http.StatusOK, unpla.NewStandResponse(unpla.OK, unpla.Message[unpla.OK], data))
}

func(c *gin.Context) {
    c.JSON(http.StatusOK, unpla.NewErrorResponse(unpla.OK, unpla.Message[unpla.OK]))
}
```


状态码说明

```
	// SystemError 系统错误
	SystemError = -1
	// OK 操作成功
	OK = 200
	// BadRequest 通用请求错误
	BadRequest = 400
	// NotFound 资源不存在
	NotFound = 404
	// TokenIsNotAvailable 服务令牌不可用
	TokenIsNotAvailable = 1000

	// DataNotFound 数据不存在
	DataNotFound = 4004
	
	//ParamRequired 缺少必填参数
	ParamRequired = 4001
	
	//DbQueryError 查询错误
	DbQueryError = 1101
	//DbInsertError 创建错误
	DbInsertError = 1102
	//DbUpdateError 更新错误
	DbUpdateError = 1103
	//PhoneNumHasBeenRegistered 手机号已经被注册
	PhoneNumHasBeenRegistered = 1301
	//PasswordPatternError 密码格式不正确
	PasswordPatternError = 1302
	//WeChatHasBeenAuthorized 微信号已被授权
	WeChatHasBeenAuthorized = 1303
	//PasswordError 密码输入不正确
	PasswordError = 1304
	//ConfirmPassError 两次密码输入不一致
	ConfirmPassError = 1305
	// 待续...
	
```
