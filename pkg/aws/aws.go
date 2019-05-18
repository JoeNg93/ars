package aws

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/go-ini/ini"
)

type signinTokenResponse struct {
	SigninToken string `json:"SigninToken"`
}

// GetRoles reads aws credentials file to get a list of AWS roles
func GetRoles() ([]string, error) {
	var roles []string

	credsPath := path.Join(os.Getenv("HOME"), ".aws", "credentials")
	cfg, err := ini.Load(credsPath)
	if err != nil {
		return nil, err
	}

	var roleSections []*ini.Section
	for _, s := range cfg.SectionStrings() {
		section := cfg.Section(s)
		if section.HasKey("role_arn") && section.HasKey("source_profile") {
			roleSections = append(roleSections, section)
		}
	}

	if len(roleSections) == 0 {
		return nil, fmt.Errorf("no roles found in %v", credsPath)
	}

	for _, s := range roleSections {
		roles = append(roles, s.Name())
	}

	return roles, nil
}

// GetSigninToken fetches signin token based on token URL
func GetSigninToken(tokenURL string) (string, error) {
	req, err := http.NewRequest("GET", tokenURL, nil)

	// client := &http.Client{}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var tokenRes signinTokenResponse
	err = json.Unmarshal(r, &tokenRes)
	if err != nil {
		return "", err
	}

	return tokenRes.SigninToken, nil
}

// GetAssumeRoleCreds fetches temporary role creds for a role
func GetAssumeRoleCreds(role string) (creds credentials.Value, region string, err error) {
	// SharedConfigEnable is needed to support profiles with assume role config
	sessOpts := session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Profile:           role,
	}
	sess := session.Must(session.NewSessionWithOptions(sessOpts))
	creds, err = sess.Config.Credentials.Get()
	if err != nil {
		return creds, "", err
	}

	return creds, *sess.Config.Region, nil
}
