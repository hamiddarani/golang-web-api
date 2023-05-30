package service_errors

const (

	// Token
	UnExpectedError = "unexpected error"
	ClaimsNotFound  = "claims not found"
	TokenRequired   = "token required"
	TokenExpired    = "token expired"
	TokenInvalid    = "token invalid"

	// OTP
	OptExists   = "otp exists"
	OtpUsed     = "otp used"
	OtpNotValid = "otp invalid"
)
