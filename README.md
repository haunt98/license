# license

[![Go](https://github.com/haunt98/license/workflows/Go/badge.svg?branch=main)](https://github.com/actions/setup-go)

Generate `LICENSE` file automatically.

Support:

- [MIT](https://choosealicense.com/licenses/mit/)

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
license --name MIT --output .
```
