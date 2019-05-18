package browser

import (
	"fmt"
	"os/exec"
)

// GetAppName returns the browser real application name based on browser short name
// Example of browserName: firefox, chrome, etc.
func GetAppName(browserName string) (string, error) {
	switch browserName {
	case "chrome":
		return "Google Chrome", nil
	case "firefox":
		return "Firefox", nil
	case "safari":
		return "Safari", nil
	case "brave":
		return "Brave Browser", nil
	default:
		return "", fmt.Errorf("browser %q is not supported", browserName)
	}
}

// Open opens an URL with browser app
func Open(appName string, url string) {
	exec.Command("open", "-a", appName, url).Run()
}
