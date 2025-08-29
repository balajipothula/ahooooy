package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	appcfg "dating-app/services/login/config"
	"dating-app/services/login/service"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
)

type OAuthHandler struct {
	Auth *service.AuthService
}

// ---------- Helper: build state (demo-safe) ----------
func buildState(provider string) string {
	// In prod: sign & store state in Redis to prevent CSRF
	return fmt.Sprintf("%s-%d", provider, time.Now().UnixNano())
}

func parseState(state string, expected string) bool {
	// In prod: verify signature & lookup
	return len(state) > 0
}

// ---------- Google ----------
func (h *OAuthHandler) GoogleLogin(c *fiber.Ctx) error {
	cfg := appcfg.GoogleConfig()
	url := cfg.AuthCodeURL(buildState("google"), oauth2.AccessTypeOffline)
	return c.Redirect(url, http.StatusTemporaryRedirect)
}

func (h *OAuthHandler) GoogleCallback(c *fiber.Ctx) error {
	state := c.Query("state")
	code := c.Query("code")
	if !parseState(state, "google") || code == "" {
		return c.Status(400).JSON(fiber.Map{"error": "invalid state or code"})
	}

	cfg := appcfg.GoogleConfig()
	tok, err := cfg.Exchange(context.Background(), code)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "token exchange failed"})
	}

	email, extID, err := fetchGoogleUserinfo(tok)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	jwt, err := h.Auth.LoginOrRegisterOAuth("google", email, extID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"token": jwt})
}

func fetchGoogleUserinfo(tok *oauth2.Token) (email, externalID string, err error) {
	req, _ := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v3/userinfo", nil)
	req.Header.Set("Authorization", "Bearer "+tok.AccessToken)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", "", err
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var data struct {
		Sub   string `json:"sub"`
		Email string `json:"email"`
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return "", "", err
	}
	if data.Email == "" || data.Sub == "" {
		return "", "", errors.New("google userinfo incomplete")
	}
	return data.Email, data.Sub, nil
}

// ---------- Facebook ----------
func (h *OAuthHandler) FacebookLogin(c *fiber.Ctx) error {
	cfg := appcfg.FacebookConfig()
	url := cfg.AuthCodeURL(buildState("facebook"))
	return c.Redirect(url, http.StatusTemporaryRedirect)
}

func (h *OAuthHandler) FacebookCallback(c *fiber.Ctx) error {
	state := c.Query("state")
	code := c.Query("code")
	if !parseState(state, "facebook") || code == "" {
		return c.Status(400).JSON(fiber.Map{"error": "invalid state or code"})
	}

	cfg := appcfg.FacebookConfig()
	tok, err := cfg.Exchange(context.Background(), code)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "token exchange failed"})
	}

	email, extID, err := fetchFacebookUserinfo(tok)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	jwt, err := h.Auth.LoginOrRegisterOAuth("facebook", email, extID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"token": jwt})
}

func fetchFacebookUserinfo(tok *oauth2.Token) (email, externalID string, err error) {
	// fields=email,id,name
	req, _ := http.NewRequest("GET", "https://graph.facebook.com/me?fields=id,email", nil)
	req.Header.Set("Authorization", "Bearer "+tok.AccessToken)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", "", err
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var data struct {
		ID    string `json:"id"`
		Email string `json:"email"`
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return "", "", err
	}
	if data.ID == "" {
		return "", "", errors.New("facebook userinfo missing id")
	}
	// Note: Facebook email may be empty if not granted; decide your fallback policy
	return data.Email, data.ID, nil
}

// ---------- Twitter (OAuth2) ----------
func (h *OAuthHandler) TwitterLogin(c *fiber.Ctx) error {
	cfg := appcfg.TwitterConfig()
	// Twitter requires code_challenge (PKCE) in many cases; for brevity we skip here.
	url := cfg.AuthCodeURL(buildState("twitter"))
	return c.Redirect(url, http.StatusTemporaryRedirect)
}

func (h *OAuthHandler) TwitterCallback(c *fiber.Ctx) error {
	state := c.Query("state")
	code := c.Query("code")
	if !parseState(state, "twitter") || code == "" {
		return c.Status(400).JSON(fiber.Map{"error": "invalid state or code"})
	}

	cfg := appcfg.TwitterConfig()
	tok, err := cfg.Exchange(context.Background(), code)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "token exchange failed"})
	}

	email, extID, err := fetchTwitterUserinfo(tok)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	jwt, err := h.Auth.LoginOrRegisterOAuth("twitter", email, extID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"token": jwt})
}

func fetchTwitterUserinfo(tok *oauth2.Token) (email, externalID string, err error) {
	// Twitter v2 userinfo: need to call /2/users/me with "Authorization: Bearer"
	// Email access requires Elevated + special permission; often you only get id/username.
	req, _ := http.NewRequest("GET", "https://api.twitter.com/2/users/me?user.fields=profile_image_url", nil)
	req.Header.Set("Authorization", "Bearer "+tok.AccessToken)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", "", err
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var data struct {
		Data struct {
			ID       string `json:"id"`
			Username string `json:"username"`
			// Email usually not provided via v2 without special access
		} `json:"data"`
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return "", "", err
	}
	if data.Data.ID == "" {
		return "", "", errors.New("twitter userinfo missing id")
	}
	// If email unavailable, synthesize a pseudo-email or store blank and rely on provider+extID
	return "", data.Data.ID, nil
}

