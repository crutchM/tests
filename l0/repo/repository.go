package repo

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
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

func (s *Repository) Get(id string) data.Order {
	value, err := s.cache.Get(id)
	if !err {
		value := s.db.GetRow(id)
		s.cache.Set(value.Uid, value, 0)
	}
	var order data.Order
	buf := &bytes.Buffer{}
	gob.NewEncoder(buf).Encode(fmt.Sprint(value))
	marshaled, _ := json.Marshal(buf.Bytes())
	json.Unmarshal(marshaled, &order)
	return order
}

func (s *Repository) fillCache() {
	for _, v := range s.db.GetAll() {
		s.cache.Set(v.Uid, v, 0)
	}
}
