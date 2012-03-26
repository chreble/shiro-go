package shiro

import (
	"errors"
)

var (
	ErrSecManagerMissing  = errors.New("Security Manager Unavailable")
	ErrAuth               = errors.New("Error AUthenticating Subject")
	ErrAccount            = errors.New("Error in accessing Account")
	ErrConcurrentAccess   = errors.New("User Already Logged In")
	ErrWrongCredential    = errors.New("Incorrect Credentials Entered")
	ErrAccountLocked      = errors.New("Account is locked")
	ErrExcessiveAttempts  = errors.New("Authentication tried too many times")
	ErrCredentialsExpired = errors.New("Credentials Expired")
	ErrAccountDisabled    = errors.New("Account disabled")
	ErrAccountInvalid     = errors.New("Account Invalid")
	ErrUnsupportedToken   = errors.New("Unsupported Token")
	ErrUnknownAlgorithm   = errors.New("Unknown Algorithm")
)

type SecurityManager struct {
}
type ShiroError struct {
}
