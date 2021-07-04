[![Go Reference](https://pkg.go.dev/badge/github.com/sitnikovv/tzchanger.svg)](https://pkg.go.dev/github.com/sitnikovv/tzchanger)
[![Go Report Card](https://goreportcard.com/badge/github.com/sitnikovv/tzchanger)](https://goreportcard.com/report/github.com/sitnikovv/tzchanger)

# golang TimeZone Changer

Change timezone without change time

# Usage

```go
package main

import (
	"fmt"
	"time"

	"github.com/sitnikovv/tzchanger"
)

func main() {
	const cleanTime = "2006-01-02 15:04:05"
	utcTime, _ := time.Parse(cleanTime, "2021-07-01 05:04:03")
	fmt.Println(utcTime)
	mskTime, _ := tzchanger.ChangeTZ(utcTime, "Europe/Moscow")
	fmt.Println(mskTime)
	vldTime, _ := tzchanger.ChangeTZ(mskTime, "Asia/Vladivostok")
	fmt.Println(vldTime)
	pdtTime, _ := tzchanger.ChangeTZ(vldTime, "America/Los_Angeles")
	fmt.Println(pdtTime)
}
```

# Output

```
2021-07-01 05:04:03 +0000 UTC
2021-07-01 05:04:03 +0300 MSK
2021-07-01 05:04:03 +1000 +10
2021-07-01 05:04:03 -0700 PDT
```