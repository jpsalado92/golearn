# [Go Class: 37 Static Analysis](https://www.youtube.com/watch?v=GyoMEerSd0I)

Static analysis is also considered as linting. Here "static" means that the code is
analyzed without running it. Some of the categories these tools fall under are:

- bugs
- commented code
- complexity
- errors
- formatting issues
- imports
- module
- performance
- sql
- style
- test
- unused

Whenever a project is set for the first time, it is a good idea to set all of these tools.
Some options are the `Makefile`, the CI/CD pipeline, a pre-commit hook, or the IDE itself.

When writing this, the most popular tool is `golangci-lint`. It is a fast linters runner for Go.

### `golint` [deprecated] use https://golangci-lint.run/ instead

`golangci-lint` is a fast linters runner for Go. It runs linters in parallel, uses
caching, supports YAML configuration, integrates with all major IDEs, and includes over a
hundred linters.

This tool includes many linters that can be consulted here. By default the following
linters are enabled:

- `errcheck`: errcheck is a program for checking for unchecked errors in Go code. These unchecked errors can be critical bugs in some cases [fast: false, auto-fix: false]
- `gosimple`: Linter for Go source code that specializes in simplifying code [fast: false, auto-fix: false]
- `govet`: Vet examines Go source code and reports suspicious constructs. It is roughly the same as 'go vet' and uses its passes. [fast: false, auto-fix: false]
- `ineffassign`: Detects when assignments to existing variables are not used [fast: true, auto-fix: false]
- `staticcheck`: It's a set of rules from staticcheck. It's not the same thing as the staticcheck binary. The author of staticcheck doesn't support or approve the use of staticcheck as a library inside golangci-lint. [fast: false, auto-fix: false]
- `unused`: Checks Go code for unused constants, variables, functions and types [fast: false, auto-fix: false]
