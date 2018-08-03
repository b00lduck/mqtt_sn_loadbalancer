package main

import (
	log "github.com/Sirupsen/logrus"
	"net"
    "time"
    "os"
    "strings"
    "bufio"
)

func main() {
	log.Info("Starting UDP packet comparator...")

	// Configuration
	port := os.Getenv("PORT")
    if len(port) == 0 {
        port = "20000"
        log.Info("Using default port")
    }

    timeout := os.Getenv("TIMEOUT")
    if len(timeout) == 0 {
        timeout = "2"
        log.Info("Using default timeout")
    }
    timeoutDuration, err := time.ParseDuration(timeout + "s")
    if err != nil {
        log.Fatal("Invalid timeout", err)
    }

    log.WithField("timeout", timeout).
        WithField("port", port).
        Info("Configuration")


    log.Info("Waiting for expected data on stdin...")


    // Read expected data
    expectedData := make([]byte, 65536)
    reader := bufio.NewReader(os.Stdin)
    expectedLength, err := reader.Read(expectedData)
    if err != nil {
        log.Fatal("Error reading from stdin", err)
    }

    log.WithField("length", expectedLength).
        Info("Got expected data from stdin")


    log.Info("Waiting for actual data via UDP socket...")

	// listen to incoming udp packets
	pc, err := net.ListenPacket("udp", ":" + port)
	if err != nil {
		log.WithField("error", err).
		    WithField("port", port).
			Fatal("Could not listen to UDP socket")
	}
	defer pc.Close()

    pc.SetReadDeadline(time.Now().Add(time.Duration(timeoutDuration)))

    buffer := make([]byte, 65536)
    actualLength, _, err := pc.ReadFrom(buffer)
    if err != nil {
        if strings.Contains(err.Error(), "timeout") {
            log.Error("Timeout")
            os.Exit(2)
        }
        log.Error(err)
        os.Exit(1)
    }

    if expectedLength != actualLength {
        log.WithField("expected", len(expectedData)).
            WithField("actual", actualLength).
            Error("Length mismatch")
        os.Exit(3)
    }

    for k,v := range expectedData {
        if buffer[k] != v {
            log.WithField("expected", v).
                WithField("actual", buffer[k]).
                WithField("position", k).
                Error("Byte mismatch")
            os.Exit(3)
        }
    }

    log.Info("Match.")
}


