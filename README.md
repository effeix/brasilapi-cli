# brasilapi-cli

CLI for the BrasilAPI project, written in Go

## Installing

### Homebrew

```sh
brew install effeix/tap/bra
```

### From source

Requires installing Golang. Check out the official tutorial at: [https://go.dev/doc/install](https://go.dev/doc/install).

```sh
git clone git@github.com:effeix/brasilapi-cli.git

make build
```

The binary will be generated at `./bin/bra`.

## Contributing

Contributions are always welcome. Feel free to drop a Pull Request or raise an Issue.

### Releasing

1. Generate a new tag with the updated version

    ```sh
    git tag vX.X.X

    git push origin vX.X.X
    ```

2. Run goreleaser

    ```sh
    GITHUB_TOKEN=<token> goreleaser release --clean
    ```

This will generate a new release and update the effeix/homebrew-tap repository automatically.
