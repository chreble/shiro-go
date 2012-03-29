package shiro

import (
	"strings"
)

const (
	WildCard      = "*"
	TokenDelim    = ":"
	SubTokenDelim = ","
)

var (
	lowersubparts []string
	caseSensitive bool
)

type permSubToken map[string]bool

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

func (p Permission) SetCaseSensitive(c bool) {
	caseSensitive = c
}

func (p Permission) CaseSensitive() bool {
	return caseSensitive
}

func (p Permission) Implies(req Permission) bool {
	available := p.tokenize()
	requested := req.tokenize()
	for i, v := range requested {
		if len(available)-1 < i {
			return true
		} else {
			m1 := available[i]
			for k1, _ := range v {
				if _, ok := m1[k1]; !ok {
					return false
				}
			}
		}
	}
	for _, v := range available {
		if _, ok := v[WildCard]; !ok {
			return false
		}
	}
	return true
}

// This method will tokenize the permission string into slice of subtoken maps
func (p Permission) tokenize() []permSubToken {
	s := strings.TrimSpace(string(p))
	if p.CaseSensitive() {
		s = strings.ToLower(s)
	}
	parts := strings.Split(s, TokenDelim)
	permToken := make([]permSubToken, 0)
	for _, v := range parts {
		pst := make(permSubToken)
		subparts := strings.Split(v, SubTokenDelim)
		for _, v2 := range subparts {
			pst[v2] = true
		}
		permToken = append(permToken, pst)
	}
	return permToken
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
