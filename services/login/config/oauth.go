package config

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Load from env. Example (Google):
// OAUTH_GOOGLE_CLIENT_ID, OAUTH_GOOGLE_CLIENT_SECRET, OAUTH_GOOGLE_REDIRECT_URL
// Facebook:
// OAUTH_FACEBOOK_CLIENT_ID, OAUTH_FACEBOOK_CLIENT_SECRET, OAUTH_FACEBOOK_REDIRECT_URL
// Twitter (OAuth2):
// OAUTH_TWITTER_CLIENT_ID, OAUTH_TWITTER_CLIENT_SECRET, OAUTH_TWITTER_REDIRECT_URL

func GoogleConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv("OAUTH_GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("OAUTH_GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("OAUTH_GOOGLE_REDIRECT_URL"),
		Scopes: []string{
			"openid", "email", "profile",
		},
		Endpoint: google.Endpoint,
	}
}

func FacebookConfig() *oauth2.Config {
	// Facebook uses standard OAuth2 endpoints via Meta
	return &oauth2.Config{
		ClientID:     os.Getenv("OAUTH_FACEBOOK_CLIENT_ID"),
		ClientSecret: os.Getenv("OAUTH_FACEBOOK_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("OAUTH_FACEBOOK_REDIRECT_URL"),
		Scopes:       []string{"public_profile", "email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.facebook.com/v18.0/dialog/oauth",
			TokenURL: "https://graph.facebook.com/v18.0/oauth/access_token",
		},
	}
}

func TwitterConfig() *oauth2.Config {
	// Twitter API v2 OAuth2
	return &oauth2.Config{
		ClientID:     os.Getenv("OAUTH_TWITTER_CLIENT_ID"),
		ClientSecret: os.Getenv("OAUTH_TWITTER_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("OAUTH_TWITTER_REDIRECT_URL"),
		Scopes:       []string{"tweet.read", "users.read", "offline.access"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://twitter.com/i/oauth2/authorize",
			TokenURL: "https://api.twitter.com/2/oauth2/token",
		},
	}
}

