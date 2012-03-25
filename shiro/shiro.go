package shiro

import (
	"errors"
)

var (
	ErrSecManagerMissing = errors.New("Security Manager Unavailable")
	ErrAuth              = errors.New("Error AUthenticating Subject")
	ErrAccount = errors.New("Error in accessing Account")
	ErrConcurrentAccess = errors.New("User Already Logged In")
	ErrWrongCredential = errors.New("Incorrect Credentials Entered")
	ErrAccountLocked = errors.New("Account is locked")
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

type Credentials struct {
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

type Token struct {
}

type Permission struct {
	Perm string
}
type RememberMe bool

type SimpleAccount struct {
	AuthInfo    AuthInfo
	Realm       string
	roles       []string
	permissions []Permission
}

type AuthInfo struct {
	Credentials interface{}
	CredSalt    []byte
	Principals  []Principal
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

func (s *SimpleAccount) Salt() []byte {
	return nil
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

func (a *AuthInfo) Merge(info AuthInfo) {
}

func (u *UserToken) Passwd() []byte {
	return nil
}

func (u *UserToken) SetPasswd(passwd []byte) {
}
