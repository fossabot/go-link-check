# [Go Link Check](./README.md) / Contributing

-   [Resources](#resources)
-   [Pre-Commit Hooks](#pre-commit-hooks)
-   [Formatting and Linting](#formatting-and-linting)
-   [Testing](#testing)
-   [Dependencies](#dependencies)

## Resources

-   [Effective Go](https://golang.org/doc/effective_go.html)
-   [Go by Example](https://gobyexample.com/)
-   [Just tell me how to use Go Modules](https://engineering.kablamo.com.au/posts/2018/just-tell-me-how-to-use-go-modules)

## Pre-Commit Hooks

Linting and Testing is performed before each commit by [Husky](https://github.com/typicode/husky), see `.husky.js` for configuration.

## Formatting and Linting

Auto-formatting is being used to help develop consistently formatted source code.

[Go Format](https://golang.org/cmd/gofmt/) for `.go` files.

```bash
gofmt -w .
```

> TODO: Add linting to CI.

[Prettier](https://github.com/prettier/prettier) for `.js`, `.md`, `.yml`, and `.json` files.

```bash
yarn install && yarn run format
```

> TODO: Add linting to CI.

[Hadolint](https://github.com/hadolint/hadolint) for `Dockerfile`

> TODO: Add to CI.

```bash
docker run --rm -i hadolint/hadolint < Dockerfile
```

## Testing

```bash
go test ./... -cover
```

## Dependencies

[Go Modules](https://github.com/golang/go/wiki/Modules) has been chosen to manage Go dependencies for this repository. Vendoring is used to install these modules into `./vendor` directory instead of into `$GOPATH` which helps to isolate this project from other Go projects being developed on the same workstation. Running `go mod vendor` to install dependencies and then providing `-mod=vendor` argument to the `go build` command enables this functionality.
