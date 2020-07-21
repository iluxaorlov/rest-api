package sqlstore_test

import (
	"github.com/iluxaorlov/rest-api/internal/app/model"
	"github.com/iluxaorlov/rest-api/internal/app/store"
	"github.com/iluxaorlov/rest-api/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDb(t, url)

	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDb(t, url)

	defer teardown("users")

	email := "user@example.org"

	s := sqlstore.New(db)
	_, err := s.User().FindByEmail(email)

	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u1 := model.TestUser(t)
	u1.Email = email
	s.User().Create(u1)
	u2, err := s.User().FindByEmail(email)

	assert.NoError(t, err)
	assert.NotNil(t, u2)
}

func TestUserRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDb(t, url)

	defer teardown("users")

	s := sqlstore.New(db)
	u1 := model.TestUser(t)
	s.User().Create(u1)
	u2, err := s.User().Find(u1.Id)

	assert.NoError(t, err)
	assert.NotNil(t, u2)
}
