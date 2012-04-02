package shiro

type Principal struct {
}

type Subject struct {
	permitted     bool
	authenticated bool
	remembered    bool
	host          string
	principals    []Principal
	secMan        SecurityManager
	session       Session
	newSession    bool
}

func (s *Subject) Principals() []Principal {
	return s.principals
}

func (s *Subject) Permitted(p Permission) (bool, error) {
	return false, ErrAuthorization
}

func (s *Subject) Authenticated() bool {
	return false
}

func (s *Subject) Remembered() bool {
	return false
}

func (s *Subject) Login(u *UserToken) error {
	return ErrAuthentication
}

func (s *Subject) Logout() {
}

func (s *Subject) HasRole(r RoleName) bool {
	return false
}

func (s *Subject) HasRoles(r []RoleName) []bool {
	return make([]bool, 0)
}

func (s *Subject) Session() *Session {
	return &Session{} 
}
