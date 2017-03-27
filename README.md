# package stack15

logging handler for log15 that uses the StackDriver Logging API of the Google Cloud Platform

Example
```
package main

import (
	"time"

	"github.com/inconshreveable/log15"

	"github.com/emicklei/log15-stackdriver-handler"
)

func main() {
	h, _ := stack15.NewHandler("<google-project-id>", "stack15-test")
	defer h.Close()
    
	srvlog := log15.New("module", "app/server")
	srvlog.SetHandler(h)

	srvlog.Info("its happing again", "year", 2017)
	srvlog.Info("laura isn't coming back", "when", time.Now())
}
```