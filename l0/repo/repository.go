package repo

import (
	"fmt"
	"tests/l0/data"
	"tests/l0/db"
)

type Repository struct {
	db    *db.DataBase
	cache *data.Cache
}

func NewRepository(db *db.DataBase, cache *data.Cache) *Repository {
	rep := Repository{db: db, cache: cache}
	rep.fillCache()
	return &rep
}

func (s *Repository) Write(value data.Order) {
	s.cache.Set(value.Uid, value, 0)
	s.db.Write(value)
}

func (s *Repository) Get(id string) interface{} {
	value, err := s.cache.Get(id)
	if !err {
		value := s.db.GetRow(id)
		if value.Uid == "" {
			return data.Order{}
		}
		s.cache.Set(value.Uid, value, 0)
		s.Get(id)
	}
	return value
}

func (s *Repository) fillCache() {
	values := s.db.GetAll()
	for _, v := range values {
		s.cache.Set(v.Uid, v, 0)
	}
	all := s.cache.GetAll()
	fmt.Println(all)
}
