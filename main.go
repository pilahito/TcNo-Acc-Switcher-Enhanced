package main

import (
    "flag"
    "fmt"
    "os"

    "pilahito.com/tcno-acc-switcher-enhanced/internal/accounts"
    "pilahito.com/tcno-acc-switcher-enhanced/internal/updatecheck"
)

const defaultReleaseURL = "https://github.com/TCNOco/TcNo-Acc-Switcher/releases/tag/2025-11-20_03"

func main() {
    checkRelease := flag.Bool("check-release", false, "Check the upstream release URL")
    updateAll := flag.Bool("update-all", false, "Run the batch update flow for all known accounts")
    releaseURL := flag.String("release-url", defaultReleaseURL, "Release URL to validate")
    flag.Parse()

    switch {
    case *checkRelease:
        if err := runCheckRelease(*releaseURL); err != nil {
            fmt.Fprintln(os.Stderr, err)
            os.Exit(1)
        }
    case *updateAll:
        if err := runUpdateAll(); err != nil {
            fmt.Fprintln(os.Stderr, err)
            os.Exit(1)
        }
    default:
        printUsage()
        os.Exit(2)
    }
}

func runCheckRelease(releaseURL string) error {
    rel, err := updatecheck.CheckRelease(releaseURL)
    if err != nil {
        return err
    }

    fmt.Printf("Release detected: %s\n", rel.Name)
    fmt.Printf("Tag: %s\n", rel.Tag)
    fmt.Printf("URL: %s\n", rel.URL)
    fmt.Printf("Status: %s\n", rel.Status())
    return nil
}

func runUpdateAll() error {
    accountsList := []accounts.Account{
        {Name: "Steam - Main", Platform: "Steam", Enabled: true},
        {Name: "Steam - Secondary", Platform: "Steam", Enabled: true},
        {Name: "Epic - Main", Platform: "Epic", Enabled: true},
        {Name: "Battle.net - Main", Platform: "Battle.net", Enabled: true},
        {Name: "Discord - Personal", Platform: "Discord", Enabled: true},
    }

    results := accounts.UpdateAll(accountsList, func(acc accounts.Account) error {
        fmt.Printf("Updating %s (%s)\n", acc.Name, acc.Platform)
        return nil
    })

    fmt.Println("Batch update summary:")
    for _, result := range results {
        if result.Err != nil {
            fmt.Printf("[FAIL] %s: %v\n", result.Account, result.Err)
            continue
        }
        fmt.Printf("[OK] %s: %s\n", result.Account, result.Message)
    }

    return nil
}

func printUsage() {
    fmt.Println("Usage:")
    fmt.Println("  go run . --check-release")
    fmt.Println("  go run . --update-all")
    fmt.Println("  go run . --release-url https://github.com/TCNOco/TcNo-Acc-Switcher/releases/tag/2025-11-20_03")
}
