# Common Functions

## Feature
* Random String
* Random Integer

## Example

```
package main

import (
	"fmt"

	"github.com/oscar3x39/gcom/random"
)

func main() {
	fmt.Println(random.String(10))
	fmt.Println(random.Int(10))
}
```