# Logger
> 📋🦄 Colorful command-line logger

## Usage

```go
package main

import (
	"github.com/mineway/logger"
)

func main() {
	// Define log directory location
	logger.SetLogLocation("/var/log/logger")
	
	logger.Fatal("Hello world from %s !", "README.md")
	// Result => [20:10:15][fatal] Hello world from README.md !
	//           os exit (1)

	logger.Error("Hello world from %s !", "README.md")
	// Result => [20:10:15][error] Hello world from README.md !

	logger.Warning("Hello world from %s !", "README.md")
	// Result => [20:10:15][warning] Hello world from README.md !

	logger.Info("Hello world from %s !", "README.md")
	// Result => [20:10:15][info] Hello world from README.md !

	logger.Success("Hello world from %s !", "README.md")
	// Result => [20:10:15][success] Hello world from README.md !

	logger.Log("Hello world from %s !", "README.md")
	// Result => [20:10:15][log] Hello world from README.md !
}
```
