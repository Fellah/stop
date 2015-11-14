# Stop

## Installation

```
$ go get -u github.com/fellah/stop
```

## Usage

Waiting for stop signal to finish Go application.

```go
package main

import (
	"github.com/fellah/stop"
}

func main() {
	go ToDoSomething()

	<- stop.Ch
	
	CloseResources()
}
```

Send stop signal.

```go
func SomeTask() {

	...
	
	stop.Fire()

	...

}
```
