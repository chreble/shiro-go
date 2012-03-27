package shiro

import (
	"strings"
)

const (
	WildCard      = "*"
	TokenDelim    = ":"
	SubTokenDelim = ","
	CaseSensitive = false
)

var (
	t             tokenSet
	AllTokens     []tokenSet
	parts         []string
	subparts      []string
	lowersubparts []string
)

type tokenSet map[string]bool

type Permission string

type RoleName string

type Role struct {
	RoleName
	permissions []Permission
}

type Account struct {
	AuthInfo  AuthInfo
	Realm     Realm
	RoleNames []string
}

type User struct {
	Name    string
	Account Account
}

func NewRole() *Role {
	p := make([]Permission, 0)
	name := RoleName("test")
	return &Role{RoleName: name, permissions: p}
}

func (p Permission) Implies(req Permission) bool {
	return false
}

func tokenize(s string) {
	s = strings.TrimSpace(s)
	parts = strings.Split(s, TokenDelim)
	t = make(tokenSet)
	AllTokens = make([]tokenSet, 0)
	for _, v := range parts {
		subparts = strings.Split(v, SubTokenDelim)
		for _, v2 := range subparts {
			t[v2] = true
			AllTokens = append(AllTokens, t)
		}
	}
}

func lowerCase() {
	lowersubparts = make([]string, 0)
	for _, v := range subparts {
		lowersubparts = append(lowersubparts, strings.ToLower(v))
	}
}

func (s *Account) AddPermission() {
}

func (r *Role) AddPermission(p Permission) {
	perm := r.Get()
	r.permissions = append(perm, p)
}

func (r *Role) AddPermissions(p []Permission) {
	for _, v := range p {
		r.AddPermission(v)
	}
}

func (r *Role) IsPermitted(p Permission) bool {
	for _, v := range r.permissions {
		if v.Implies(p) {
			return true
		}
	}
	return false
}

func (s *Account) AddPermissions() {
}

func (s *Account) AddRole() {
}

func (s *Account) AddRoles() {
}

func (s *Account) Creds() Credentials {
	return Credentials{}
}

func (s *Account) Lock() bool {
	return false
}

func (r *Role) Get() []Permission {
	return r.permissions
}

func (s *Account) Permissions() []Permission {
	return []Permission{}
}

func (s *Account) Expired() bool {
	return false
}

func (s *Account) SetPermissions(p []Permission) {
}

func (s *Account) SetRoles(r []string) {
}
