package logger

type Category string
type SubCategory string
type ExtraKey string

const (
	General         Category = "General"
	Internal        Category = "Internal"
	Postgres        Category = "Postgres"
	Redis           Category = "Redis"
	Validation      Category = "Validation"
	RequestResponse Category = "RequestResponse"
)

const (
	Startup         SubCategory = "Startup"
	ExternalService SubCategory = "ExternalService"
	Migration       SubCategory = "Migration"
	Rollback        SubCategory = "Rollback"
	Select          SubCategory = "Select"
	Insert          SubCategory = "Insert"
	Update          SubCategory = "Update"
	Delete          SubCategory = "Delete"

	Api          SubCategory = "Api"
	HashPassword SubCategory = "HashPassword"
)

const (
	AppName      ExtraKey = "AppName"
	LoggerName   ExtraKey = "LoggerName"
	ClientIp     ExtraKey = "ClientIp"
	HostIp       ExtraKey = "HostIp"
	Method       ExtraKey = "Method"
	StatusCode   ExtraKey = "StatusCode"
	BodySize     ExtraKey = "BodySize"
	Path         ExtraKey = "Path"
	Latency      ExtraKey = "Latency"
	RequestBody  ExtraKey = "RequestBody"
	ResponseBody ExtraKey = "ResponseBody"
	ErrorMessage ExtraKey = "ErrorMessage"
)
