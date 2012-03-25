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
)

type UserToken struct {
	Name       string
	passwd     []byte
	RememberMe bool
	Host       string
}

type SecurityManager struct {
}

type Subject struct {
}

type ShiroError struct {
}

type Principal struct {
}

type Authenticator interface {
	Cred()
	Principals()
	Authenticate()
}

type AuthListener interface {
	OnFailure()
	OnLogOut()
	OnSuccess()
}

type Permission struct {
	Perm string
}

type Role struct {
	Name       string
	Permission []Permission
}

type SimpleAccount struct {
	AuthInfo AuthInfo
	Realm    Realm
	Roles    []Role
}

type User struct {
	Name    string
	Account SimpleAccount
}

func (s *SimpleAccount) AddPermission() {
}

func (s *SimpleAccount) AddPermissions() {
}

func (s *SimpleAccount) AddRole() {
}

func (s *SimpleAccount) AddRoles() {
}

func (s *SimpleAccount) Creds() Credentials {
	return Credentials{}
}

func (s *SimpleAccount) Lock() bool {
	return false
}

func (s *SimpleAccount) Permissions() []Permission {
	return []Permission{}
}

func (s *SimpleAccount) Expired() bool {
	return false
}

func (s *SimpleAccount) SetPermissions(p []Permission) {
}

func (s *SimpleAccount) SetRoles(r []string) {
}
