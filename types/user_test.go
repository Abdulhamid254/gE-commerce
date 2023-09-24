package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestNewUser(t *testing.T) {
	user, err := NewUser("a@gmail.com","ghost0010")
	assert.Nil(t, err)
	// below was trying to make the length of the password not to be too small
	// assert.(t, 7, len(user.EncryptedPassword))

	assert.NotNil(t, user.EncryptedPassword)
}

func TestUserPassword(t *testing.T){
	pw := "ghost0010"
	user, err := NewUser("a@gmail.com", pw)
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword(pw))


	assert.False(t, user.ValidatePassword("ghost0016780"))
}