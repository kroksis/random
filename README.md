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
Copyright © 2016 Eduards Bistrums

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
