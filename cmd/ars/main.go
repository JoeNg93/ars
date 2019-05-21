package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/JoeNg93/ars/pkg/aws"
	"github.com/JoeNg93/ars/pkg/browser"
	"log"
	"strings"
	"time"

	"github.com/manifoldco/promptui"
)

func promptRole(roles []string) (string, error) {
	prompt := promptui.Select{
		Label: "Select role",
		Items: roles,
		Searcher: func(input string, idx int) bool {
			// Searcher will be called on each element in "Items", idx is the current index of the
			// loop. If Searcher returns true for an element -> that element will be included
			// as search result
			role := roles[idx]

			if strings.Contains(role, input) {
				return true
			}

			return false
		},
	}

	_, role, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return role, nil
}

func prompDestinationURL() (string, error) {
	prompt := promptui.Prompt{
		Label: "Redirect URL",
		Validate: func(input string) error {
			if input == "" {
				return errors.New("redirect url should not be empty")
			}
			return nil
		},
	}
	return prompt.Run()
}

func main() {
	browserName := flag.String("browser", "chrome", "Browser to open AWS Console. Accepted values: firefox, chrome, safari, brave")
	destinationURL := flag.String("redirect-url", "https://<profile_region>.console.aws.amazon.com", "Redirect URL after login")
	askDestinationURL := flag.Bool("ask-redirect-url", false, "Prompt redirect URL")

	flag.Parse()

	appName, err := browser.GetAppName(*browserName)
	if err != nil {
		log.Fatalln(err)
	}

	roles, err := aws.GetRoles()
	if err != nil {
		log.Fatalf("get roles failed: %v", err)
	}

	role, err := promptRole(roles)
	if err != nil {
		log.Fatalf("prompt role failed: %v", err)
	}

	roleCreds, region, err := aws.GetAssumeRoleCreds(role)
	if err != nil {
		log.Fatalf("get role creds failed: %v", err)
	}

	tokenURL := aws.TokenURL{
		SessionID:    roleCreds.AccessKeyID,
		SessionKey:   roleCreds.SecretAccessKey,
		SessionToken: roleCreds.SessionToken,
	}
	token, err := aws.GetSigninToken(tokenURL.String())
	if err != nil {
		log.Fatalf("get token failed: %v", err)
	}

	// Get signin URL
	if *askDestinationURL {
		url, err := prompDestinationURL()
		if err != nil {
			log.Fatalln(err)
		}
		*destinationURL = url
	} else if *destinationURL == "https://<profile_region>.console.aws.amazon.com" {
		*destinationURL = fmt.Sprintf("https://%s.console.aws.amazon.com", region)
	}

	loginURL := aws.LoginURL{
		DestinationURL: *destinationURL,
		SigninToken:    token,
	}

	fmt.Printf("Login URL: %v", loginURL)

	// Get signout URL
	logoutURL := aws.LogoutURL{}

	browser.Open(appName, logoutURL.String())
	time.Sleep(2 * time.Second) // Wait for log out
	browser.Open(appName, loginURL.String())
}
