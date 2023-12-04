package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"golang.org/x/oauth2"
)

type TokenType string

const (
	Bearer TokenType = "Bearer"
)

type GoogleOauth struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresIn    int64     `json:"expires_in"`
	Scope        string    `json:"scope"`
	TokenType    TokenType `json:"token_type"`
	IDToken      string    `json:"id_token"`
}

func GetGoogleOauthToken(code string, config oauth2.Config) (*GoogleOauth, error) {
	values := url.Values{}
	values.Add("grant_type", "authorization_code")
	values.Add("code", code)
	values.Add("client_id", config.ClientID)
	values.Add("client_secret", config.ClientSecret)
	values.Add("redirect_uri", config.RedirectURL)
	query := values.Encode()

	req, err := http.NewRequest(http.MethodPost, config.Endpoint.TokenURL, bytes.NewBufferString(query))
	if err != nil {
		return nil, err // TODO: should custom err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err // TODO: should custom err
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		responseErr, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err // TODO: should custom err
		}
		return nil, fmt.Errorf("failed to get google token, error: %v", string(responseErr))
	}

	googleResponse := &GoogleOauth{}
	if err := json.NewDecoder(resp.Body).Decode(&googleResponse); err != nil {
		return nil, err // TODO: should custom err
	}

	return googleResponse, nil
}

type GoogleUserReply struct {
	ID            string `json:"ID"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

func (g *GoogleOauth) GetGoogleUserInfo() (*GoogleUserReply, error) {
	if g.TokenType != Bearer {
		return nil, fmt.Errorf("expected is Bearer type")
	}

	url := fmt.Sprintf("%s?access_token=%s&alt=json", "https://www.googleapis.com/oauth2/v1/userinfo", g.AccessToken)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create new get user info request, error: %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", g.IDToken))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to do request to get google user info, error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		errResponse, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read google user info during error, status_code: %d error: %v", resp.StatusCode, string(errResponse))
		}
	}

	var googleUserInfo *GoogleUserReply
	if err := json.NewDecoder(resp.Body).Decode(&googleUserInfo); err != nil {
		return nil, fmt.Errorf("failed to decode google user info response, error: %v", err)
	}

	return googleUserInfo, nil
}
