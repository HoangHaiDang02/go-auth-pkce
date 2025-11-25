package web_app

import (
	"fmt"
	authenticate_service "go-oauth-pkce/src/internals/authenticate_service/handlers"
	authorzation "go-oauth-pkce/src/internals/authorization_service/handlers"
	resource "go-oauth-pkce/src/internals/resource_service/handlers"
	"go-oauth-pkce/src/ultils"
)

func Application() {
	codeVerifier := ultils.GenerateRandomString()
	fmt.Println("Generated Code Verifier:", codeVerifier)
	codeChallenge := ultils.HashingCodeChallenge(codeVerifier)
	fmt.Println("Generated Code Challenge:", codeChallenge)

	authorzation.TokenHolder(codeChallenge)
	authenticateCode := authenticate_service.Authenticate()
	authorizationCode := authorzation.Authorization(authenticateCode)
	accessToken := authorzation.GetResourceAccessToken(authorizationCode, codeVerifier)

	if accessToken == "" {
		fmt.Println("Failed to get access token")
		return
	}

	resource.GetResource(accessToken)
}
