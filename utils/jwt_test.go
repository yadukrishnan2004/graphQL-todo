package utils

import (
	"os"
	"testing"
)


func init() {
	os.Setenv("JWT_SECRET", "supersecret")
}
func TestGenerateToken(t *testing.T) {

	token, err := GenerateToken(1)

	if err != nil {
		t.Errorf("failed to generate token: %v", err)
	}

	if token == "" {
		t.Errorf("token is empty")
	}
}

func TestValidateToken(t *testing.T) {

	tokenString, _ := GenerateToken(1)

	token, err := ValidateToken(tokenString)

	if err != nil {
		t.Errorf("failed to validate token: %v", err)
	}

	if !token.Valid {
		t.Errorf("token is invalid")
	}
}

func TestValidateInvalidToken(t *testing.T) {

	token, err := ValidateToken("invalidtoken")

	if err == nil && token.Valid {
		t.Errorf("invalid token passed validation")
	}
}