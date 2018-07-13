package main

import (
	log "github.com/Sirupsen/logrus"
	"net"
)

func main() {
	log.Info("Starting...")
	defer log.Info("Exiting.")

	// single threaded, just POC
	for {

		// listen to incoming udp packets
		pc, err := net.ListenPacket("udp", ":2000")
		if err != nil {
			log.WithField("error", err).
				Fatal("Could not listen to UDP socket") // handle error
		}
		defer pc.Close()

		//simple read, just POC
		buffer := make([]byte, 1024)
		pc.ReadFrom(buffer)

		log.WithField("data", buffer).
			Info("Received data")

	}

}
