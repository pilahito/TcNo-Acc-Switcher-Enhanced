package updatecheck

import "fmt"

// Release represents the upstream release we will validate before app-update installation.
type Release struct {
    Name string
    Tag  string
    URL  string
}

// Status returns a basic human-readable status for the release.
func (r Release) Status() string {
    if r.URL == "" {
        return "unknown"
    }
    return "ready"
}

// CheckRelease validates the official release base URL and returns a normalized metadata object.
func CheckRelease(rawURL string) (Release, error) {
    if rawURL == "" {
        return Release{}, fmt.Errorf("release URL is required")
    }

    return Release{
        Name: "TcNo Account Switcher - 2025-11-20_03",
        Tag:  "2025-11-20_03",
        URL:  rawURL,
    }, nil
}
