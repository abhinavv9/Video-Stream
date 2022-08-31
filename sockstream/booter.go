package sockstream

import "gocv.io/x/gocv"

type Booter struct {
	hub     *Hub
	router  *Router
	capture *Capture
}

func NewBooter() *Booter {
	capCh := make(chan gocv.Mat)
	hub := NewHub(capCh)

	return &Booter{
		hub:     hub,
		router:  NewRouter(hub),
		capture: NewCapture(hub),
	}
}

func (booter *Booter) Run() {
	go booter.hub.Run()
	go booter.capture.Run()
	booter.router.Run()
}
