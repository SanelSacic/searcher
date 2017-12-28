// This program provides sample web service that uses concurrency
// and channels to perform a coordinated set of asynchronous searches.
package main

import (
	"expvar"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/gocode/SearchForNews/service"
)

// init is  called before main. We are using init to
// set the logging package.
func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	log.SetOutput(os.Stdout)
}

// expvars is adding the goroutine counts to the variable set.
func expvars() {

	// Add goroutine counts to the variable set.
	gr := expvar.NewInt("goroutines")
	go func() {
		for _ = range time.Tick(time.Millisecond * 250) {
			gr.Set(int64(runtime.NumGoroutine()))
		}
	}()
}

func main() {
	expvars()
	service.Run()
}
