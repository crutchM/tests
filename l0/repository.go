package l0

type Repository struct {
	db    *DataBase
	cache *Cache
}

func NewRepository(db *DataBase, cache *Cache) *Repository {
	return &Repository{db: db, cache: cache}
}

func (s *Repository) Write(value Order) {
	s.cache.Set(value.Uid, value, 0)
	s.db.Write(value)
}

func (s *Repository) Get(id string) {

}

func (s *Repository) fillCache() {

}
