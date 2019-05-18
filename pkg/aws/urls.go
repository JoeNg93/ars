package aws

import (
	"encoding/json"
	"net/url"
)

// LoginURL represents AWS login URL
type LoginURL struct {
	DestinationURL string
	SigninToken    string
}

func (u LoginURL) String() string {
	loginURL, _ := url.Parse("https://signin.aws.amazon.com")

	loginURL.Path = "/federation"

	query := url.Values{}
	query.Add("Action", "login")
	query.Add("Issuer", "")
	query.Add("Destination", u.DestinationURL)
	query.Add("SigninToken", u.SigninToken)
	loginURL.RawQuery = query.Encode()

	return loginURL.String()
}

// LogoutURL represents AWS logout URL
type LogoutURL struct{}

func (u LogoutURL) String() string {
	logoutURL, _ := url.Parse("https://signin.aws.amazon.com")

	logoutURL.Path = "/oauth"

	query := url.Values{}
	query.Add("Action", "logout")
	logoutURL.RawQuery = query.Encode()

	return logoutURL.String()
}

// TokenURL represents AWS token URL to fetch signin token
type TokenURL struct {
	SessionID    string
	SessionKey   string
	SessionToken string
}

func (u TokenURL) String() string {
	sessCfg := map[string]string{
		"sessionId":    u.SessionID,
		"sessionKey":   u.SessionKey,
		"sessionToken": u.SessionToken,
	}
	cfg, _ := json.Marshal(sessCfg)

	tokenURL, _ := url.Parse("https://signin.aws.amazon.com")

	tokenURL.Path = "/federation"

	query := url.Values{}
	query.Add("Action", "getSigninToken")
	query.Add("Session", string(cfg))
	tokenURL.RawQuery = query.Encode()

	return tokenURL.String()
}
