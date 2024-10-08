package response

const (
	ErrCodeSuccess      = 20001 // Success
	ErrCodeParamInvalid = 20003 // Email is invalid
	ErrCodeInvalidToken = 20004 // Invalid Authen Token
)

var msg = map[int]string{
	ErrCodeSuccess:      "success",
	ErrCodeParamInvalid: "Email is invalid",
	ErrCodeInvalidToken: "Invalid Token",
}
