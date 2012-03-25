package shiro

type Realm struct {
	Authc AuthInfo
	Name  string
}

type CacheRealm struct {
	Principal Principal
	cache     bool
	Realm     Realm
}

type AuthcRealm struct {
	CacheRealm CacheRealm
	CacheKey   string
	CacheName  string
}

type SimpleAccRealm struct {
	Roles map[string]Role
	Users map[string]SimpleAccount
}

func (s *SimpleAccRealm) AccExist(user string) bool {
	return false
}

func (s *SimpleAccRealm) Add(acc SimpleAccount) {
}

func (s *SimpleAccRealm) AddAccount(user, passwd string) {
}
