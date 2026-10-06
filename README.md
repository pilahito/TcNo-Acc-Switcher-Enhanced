# TcNo Account Switcher Enhanced

This repository is a fork-oriented enhancement workspace for building a stronger, self-updating, batch-update version of the official TcNo Account Switcher.

Base release used for alignment:
- https://github.com/TCNOco/TcNo-Acc-Switcher/releases/tag/2025-11-20_03

The goal of this fork is to add the following features on top of the upstream app:

- Update the app from inside the app
- Update all configured accounts in one action
- Better release validation and rollback safety
- Cleaner batch-progress reporting and summary
- Stronger troubleshooting for Epic/Battle.net/Discord/Ubisoft issues

## Features currently modeled in this repo

- `CheckRelease` helper for upstream release validation flow
- `UpdateAll` batch logic for accounts
- CLI demo for checking releases and processing all accounts
- Release-oriented structure suitable for a first public fork

## Quick start

```bash
go run . --check-release

go run . --update-all
```

## Project structure

```text
cmd/                 optional future CLI integration
internal/
  accounts/
  updatecheck/
main.go              demo entry point
README.md
CHANGELOG.md
```

## Release plan

1. Align this fork to the latest stable official release.
2. Add app-internal update flow.
3. Add `Update All` actions in the UI.
4. Add release verification and rollback safety.
5. Publish the first tagged public release.

## Notes

This repository is intentionally structured as a strong enhancement base. It is not the upstream project itself; it is a fork-ready enhancement layer built around the official release.
