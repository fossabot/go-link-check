# [Go Link Check](https://github.com/dbtedman/go-link-check)

Link check automation tool.

[![Actions Status](https://github.com/dbtedman/go-link-check/workflows/test/badge.svg)](https://github.com/dbtedman/go-link-check/actions)
[![codecov](https://codecov.io/gh/dbtedman/go-link-check/branch/master/graph/badge.svg)](https://codecov.io/gh/dbtedman/go-link-check)
[![Known Vulnerabilities](https://snyk.io/test/github/dbtedman/go-link-check/badge.svg)](https://snyk.io/test/github/dbtedman/go-link-check)
![Go Version](https://img.shields.io/static/v1?label=Go&message=v1.13&color=blue&style=flat)

## Getting Started

### Install

Install Go module dependencies.

```bash
go mod vendor
```

### Build

Build executable from Go source code.

```bash
go build -v -mod=vendor -o ./go-link-check
```

### Run

Run `go-link-check` to generate link check report.

```bash
./go-link-check -url="https://danieltedman.com"
```

### Help

To learn more about the available configuration options, view the help.

```bash
./go-link-check -help
```

You will see something like the following.

```
Usage of ./go-link-check:
  -outFile string
        (optional) Path to output report csv to. (default "results.csv")
  -url string
        Website URL to parse for links to validate.
```

### Docker

Alternatively [Docker](https://www.docker.com/) can be used to build and run the application.

```bash
docker run -it --rm $(docker build --quiet .) -url="https://danieltedman.com"
```

## Want to lean more?

-   See our [Contributing Guide](CONTRIBUTING.md) for details on how this repository is developed.
-   See our [Changelog](CHANGELOG.md) for details on which features, improvements, and bug fixes have been implemented
-   See our [License](LICENSE.md) for details on how you can use the code in this repository.
-   See our [Security Guide](SECURITY.md) for details on how security is considered.
