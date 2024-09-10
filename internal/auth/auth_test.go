package auth

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Stern-Ritter/gophkeeper/internal/model"
)

func TestGetPasswordHash(t *testing.T) {
	testCases := []struct {
		password string
		wantErr  bool
	}{
		{"password", false},
		{" ", true},
		{"", true},
	}

	for _, tt := range testCases {
		hash, err := GetPasswordHash(tt.password)
		if tt.wantErr {
			assert.Error(t, err, "expected an error for password: %s", tt.password)
			assert.Empty(t, hash, "expected empty hash for password: %s", tt.password)
		} else {
			assert.NoError(t, err, "unexpected error for password: %s", tt.password)
			assert.NotEmpty(t, hash, "expected non-empty hash for password: %s", tt.password)
		}
	}
}

func TestCheckPasswordHash(t *testing.T) {
	password := "password"
	hash, err := GetPasswordHash(password)
	require.NoError(t, err, "unexpected error hashing password: %s", password)

	valid := CheckPasswordHash(password, hash)
	assert.True(t, valid, "password should match the hash")

	invalid := CheckPasswordHash("wrong password", hash)
	assert.False(t, invalid, "password should not match the hash")
}

func TestNewToken(t *testing.T) {
	secret := "secret"
	user := model.User{
		ID:    "42",
		Login: "test user",
	}
	duration := time.Hour

	token, err := NewToken(user, secret, duration)
	assert.NoError(t, err, "unexpected error generating token")
	assert.NotEmpty(t, token, "expected non-empty token")

	parsedToken, err := ValidateToken(token, secret)
	assert.NoError(t, err, "unexpected error validating token")
	assert.NotNil(t, parsedToken, "expected non-nil token after validation")

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	assert.True(t, ok, "expected claims to be of type jwt.MapClaims")
	assert.Equal(t, user.ID, claims["uid"], "user ID in claims does not match")
	assert.Equal(t, user.Login, claims["login"], "login in claims does not match")
}

func TestValidateToken(t *testing.T) {
	secret := "secret"
	user := model.User{
		ID:    "42",
		Login: "test user",
	}
	duration := time.Hour

	token, err := NewToken(user, secret, duration)
	require.NoError(t, err, "unexpected error generating token")
	assert.NotEmpty(t, token, "expected non-empty token")

	parsedToken, err := ValidateToken(token, secret)
	assert.NoError(t, err, "unexpected error validating token")
	assert.NotNil(t, parsedToken, "expected non-nil token after validation")

	expiredToken := "invalid token"

	parsedToken, err = ValidateToken(expiredToken, secret)
	assert.Error(t, err, "expected error for invalid token")
	assert.Nil(t, parsedToken, "expected nil token after validation")

	invalidToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":   user.ID,
		"login": user.Login,
		"exp":   time.Now().Add(-time.Hour).Unix(),
	}).SignedString([]byte("wrong secret"))
	require.NoError(t, err, "unexpected error generating invalid token")

	_, err = ValidateToken(invalidToken, secret)
	assert.Error(t, err, "expected error for token with wrong signature")
}
