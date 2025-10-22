## 介绍
golang项目：提供一个标准的数据格式返回

## 使用方法

```
go get -u github.com/wxz1211/unpla_standard_response
```

在gin中使用

```go
import (
    unpla "github.com/wxz1211/unpla_standard_response"
)
type Data struct {
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

	// DataNotFound 数据不存在
	DataNotFound = 4004
	
	//ParamRequired 缺少必填参数
	ParamRequired = 4001
	
```
