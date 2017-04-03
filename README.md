# Logger
A Simple color logger for Go

## Start using it

Download and install it:
```shell
$ go get github.com/alissonbrunosa/logger
```
Import it in your code:
```go
import "github.com/alissonbrunosa/logger"
```
## Usage

```Go
package main

import "github.com/alissonbrunosa/logger"

func main() {
	log := logger.New()

	log.Debug("debug")
	log.Info("info")
	log.Warn("warn")
	log.Error("error")
}
```

## Output
![Result output](https://github.com/alissonbrunosa/logger/blob/master/result.png)
