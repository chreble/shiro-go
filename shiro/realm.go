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

type AccountRealm struct {
	Roles map[string]Role
	Users map[string]Account
}

func (s *AccountRealm) AccExist(user string) bool {
	return false
}

func (s *AccountRealm) Add(acc Account) {

}

func (s *AccountRealm) AddAccount(user, passwd string) {
}
