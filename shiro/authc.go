package shiro

import (
	"hash"
)

type Session struct {
}

type Principal struct {
}

type Subject struct {
	authcd     bool
	host       string
	principals []string
	secMan     SecurityManager
	session    Session
	newSession bool
}

type Credentials struct {
}

type AuthInfo struct {
	Credentials interface{}
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

func LoginN(token UserToken, aggregate AuthInfo) AuthInfo {
	return AuthInfo{}
}

func Login(token UserToken, realm Realm, realmInfo, aggregate AuthInfo) AuthInfo {
	return AuthInfo{}
}

func BeforeLoginN(realms []Realm, token UserToken) AuthInfo {
	return AuthInfo{}
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
