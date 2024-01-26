package unpla

const (
	// SystemError 系统错误
	SystemError = -1
	// OK 操作成功
	OK = 200
	// NoContent 内容为空
	NoContent = 204
	// BadRequest 通用请求错误
	BadRequest = 400
	// Authorization 访问被拒绝
	Authorization = 401
	// NotFound 资源不存在
	NotFound = 404
	// TokenIsNotAvailable 服务令牌不可用
	TokenIsNotAvailable = 1000

	// DataNotFound 数据不存在
	DataNotFound = 4004

	//ParamRequired 缺少必填参数
	ParamRequired = 4001
	// LogicalException 逻辑异常
	LogicalException = 4002
	// AccountFrozen 账户被冻结
	AccountFrozen = 4003

	//DbQueryError 查询错误
	DbQueryError = 1101
	//DbInsertError 创建错误
	DbInsertError = 1102
	//DbUpdateError 更新错误
	DbUpdateError = 1103
	//DbDeleteError 更新错误
	DbDeleteError = 1104
	//PhoneNumHasBeenRegistered 手机号已经被注册
	PhoneNumHasBeenRegistered = 1301
	//PasswordPatternError 密码格式不正确
	PasswordPatternError = 1302
	//IdentifierHasBeenAuthorized 第三方唯一标识已被授权
	IdentifierHasBeenAuthorized = 1303
	//PasswordError 密码输入不正确
	PasswordError = 1304
	//ConfirmPassError 两次密码输入不一致
	ConfirmPassError = 1305
	//PhoneNumPatternError 手机号格式不正确
	PhoneNumPatternError = 1306
	// 待续...

)

// Message 状态码对应的信息
var Message map[int]string

func init() {
	Message = make(map[int]string)
	Message[SystemError] = "系统错误"
	Message[OK] = "操作成功"
	Message[NoContent] = "内容为空"
	Message[BadRequest] = "请求参数错误"
	Message[Authorization] = "访问被拒绝"
	Message[NotFound] = "资源不存在"
	Message[TokenIsNotAvailable] = "服务令牌不可用"
	Message[DataNotFound] = "数据不存在"
	Message[ParamRequired] = "缺少必填参数"
	Message[LogicalException] = "逻辑异常"
	Message[AccountFrozen] = "账户被冻结"
	Message[DbQueryError] = "查询错误"
	Message[DbInsertError] = "创建错误"
	Message[DbUpdateError] = "更新错误"
	Message[DbDeleteError] = "删除错误"
	Message[PhoneNumHasBeenRegistered] = "手机号已经被注册"
	Message[PasswordPatternError] = "密码格式不正确"
	Message[IdentifierHasBeenAuthorized] = "第三方唯一标识已被授权"
	Message[PasswordError] = "密码输入不正确"
	Message[ConfirmPassError] = "两次密码输入不一致"
	Message[PhoneNumPatternError] = "手机号格式不正确"
}
