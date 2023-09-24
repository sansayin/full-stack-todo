package globalkey

/**
global constant key
*/

// 软删除
var (
	DelStateNo  int64 = 0 // 未删除
	DelStateYes int64 = 1 // 已删除
)

// 时间格式化模版
var (
	DateTimeFormatTplStandardDateTime = "Y-m-d H:i:s"
	DateTimeFormatTplStandardDate     = "Y-m-d"
	DateTimeFormatTplStandardTime     = "H:i:s"
)
