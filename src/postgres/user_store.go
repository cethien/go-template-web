package postgres

import (
	"database/sql"

	"github.com/cethien/go-template-web/src/domain"
)

type UserStore struct {
	*sql.DB
}

func (s *UserStore) Create(u *domain.User) error {
	panic("not implemented") // TODO: Implement
}

func (s *UserStore) GetMany() ([]*domain.User, error) {
	panic("not implemented") // TODO: Implement
}

func (s *UserStore) Get(id int) (*domain.User, error) {
	panic("not implemented") // TODO: Implement
}

func (s *UserStore) Update(u *domain.User) error {
	panic("not implemented") // TODO: Implement
}

func (s *UserStore) Delete(u *domain.User) error {
	panic("not implemented") // TODO: Implement
}
