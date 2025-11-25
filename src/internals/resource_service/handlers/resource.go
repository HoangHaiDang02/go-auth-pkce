package resource

import (
	"fmt"
	authorzation "go-oauth-pkce/src/internals/authorization_service/handlers"
)

func GetResource(accessToken string) {
	if !verifyAccessToken(accessToken) {
		fmt.Println("Access token is invalid")
		return
	}

	fmt.Println("This is resource.")
}

func verifyAccessToken(accessToken string) bool {
	if accessToken == authorzation.AccessToken {
		return true
	}

	return false
}
