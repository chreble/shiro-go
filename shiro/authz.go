package shiro

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
	RoleNames    []string
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
