package auth

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJwtWorksWithValidInput(t *testing.T) {

	userId := "12367"
	username := "hassam"
	secretkey := "1234"

	token, err := GenerateJWT(userId, username, []byte(secretkey))
	log.Println(token)

	assert.NoError(t, err)
	assert.NotNil(t, token)

	claims, err := ParseJWT(token, []byte(secretkey))

	assert.NoError(t, err)

	log.Println(claims.UserID)
	log.Println(claims.Username)

	assert.Equal(t, userId, claims.UserID)
	assert.Equal(t, username, claims.Username)
}

func TestJwtReturnSafelyWithInvalidInput(t *testing.T) {

	tests := []struct {
		testLabel string
		userId    string
		username  string
		secretkey []byte
	}{
		{
			"input with empty id should throw error", "", "hassam", []byte("123"),
		},
		{
			"input with empty username should throw error", "123", "", []byte("123"),
		},
	}

	for _, test := range tests {
		t.Run(test.testLabel, func(t *testing.T) {
			token, err := GenerateJWT(test.userId, test.username, []byte(test.secretkey))
			log.Println(token)

			assert.Error(t, err)
			assert.Equal(t, "", token)
		})
	}
}
