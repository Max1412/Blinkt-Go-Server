package serverbackend

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/Max1412/blinkt_server/internal/pkg/ledapps"
	"github.com/alexellis/blinkt_go"
)

var ledWG sync.WaitGroup
var standardWG sync.WaitGroup
var stopchan = make(chan bool)

// TODOS:
// - use http pages with links, use /led/...
// - link to stop when starting a task
// - Add wake up light app
// - Add other apps from examples

func work() {
	fmt.Printf("I do work, getting called in a loop\n")
	time.Sleep(1 * time.Second)
}

func waitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return false // completed normally
	case <-time.After(timeout):
		return true // timed out
	}
}

// executes something that is passed asnychronously and can be stopped
// loopFunc must not contain any endlessly running loops
func executeStoppable(lApp ledapps.LedAppInterface, wg *sync.WaitGroup) {
	// signal being done
	defer wg.Done()

	lApp.Setup()

	defer func() {
		lApp.Cleanup()
	}()

	for {
		select {
		default:
			lApp.Loop() // do work here (no endless loops!)
		case <-stopchan:
			// stop
			return
		}
	}
}

// handles "/SolidColor/"
func HandlerLEDSolidColor(w http.ResponseWriter, r *http.Request) {
	if waitTimeout(&ledWG, time.Second) {
		fmt.Fprintf(w, "There is already a LED task running or stuck")
	} else {
		blinktApp := &ledapps.SolidColor{}
		ledWG.Add(1)
		go executeStoppable(blinktApp, &ledWG)
		fmt.Fprintf(w, "Started LED task")
	}
}

// handles "/Progress/"
func HandlerLEDProgress(w http.ResponseWriter, r *http.Request) {
	if waitTimeout(&ledWG, time.Second) {
		fmt.Fprintf(w, "There is already a LED task running or stuck")
	} else {
		blinktApp := &ledapps.Progress{}
		ledWG.Add(1)
		go executeStoppable(blinktApp, &ledWG)
		fmt.Fprintf(w, "Started LED task")
	}
}

// handles "/WakeUp/"
func HandlerLEDWakeUp(w http.ResponseWriter, r *http.Request) {
	if waitTimeout(&ledWG, time.Second) {
		fmt.Fprintf(w, "There is already a LED task running or stuck")
	} else {
		blinktApp := &ledapps.WakeUp{}
		ledWG.Add(1)
		go executeStoppable(blinktApp, &ledWG)
		fmt.Fprintf(w, "Started LED task")
	}
}

// handles "/stop/"
// needs handlers for stopping other WGs when implemented
func HandlerStopAsync(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "stopping...\n")
	stopchan <- true // tell the async func to stop
	ledWG.Wait()
	fmt.Fprintf(w, "Stopped.")
}

// handles "/"
func Handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	if len(path) == 0 {
		fmt.Fprint(w, "Welcome to the server")
	} else {
		fmt.Fprintf(w, "Nothing found at: %s!", path)
	}
}

func LedCleaner() {
	blinkt := blinkt.NewBlinkt(0.5)
	blinkt.Clear()
	blinkt.Show()
}
