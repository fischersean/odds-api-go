*This project has migrated to [sourcehut](https://git.sr.ht/~swf/odds-api-go)*

# odds-api-go 
[![Go Reference](https://pkg.go.dev/badge/github.com/fischersean/odds-api-go.svg)](https://pkg.go.dev/github.com/fischersean/odds-api-go)

Unofficial Go client library for [the-odds-api](https://the-odds-api.com).

## Install
```go
import "github.com/fischersean/odds-api-go"
```

Then run `go mod tidy`
## Usage
### Getting valid sports

```go
package main

import (
	"fmt"
	"github.com/fischersean/odds-api-go"
)

func main() {
	c := oddsapi.NewClient("YOUR API KEY HERE")
	s, err := c.GetSports(oddsapi.GetSportsInput{
		IncludeInactive: false,
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(s)
}
```

### Getting events data
```go
package main

import (
	"fmt"
	"github.com/fischersean/odds-api-go"
)

func main() {
	c := oddsapi.NewClient("YOUR API KEY HERE")
	s, err := c.GetEvents(oddsapi.GetEventsInput{
		Sports:  []string{"americanfootball_nfl"},
		Regions: []oddsapi.RegionType{oddsapi.RegionUS},
		Markets: []oddsapi.MarketType{oddsapi.MarketMoneyline},
	})
	if err != nil {
		fmt.Println(err.Error())
        	return
	}
	fmt.Println(s)
}
```
