package auth

import (
	"crypto/rand"
	"crypto/rsa"
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

func TestValidateToken_ValidToken(t *testing.T) {
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
}

func TestValidateToken_InvalidSigningMethod(t *testing.T) {
	secret := "secret"
	user := model.User{
		ID:    "42",
		Login: "test user",
	}
	duration := time.Hour

	privateKey, err := generateRSAKey()
	require.NoError(t, err, "unexpected error generating RSA key")

	invalidSigningMethodToken, err := generateRSAToken(privateKey, jwt.MapClaims{
		"uid":   user.ID,
		"login": user.Login,
		"exp":   time.Now().Add(duration).Unix(),
	})
	require.NoError(t, err, "unexpected error generating invalid signing method token")

	token, err := ValidateToken(invalidSigningMethodToken, secret)
	assert.Error(t, err, "expected error for token with unexpected signing method")
	assert.Empty(t, token, "expected empty token for invalid signing method")
}

func TestValidateToken_MalformedToken(t *testing.T) {
	secret := "secret"
	malformedToken := "not token"

	_, err := ValidateToken(malformedToken, secret)
	assert.Error(t, err, "expected error for malformed token")
}

func TestValidateExpiredToken(t *testing.T) {
	secret := "secret"
	user := model.User{
		ID:    "42",
		Login: "test user",
	}
	duration := -time.Hour

	expiredToken, err := NewToken(user, secret, duration)
	require.NoError(t, err, "unexpected error generating expired token")

	_, err = ValidateToken(expiredToken, secret)
	assert.Error(t, err, "expected error for expired token")
}

func generateRSAKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2048)
}

func generateRSAToken(privateKey *rsa.PrivateKey, claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(privateKey)
}
