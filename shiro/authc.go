package shiro

import (
	"hash"
	"time"
)

type Session struct {
	id         string
	start      time.Time
	stop       time.Time
	lastAccess time.Time
	timeout    time.Duration
	expired    bool
	host       string
	attributes map[string]interface{}
}



type Credentials struct {
}

type AuthInfo struct {
	Credentials Credentials
	salt        []byte
	Principals  []Principal
}

type UserToken struct {
	Name       string
	passwd     []byte
	RememberMe bool
	Host       string
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

func Verify(token UserToken, info AuthInfo) bool {
	return false
}

func VerifySimpleHash(text string, passwd hash.Hash) bool {
	return false
}

func VerifyPasswd(input, saved string) bool {
	return false
}
func EncodePasswd(text string) string {
	return ""
}

func StoredPasswd(info AuthInfo) Credentials {
	return Credentials{}
}

func SubmittedPasswd(token UserToken) Credentials {
	return Credentials{}
}

func Principals(token UserToken) []Principal {
	return nil
}

func (a *AuthInfo) Merge(info AuthInfo) {
}

func (u *UserToken) Passwd() []byte {
	return nil
}

func (u *UserToken) SetPasswd(passwd []byte) {
}

func (s *Account) Salt() []byte {
	return nil
}

