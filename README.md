# random
Go library for randomness

## Install
1. go get github.com/kroksis/random
2. import "github.com/kroksis/random"

## Example
```go
package main

import (
	"fmt"
	"github.com/kroksis/random"
	"time"
)

func main() {
	for {
		time.Sleep(1 * time.Second)
		fmt.Println(random.String(64))
		fmt.Println(random.StringReadable(64))
		fmt.Println(random.Int(64))
	}
}
```

## License

BSD licensed. See the [LICENSE](https://github.com/kroksis/random/blob/master/LICENSE) file for details.