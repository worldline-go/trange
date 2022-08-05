# trange

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
