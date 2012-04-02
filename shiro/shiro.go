package shiro

import (
	"errors"
	"log"
	"os"
)

// List of common Authentication and Authorization errors
var (
	ErrSecManagerMissing  = errors.New("Security Manager Unavailable")
	ErrAuthentication     = errors.New("Error AUthenticating Subject")
	ErrAuhorization       = errors.New("User is not authorized for this role")
	ErrEmptyPermission    = errors.New("Permission string is empty")
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

var (
	Log *log.Logger
	LogFileName = "../shiro.log"
)

type SecurityManager struct {
}
type ShiroError struct {
}

func init() {
	file, err := os.Create(LogFileName)
	if err != nil {
		panic(err)
	}
	Log = log.New(file, "Error: ", 1)
}
