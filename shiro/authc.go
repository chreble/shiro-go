package shiro

import (
	"hash"
)

type HashedPasswd struct {
	algo    string
	encType string
	iter    uint8
}
type AuthInfo struct {
	Credentials interface{}
	salt        []byte
	Principals  []Principal
}

type Token struct {
}
type Credentials struct {
}

func Verify(token UserToken, info AuthInfo) bool {
	return false
}

func VerifyHashedPasswd(text string, passwd hash.Hash) bool {
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

func (s *SimpleAccount) Salt() []byte {
	return nil
}
