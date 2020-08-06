package net

import (
	"errors"
	"fmt"
	"net"

	"gopkg.in/irc.v3"
)

type ircWriter struct {
	client  *irc.Client
	channel string
}

func (i ircWriter) Write(msg string) error {
	return i.client.Write("PRIVMSG " + i.channel + " :" + msg)
}

func mkIRC(config Config) (Connection, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", config.Host, config.Port))
	if err != nil {
		return nil, errors.New("create irc socket: " + err.Error())
	}

	conf := irc.ClientConfig{
		Nick: config.User,
		Pass: config.Pass,
		User: config.User,
		Name: config.User,
		Handler: irc.HandlerFunc(func(c *irc.Client, m *irc.Message) {
			if m.Command == "001" {
				c.Write("JOIN " + config.Channel)
			}
		}),
	}

	client := irc.NewClient(conn, conf)

	// FIXME this is pretty ugly
	go client.Run()
	/* if err := client.Run(); err != nil {
		return nil, errors.New("start irc client: " + err.Error())
	} */

	return ircWriter{client: client, channel: config.Channel}, nil
}
