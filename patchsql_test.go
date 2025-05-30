package patchsql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestUser struct {
	ID       int     `db:"id"`
	Username string  `db:"username"`
	Email    *string `db:"email"`
	Age      int     `db:"age"`
	Note     *string `db:"note"`
	Unused   string  // no db tag
}

type TestUserDifferentTags struct {
	ID       int     `sql:"id"`
	Username string  `sql:"username"`
	Email    *string `sql:"email"`
	Age      int     `sql:"age"`
	Note     *string `sql:"note"`
	Unused   string  // no db tag
}

func TestBuildSetClause(t *testing.T) {
	username := "johndoe"
	email := "user@example.com"
	emailPtr := &email
	age := 25
	user := TestUser{
		ID:       0,
		Username: username,
		Email:    emailPtr,
		Age:      age,
		Note:     nil,
		Unused:   "value",
	}

	query, args, _ := BuildSetClause(user)

	expectedQuery := "username = $1, email = $2, age = $3"
	expectedArgs := []any{username, emailPtr, age}

	assert.Equal(t, expectedQuery, query)
	assert.Equal(t, expectedArgs, args)
}

func TestBuildSetClauseDefault(t *testing.T) {
	username := "johndoe"
	email := "user@example.com"
	emailPtr := &email
	age := 25
	user := TestUser{
		ID:       0,
		Username: username,
		Email:    emailPtr,
		Age:      age,
		Note:     nil,
		Unused:   "value",
	}

	query, args, _ := BuildSetClause(user)

	expectedQuery := "username = $1, email = $2, age = $3"
	expectedArgs := []any{username, emailPtr, age}

	assert.Equal(t, expectedQuery, query)
	assert.Equal(t, expectedArgs, args)
}

func TestBuildSetClauseWithIndex(t *testing.T) {
	username := "johndoe"
	email := "user@example.com"
	emailPtr := &email
	age := 25
	user := TestUser{
		ID:       0,
		Username: username,
		Email:    emailPtr,
		Age:      age,
		Note:     nil,
		Unused:   "value",
	}

	query, args, _ := BuildSetClause(user, WithIndex(3))

	expectedQuery := "username = $3, email = $4, age = $5"
	expectedArgs := []any{username, emailPtr, age}

	assert.Equal(t, expectedQuery, query)
	assert.Equal(t, expectedArgs, args)
}

func TestBuildSetClauseWithTag(t *testing.T) {
	username := "johndoe"
	email := "user@example.com"
	emailPtr := &email
	age := 25
	user := TestUserDifferentTags{
		ID:       0,
		Username: username,
		Email:    emailPtr,
		Age:      age,
		Note:     nil,
		Unused:   "value",
	}

	query, args, _ := BuildSetClause(user, WithTag("sql"))

	expectedQuery := "username = $1, email = $2, age = $3"
	expectedArgs := []any{username, emailPtr, age}

	assert.Equal(t, expectedQuery, query)
	assert.Equal(t, expectedArgs, args)
}
