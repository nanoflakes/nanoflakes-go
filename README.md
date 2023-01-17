# Nanoflakes - Go

[![GitHub issues](https://img.shields.io/github/issues/nanoflakes/nanoflakes-go)](https://github.com/nanoflakes/nanoflakes-go/issues)
[![License](https://img.shields.io/github/license/nanoflakes/nanoflakes-go)](https://github.com/nanoflakes/nanoflakes-go/tree/master/LICENSE)
[![Twitter](https://img.shields.io/twitter/url?style=social&url=https%3A%2F%2Fgithub.com%2Fnanoflakes%2Fnanoflakes-go)](https://twitter.com/intent/tweet?text=Wow:&url=https%3A%2F%2Fgithub.com%2Fnanoflakes%2Fnanoflakes-go)

Reference implementation of [nanoflakes](https://github.com/nanoflakes/nanoflakes) for Go.

Licensed under the [MIT License](https://github.com/nanoflakes/nanoflakes-java/blob/master/LICENSE).

### Installation

Run the following command to install the package:

```shell
$ go get github.com/nanoflakes/nanoflakes-go
```

### Usage

- Use `nanoflakes.localGenerator(epoch, generatorId)` to create a local nanoflake generator.
    - You can get an epoch by calling `time.Now().UnixMilli()`. Nanoflake's epoch is the Unix timestamp in milliseconds.
    - A generator ID must be in the 0-1023 range.
- Use `nanoflakes.NanoflakeGenerator.next()` to get a new nanoflake.
- The `Nanoflake` struct is the result type of `NanoflakeGenerator.next()`. It can be used as-is, or getting it's raw or encoded value. It also features utility methods such as getting the creation time of the nanoflake. 