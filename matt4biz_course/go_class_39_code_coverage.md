# [Go Class: 39 Code Coverage](https://www.youtube.com/watch?v=HfCsfuVqpcM&t=14s)

## Basic Rules

Code coverage is a metric that measures the percentage of code that is executed during the tests.

In order to generate a code coverage report, we can use the `-cover` flag with the `go test` command.

```bash
go test -cover
```

You can also generate an HTML report by using the `-coverprofile` flag.

```bash
go test -coverprofile=coverage.out
```

This will generate a `coverage.out` file that can be used to generate an HTML report.

```bash
go tool cover -html=coverage.out
```

When generating the coverage report, you can also use the `-covermode` flag to specify the coverage mode.
The default mode is `set`, which only reports whether a line was executed or not. The `count` mode reports the number of times a line was executed.

```bash
go test -covermode=count
```

