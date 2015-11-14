package stop

import (
	"os"
	"os/signal"
	"syscall"
)

// Ch – chanel to fetch stop signal.
var Ch = make(chan os.Signal, 1)

func init() {
	signal.Notify(Ch, syscall.SIGINT, syscall.SIGTERM)
}

// Fire termination signal to the chanel.
func Fire() {
	Ch <- syscall.SIGTERM
}
