package main

import "github.com/abhinavv9/Video-stream/sockstream"

func main() {
	booter := sockstream.NewBooter()
	booter.Run()
}
