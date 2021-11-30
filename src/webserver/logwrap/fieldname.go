package logwrap

// 统一设定结构化日志中的字段命名
const (
	AppName       = "App"        // 具体组件，如gateway、webserver等等
	InstanceID    = "InstanceID" // 具体组件实例ID，可使用uuid V4
	ModuleName    = "Module"     // 具体模块，如global、user、project等等
	RequestID     = "RequestID"  // 可用于分布式系统日志追踪
	RequestMethod = "RequestMethod"
	RequestPath   = "RequestPath"
	RequestIP     = "IP"
	URLQuery      = "URLQuery"
	JSONParam     = "JSONParam"
	FormParam     = "FormParam"
	FileName      = "File"
	LineName      = "Line"
	FunctionName  = "Function"
)
