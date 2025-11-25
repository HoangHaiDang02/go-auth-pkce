package authorzation

import (
	"fmt"
	"go-oauth-pkce/src/ultils"
)

var AccessToken string
var ChallengeCodeHolder string

func TokenHolder(challengeCode string) {
	ChallengeCodeHolder = challengeCode
	fmt.Println("User requested login action")
	fmt.Println("Saved Challenge Code:", challengeCode)
	fmt.Println("Request User Login Form:", challengeCode)
	return
}

func Authorization(authenticateCode string) string {
	fmt.Println("Authorization Received Authenticate Code:", authenticateCode)
	// check permission
	authorizationCode := ultils.GenerateRandomString()
	fmt.Println("Generate Authorization Code:", authorizationCode)
	return authorizationCode
}

func GetResourceAccessToken(authorizationCode string, codeVerifier string) string {
	fmt.Println("Received Authorization Code:", authorizationCode)

	if !verifyPKCE(codeVerifier) {
		fmt.Println("Invalid codeVerifier")
		return ""
	}
	accessToken := ultils.GenerateRandomString()
	AccessToken = accessToken

	return accessToken
}

func verifyPKCE(codeVerifier string) bool {
	if ultils.HashingCodeChallenge(codeVerifier) != ChallengeCodeHolder {
		fmt.Println("PKCE Verify failed")
		return false
	}
	fmt.Println("PKCE Verify successfully")

	return true
}
