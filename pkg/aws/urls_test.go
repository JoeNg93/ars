package aws_test

import (
	"fmt"
	"go-scripts/switch-aws-roles/pkg/aws"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogoutURL(t *testing.T) {
	expected := "https://signin.aws.amazon.com/oauth?Action=logout"
	logoutURL := aws.LogoutURL{}
	actual := logoutURL.String()
	assert.Equal(t, expected, actual)
}

func TestLoginURL(t *testing.T) {
	tests := []struct {
		destinationURL string
		signinToken    string
		expected       string
	}{
		{
			destinationURL: "https://console.aws.amazon.com",
			signinToken:    "foo",
			expected:       "https://signin.aws.amazon.com/federation?Action=login&Destination=https%3A%2F%2Fconsole.aws.amazon.com&Issuer=&SigninToken=foo",
		},
		{
			destinationURL: "",
			signinToken:    "bar",
			expected:       "https://signin.aws.amazon.com/federation?Action=login&Destination=&Issuer=&SigninToken=bar",
		},
	}
	for _, test := range tests {
		loginURL := aws.LoginURL{
			DestinationURL: test.destinationURL,
			SigninToken:    test.signinToken,
		}
		actual := loginURL.String()
		assert.Equal(t, test.expected, actual, fmt.Sprintf("stringify of %+v", loginURL))
	}
}

func TestTokenURL(t *testing.T) {
	tokenURL := aws.TokenURL{
		SessionID:    "my_session_id",
		SessionKey:   "my_session_key",
		SessionToken: "my_session_token",
	}
	expected := "https://signin.aws.amazon.com/federation?Action=getSigninToken&Session=%7B%22sessionId%22%3A%22my_session_id%22%2C%22sessionKey%22%3A%22my_session_key%22%2C%22sessionToken%22%3A%22my_session_token%22%7D"
	actual := tokenURL.String()
	assert.Equal(t, expected, actual)
}
