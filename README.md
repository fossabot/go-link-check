# [Go Link Check](https://github.com/dbtedman/go-link-check)

Link check automation tool.

[![Actions Status](https://github.com/dbtedman/go-link-check/workflows/test/badge.svg)](https://github.com/dbtedman/go-link-check/actions)
[![Known Vulnerabilities](https://snyk.io/test/github/dbtedman/go-link-check/badge.svg)](https://snyk.io/test/github/dbtedman/go-link-check)
![Go Version](https://img.shields.io/static/v1?label=Go&message=v1.13&color=blue&style=flat)

## Getting Started

```bash
go mod vendor \
    && go build -v -mod=vendor -o ./go-link-check \
    && ./go-link-check
```

or with [Docker](https://www.docker.com/)

```bash
docker run -it --rm $(docker build --quiet .)
```

## Want to lean more?

-   See our [Contributing Guide](CONTRIBUTING.md) for details on how this repository is developed.
-   See our [Changelog](CHANGELOG.md) for details on which features, improvements, and bug fixes have been implemented
-   See our [License](LICENSE.md) for details on how you can use the code in this repository.
-   See our [Security Guide](SECURITY.md) for details on how security is considered.
