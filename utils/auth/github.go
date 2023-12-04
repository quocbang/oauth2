package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/oauth2"
)

type GetGithubUserInfoReply struct {
	ID        int64  `json:"id"`
	Login     string `json:"login"`
	Location  string `json:"location"`
	AvatarURL string `json:"avatar_url"`
	Name      string `json:"name"`
}

func GetGithubUserInfo(ctx context.Context, t *oauth2.Token) (*GetGithubUserInfoReply, error) {
	req, err := http.NewRequest(http.MethodGet, "https://api.github.com/user", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to start new request, error: %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.AccessToken))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to do request, error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		errResponse, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read google user info during error, error: %v", err)
		}
		return nil, fmt.Errorf("response with failed, status_code: %d details: %v", resp.StatusCode, string(errResponse))
	}

	var userInfo GetGithubUserInfoReply
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, fmt.Errorf("failed to decode body, error: %v", err)
	}

	return &userInfo, nil
}
