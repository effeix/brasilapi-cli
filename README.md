# brasilapi-cli

CLI for the BrasilAPI project, written in Go

## Release

1. Generate new tag with new version

2. Run goreleaser

```sh
GITHUB_TOKEN=<token> goreleaser release --clean
```

This will generate a new release and update the effeix/homebrew-tap repository automatically.
