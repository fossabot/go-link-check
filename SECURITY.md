# [Go Link Check](./README.md) / Security Guide

-   [Dependency Vulnerability Scanning](#dependency-vulnerability-scanning)
-   [Security Disclosure Policy](#security-disclosure-policy)
-   [Security Update Policy](#security-update-policy)
-   [Security Related Configuration](#security-related-configuration)
-   [Known Security Gaps and Future Enhancements](#known-security-gaps-and-future-enhancements)

## Dependency Vulnerability Scanning

Dependency scanning provided by [Snyk](https://snyk.io/test/github/dbtedman/go-link-check) for NPM and Docker dependencies.

```bash
yarn audit:npm && yarn audit:docker
```

## Security Disclosure Policy

Email [dbtedman@gmail.com](mailto:dbtedman@gmail.com) with details about the security issue.

## Security Update Policy

_Content to come._

## Security Related Configuration

_Content to come._

## Known Security Gaps and Future Enhancements

-   Go modules are not included in dependency vulnerability scanning.
