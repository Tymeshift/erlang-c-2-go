# erlang-c-js-2-go

## Installation
``` go get github.com/Tymeshift/erlang-c-go```

## Example usage
``` go
package main

import (
	"github.com/Tymeshift/erlang-c-go"
	"fmt"
)

func main() {
	// GetNumberOfAgents(volume, intervalLength, aht, targetServiceLevel, targetTime, maxOccupancy, shrinkage)
	numberOfAgents, err := erlangc.GetNumberOfAgents(50, 3600, 420, 0.8, 90, 0.8, 0.2)
	if err != nil {
		// Handle error
	}
	
	fmt.Println(numberOfAgents)
}

```
