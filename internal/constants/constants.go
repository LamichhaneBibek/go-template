package constants

const (
	DefaultUserName     string = "admin_username"
	DefaultUserEmail    string = "adminemail@gmail.com"
	DefaultUserPassword string = "admin_password"

	AuthorizationHeaderKey string = "Authorization"
	UserIdKey              string = "userId"
	FirstNameKey           string = "firstName"
	LastNameKey            string = "lastName"
	EmailKey               string = "email"
	UsernameKey            string = "username"
	ExpiredTimeKey         string = "exp"
)

const (
	// Token
	UnExpectedError = "Expected error"
	ClaimsNotFound  = "Claims not found"
	TokenRequired   = "token required"
	TokenExpired    = "token expired"
	TokenInvalid    = "token invalid"

	// User
	EmailExists      = "Email exists"
	UsernameExists   = "Username exists"
	PermissionDenied = "Permission denied"

	// DB
	RecordNotFound = "record not found"
)

type ResultCode int

const (
	Success         ResultCode = 0
	ValidationError ResultCode = 40001
	AuthError       ResultCode = 40101
	ForbiddenError  ResultCode = 40301
	NotFoundError   ResultCode = 40401
	LimiterError    ResultCode = 42901
	OtpLimiterError ResultCode = 42902
	CustomRecovery  ResultCode = 50001
	InternalError   ResultCode = 50002
)
