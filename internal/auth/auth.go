package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Oauth struct {
	Conf        *oauth2.Config
	AuthCodeUrl string
}

type UserInfo struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Picture    string `json:"picture"`
}

func NewOauth(redirectUrl string, clientID string, clientSecret string, scope []string) *Oauth {
	config := &oauth2.Config{
		RedirectURL:  redirectUrl,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     google.Endpoint,
		Scopes:       scope,
	}
	AuthCodeUrl := config.AuthCodeURL("state")
	return &Oauth{
		Conf:        config,
		AuthCodeUrl: AuthCodeUrl,
	}
}

func (oa *Oauth) GetGoogleUserInfo(code string) (*UserInfo, error) {
	tok, errt := oa.Conf.Exchange(context.TODO(), code)
	if errt != nil {
		log.Fatal(errt)
	}
	res, err := http.Get(fmt.Sprintf("https://www.googleapis.com/oauth2/v2/userinfo?access_token=%s", tok.AccessToken))
	if err != nil {
		return nil, err
	}

	defer (func() {
		if err := res.Body.Close(); err != nil {
			log.Printf("failed to close response body: %v", err)
		}
	})()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	userInfo := &UserInfo{}
	err = json.Unmarshal(body, userInfo)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}
