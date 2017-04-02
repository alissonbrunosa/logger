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
	logger.Debug("debug")
	logger.Info("info")
	logger.Warn("warn")
	logger.Error("error")
}
```

## Output
![Result output](https://github.com/alissonbrunosa/logger/blob/master/result.png)
