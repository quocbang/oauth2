package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/oauth2"
)

type TokenType string

const (
	Bearer TokenType = "Bearer"
)

type GetGoogleUserInfoReply struct {
	ID            string `json:"ID"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

func GetUserInfo(ctx context.Context, t *oauth2.Token) (*GetGoogleUserInfoReply, error) {
	if t.TokenType != string(Bearer) {
		return nil, fmt.Errorf("expected is Bearer type")
	}

	url := fmt.Sprintf("%s?access_token=%s&alt=json", "https://www.googleapis.com/oauth2/v1/userinfo", t.AccessToken)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create new get user info request, error: %v", err)
	}
	idToken, ok := t.Extra("id_token").(string)
	if !ok {
		return nil, fmt.Errorf("can't find id token")
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", idToken))

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

	var googleUserInfo *GetGoogleUserInfoReply
	if err := json.NewDecoder(resp.Body).Decode(&googleUserInfo); err != nil {
		return nil, fmt.Errorf("failed to decode google user info response, error: %v", err)
	}

	return googleUserInfo, nil
}
