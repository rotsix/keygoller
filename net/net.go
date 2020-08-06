package net

import (
	"errors"
	"fmt"
)

// Connection to the listener
type Connection interface {
	Write(string) error
}

// Config stores informations to connect to listener
type Config struct {
	Host string
	Port int
	// auth
	User string
	Pass string
	// irc
	Channel string
	// http
	URL     string
	Request string
}

// Init a new connection
func Init(protocol string, config Config) (connection Connection, err error) {
	switch protocol {
	case "irc":
		connection, err = mkIRC(config)
	case "http":
		connection, err = mkHTTP(config)
	case "debug":
		connection, err = mkDEBUG(config)
	default:
		return nil, errors.New("protocol not supported")
	}
	return
}

// debug connection

type debugWriter struct{}

func (d debugWriter) Write(msg string) error {
	fmt.Println("> " + msg)
	return nil
}

func mkDEBUG(config Config) (Connection, error) {
	return debugWriter{}, nil
}
