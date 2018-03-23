# Blinkt-Go-Server

A webserver written in Go, supposed to run on a Raspberry Pi to remotely control [Blinkt!](https://shop.pimoroni.de/products/blinkt) LEDs.
It uses the following libraries:
* https://github.com/alexellis/blinkt_go
* https://github.com/lucasb-eyer/go-colorful

This is my first project written in Go.

I'm trying to use the Go project layout: https://github.com/golang-standards/project-layout

### Usage:

* Download the repository: `go get github.com/Max1412/Blinkt-Go-Server`
* Compile and run the `server` executable
* Navigate to `<your-ip-here>:8080/`
* Run an LED-App asynchonously: `<your-ip-here>:8080/<led-app-name>`, e.g. `SolidColor`
* Stop the LED-App: `<your-ip-here>:8080/stop`

### Adding new LED-Apps:

All new Apps must implement the Interface:
```
type LedAppInterface interface {
    Setup()
    Loop()
    Cleanup()
}
```

* `Setup()`: Set inital parameters, initialize the Blinkt!-LEDs
* `Loop()`: This will be called from the backend in a loop. Do not use endless loops in here!
* `Cleanup()`: Clear the LEDs. Will be called when navigating to `<your-ip-here>:8080/stop`

### Future Plans
- [ ] Nicer front-end with links to all possible LED-Apps
- [ ] Extend the interface to pass parameters via URLs or through the HTML pages
- [ ] Automatically start the `WakeUp`-App at a specific time of the day (given by user input)
