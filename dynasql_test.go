package dynasql

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

func TestGenSetClauseFromFlatStruct(t *testing.T) {
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

	query, args := GenSetClauseFromFlatStruct(user)

	expectedQuery := "SET username = $1, email = $2, age = $3"
	expectedArgs := []any{username, emailPtr, age}

	assert.Equal(t, expectedQuery, query)
	assert.Equal(t, expectedArgs, args)
}
