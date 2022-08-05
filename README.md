# trange

[![Codecov](https://img.shields.io/codecov/c/github/worldline-go/trange?logo=codecov&style=flat-square)](https://app.codecov.io/gh/worldline-go/trange)
[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/worldline-go/trange/Test?logo=github&style=flat-square&label=ci)](https://github.com/worldline-go/trange/actions)
[![Go Reference](https://pkg.go.dev/badge/github.com/worldline-go/trange.svg)](https://pkg.go.dev/github.com/worldline-go/trange)

This repository help you to get the date range by given date

## Usage

Install

```go
go get github.com/worldline-go/trange
```

Example

```go
package main

import (
    "fmt"
    "time"

    "github.com/worldline-go/trange"
)

func main() {
    loc, _ := time.LoadLocation("Europe/Amsterdam")

    start, end, _ := trange.Day("2022-08-05", loc)

    fmt.Println("start", start, "end", end)
}
```

## Todo

- [x] Day
- [ ] Week
- [ ] Month
- [ ] Year

## Inspired by

[moment.js](https://github.com/moment/moment)
