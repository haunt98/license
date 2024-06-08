# license

[![Go](https://github.com/haunt98/license/actions/workflows/go.yaml/badge.svg)](https://github.com/haunt98/license/actions/workflows/go.yaml)
[![gitleaks](https://github.com/haunt98/license/actions/workflows/gitleaks.yml/badge.svg)](https://github.com/haunt98/license/actions/workflows/gitleaks.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/haunt98/license.svg)](https://pkg.go.dev/github.com/haunt98/license)
[![Latest Version](https://img.shields.io/github/v/tag/haunt98/license)](https://github.com/haunt98/license/tags)

Generate license file (`LICENSE`, `COPYING`, ...) automatically.

Support:

- [MIT](https://choosealicense.com/licenses/mit/)
- [GNU GPLv3](https://choosealicense.com/licenses/gpl-3.0/)
- [Apache License 2.0](https://choosealicense.com/licenses/apache-2.0/)

## Install

With Go version `>= 1.16`:

```sh
go install github.com/haunt98/license@latest
```

With Go version `< 1.16`:

```sh
GO111module=on go get github.com/license/changeloguru
```

## Usage

```sh
license generate
```
