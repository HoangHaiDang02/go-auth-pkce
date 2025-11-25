package authenticate_service

import (
	"fmt"
	"go-oauth-pkce/src/ultils"
)

func Authenticate() string {
	authenticateToken := ultils.GenerateRandomString()
	fmt.Println("Generated Authenticate:", authenticateToken)
	return authenticateToken
}
