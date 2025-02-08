package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Environment allows for running examples.
//
// The Realm ID, Refresh token, and Access token can be retrieved from the Intuit OAuth 2.0 Playground for a sandbox company.
type Environment struct {
	RealmID      string
	ClientID     string
	ClientSecret string
	RedirectURL  string
	RefreshToken string
	AccessToken  string
}

func mustLoadEnvVariable(key string) string {
	val, exists := os.LookupEnv(key)
	if !exists {
		log.Panicf("Missing required environment variable: %s", key)
	}
	return val
}

func GetEnvironment() Environment {
	if err := godotenv.Load(".env"); err == nil {
		log.Println(".env variables loaded")
	}

	return Environment{
		RealmID:      mustLoadEnvVariable("REALM_ID"),
		ClientID:     mustLoadEnvVariable("CLIENT_ID"),
		ClientSecret: mustLoadEnvVariable("CLIENT_SECRET"),
		RedirectURL:  mustLoadEnvVariable("REDIRECT_URL"),
		RefreshToken: mustLoadEnvVariable("REFRESH_TOKEN"),
		AccessToken:  mustLoadEnvVariable("ACCESS_TOKEN"),
	}
}
