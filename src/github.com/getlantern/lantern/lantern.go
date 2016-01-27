// Package lantern provides an embeddable client-side web proxy
package lantern

import (
	"fmt"
	"sync"
	"time"

	"github.com/getlantern/appdir"
	"github.com/getlantern/flashlight"
	"github.com/getlantern/flashlight/client"
	"github.com/getlantern/flashlight/config"
	"github.com/getlantern/golog"
)

var (
	log = golog.LoggerFor("lantern")

	startOnce sync.Once
)

// Start starts a client proxy at a random address. It blocks up till the given
// timeout waiting for the proxy to listen, and returns the address at which it
// is listening.
func Start(appName string, timeout time.Duration) (string, error) {
	startOnce.Do(func() {
		go run(appName)
	})
	addr, ok := client.Addr(timeout)
	if !ok {
		return "", fmt.Errorf("Proxy didn't start within given timeout")
	}
	return addr.(string), nil
}

func run(appName string) {
	flashlight.Start(appdir.General("lantern_"+appName),
		false,
		func() bool { return true },
		make(map[string]interface{}),
		func(cfg *config.Config) bool { return true },
		func(cfg *config.Config) {},
		func(cfg *config.Config) {},
		func(err error) {},
	)
}
