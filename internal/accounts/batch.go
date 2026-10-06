package accounts

import "fmt"

// Account represents a configured game/platform account.
type Account struct {
    Name     string
    Platform string
    Enabled  bool
}

// Result stores the outcome of a single account update.
type Result struct {
    Account string
    Message string
    Err     error
}

// UpdateAll attempts the same update flow across every enabled account.
func UpdateAll(accounts []Account, updateFn func(Account) error) []Result {
    results := make([]Result, 0, len(accounts))

    for _, acc := range accounts {
        if !acc.Enabled {
            results = append(results, Result{
                Account: acc.Name,
                Message: "skipped (disabled)",
            })
            continue
        }

        if err := updateFn(acc); err != nil {
            results = append(results, Result{
                Account: acc.Name,
                Message: "failed",
                Err:      fmt.Errorf("%s: %w", acc.Name, err),
            })
            continue
        }

        results = append(results, Result{
            Account: acc.Name,
            Message: "updated successfully",
        })
    }

    return results
}
