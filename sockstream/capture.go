package sockstream

import (
	"log"

	"gocv.io/x/gocv"
)

// Capture is a wrapper around capturing video streams.
type Capture struct {
	dev   *gocv.VideoCapture
	frame gocv.Mat
	hub   *Hub
}

// NewCapture returns a reference to a new instance of Capture.
func NewCapture(hub *Hub) *Capture {
	dev, err := gocv.OpenVideoCapture(0)

	if err != nil {
		log.Fatal(err)
	}

	return &Capture{
		dev:   dev,
		frame: gocv.NewMat(),
		hub:   hub,
	}
}

// Run the capture process.
func (capture *Capture) Run() {
	defer capture.dev.Close()
	defer capture.frame.Close()

	for {
		capture.dev.Read(&capture.frame)
		buf, _ := gocv.IMEncodeWithParams(".jpg", capture.frame, []byte{IMWriteJpegQuality, quality})

		capture.hub.broadcast <- buf
	}
}
