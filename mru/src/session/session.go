package session

import (
	"newcache"
	"result"
	"rut"
	"util"
)

type Session struct {
	Username   string
	Password   string
	ID         string
	RUT        *rut.RUT             //For run script
	Result     <-chan result.Result //For run script
	CaseResult map[string]chan *result.Result
	NewCache   *newcache.NewCache
}

func New(name, pass string) *Session {
	id := util.GenerateSessionIDByUserNameAndPassword(name, pass)
	cache := newcache.New(id)
	cache.Restore()

	return &Session{
		Username:   name,
		Password:   pass,
		ID:         id,
		NewCache:   cache,
		Result:     make(<-chan result.Result),
		CaseResult: make(map[string]chan *result.Result),
	}
}

func GenerateSessionIDByUserNameAndPassword(name, pass string) string {
	return util.GenerateSessionIDByUserNameAndPassword(name, pass)
}

func (s *Session) Destroy() {
	s.NewCache.Save()
	s.NewCache = nil
}
